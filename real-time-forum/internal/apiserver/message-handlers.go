package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"DIV-01/real-time-forum/internal/store"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type interlocutor struct {
	roomManager *roomManager
	room        store.RoomRepository
	user        *model.User
	conn        *websocket.Conn
}

func (il *interlocutor) sendMessage(m *model.Message) {
	if err := il.conn.WriteJSON(m); err != nil {
		fmt.Println(err)
	}
}

type inMessage struct {
	Message string `json:"message"`
}

func (s *server) messageWsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		s.error(w, http.StatusUnauthorized, errors.New("No cookie"))
		return
	}
	user, err := s.cookies.Check(session.Value)
	if err != nil {
		s.error(w, http.StatusUnauthorized, err)
		return
	}
	if !strings.Contains(r.Header.Get("Origin"), r.Host) {
		s.error(w, http.StatusForbidden, errors.New("Origin not allowed"))
		return
	}
	roomID, err := strconv.Atoi(r.URL.Query().Get("room_id"))
	if err != nil {
		s.error(w, http.StatusBadRequest, errors.New("room_id is invalid"))
		return
	}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		s.error(w, http.StatusInternalServerError, errors.New("Could not open websocket connection"))
		return
	}

	messages, err := s.store.Room().GetMessages(roomID, 0)
	il := &interlocutor{
		roomManager: s.rooms[roomID],
		room:        s.store.Room(),
		user:        user,
		conn:        conn,
	}
	fmt.Println(messages[0])
	il.roomManager.interlocutors = append(il.roomManager.interlocutors, il)

	for _, m := range messages {
		il.sendMessage(m)
	}
	go il.monitorMessages(roomID)
}

func (il *interlocutor) monitorMessages(roomID int) {
	go func() {
		for {
			_, d, err := il.conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message.", err)
				continue
			}
			in := &inMessage{}
			json.Unmarshal(d, in)

			msg := &model.Message{
				Timestamp: time.Now(),
				User:      il.user,
				Text:      in.Message,
			}

			err = il.room.NewMessage(roomID, msg)
			if err != nil {
				fmt.Println("Error saving message.", err)
				continue
			}
			for _, i := range il.roomManager.interlocutors {
				i.sendMessage(msg)
			}
		}
	}()
}
