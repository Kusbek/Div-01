package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type interlocutor struct {
	room *room
	user *model.User
	conn *websocket.Conn
}

func (il *interlocutor) sendMessage(m *message) {
	if err := il.conn.WriteJSON(m); err != nil {
		fmt.Println(err)
	}
}

type room struct {
	mu            *sync.Mutex
	ID            int
	messages      []*message
	interlocutors []*interlocutor
}

type message struct {
	Timestamp time.Time   `json:"timestamp"`
	User      *model.User `json:"user"`
	Text      string      `json:"text"`
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
	room, ok := s.rooms[roomID]
	fmt.Println(room)
	if !ok {
		s.error(w, http.StatusInternalServerError, errors.New("4e to ne mogu naiti takoy room"))
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		s.error(w, http.StatusInternalServerError, errors.New("Could not open websocket connection"))
		return
	}
	il := &interlocutor{
		room: room,
		user: user,
		conn: conn,
	}
	room.interlocutors = append(room.interlocutors, il)
	for _, m := range room.messages {
		il.sendMessage(m)
	}

	go il.monitorMessages()
}

func (il *interlocutor) monitorMessages() {
	go func() {
		for {
			_, d, err := il.conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message.", err)
				return
			}
			in := &inMessage{}
			json.Unmarshal(d, in)

			msg := &message{
				Timestamp: time.Now(),
				User:      il.user,
				Text:      in.Message,
			}

			il.room.mu.Lock()
			il.room.messages = append(il.room.messages, msg)
			for _, i := range il.room.interlocutors {
				i.sendMessage(msg)
			}
			il.room.mu.Unlock()
		}
	}()
}
