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
	guest       *guest
}

func (il *interlocutor) sendMessage(m *model.Message) {
	type messageWrapper struct {
		Message *model.Message `json:"message"`
	}
	if err := il.conn.WriteJSON(&messageWrapper{Message: m}); err != nil {
		fmt.Println(err)
	}
}

type inMessage struct {
	Message string `json:"message"`
}

func (s *server) handleGetMessages(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		s.error(w, http.StatusUnauthorized, errors.New("No cookie"))
		return
	}
	_, err = s.cookies.Check(session.Value)
	if err != nil {
		s.error(w, http.StatusUnauthorized, err)
		return
	}
	roomID, err := strconv.Atoi(r.URL.Query().Get("room_id"))
	if err != nil {
		s.error(w, http.StatusBadRequest, errors.New("room_id is invalid"))
		return
	}

	pageNum, err := strconv.Atoi(r.URL.Query().Get("page_num"))
	if err != nil {
		s.error(w, http.StatusBadRequest, errors.New("page_num is invalid"))
		return
	}

	messages, err := s.store.Room().GetMessages(roomID, 10*pageNum)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, map[string]interface{}{
		"next":     pageNum + 1,
		"messages": messages,
	})
}

func (s *server) messageWsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		s.error(w, http.StatusUnauthorized, errors.New("Not Authorized"))
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

	// messages, err := s.store.Room().GetMessages(roomID, 0)
	il := &interlocutor{
		roomManager: s.rooms[roomID],
		room:        s.store.Room(),
		user:        user,
		conn:        conn,
		guest:       s.guests[user.ID],
	}

	il.roomManager.mu.Lock()
	il.roomManager.interlocutors = append(il.roomManager.interlocutors, il)
	il.roomManager.mu.Unlock()
	// for _, m := range messages {
	// 	il.sendMessage(m, false)
	// }
	go il.monitorMessages(roomID, s)
}

func (il *interlocutor) monitorMessages(roomID int, s *server) {
	fmt.Println(il.roomManager.interlocutors)
	go func() {
		for {
			_, d, err := il.conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message.", err)
				il.roomManager.mu.Lock()
				for i, interlocutor := range il.roomManager.interlocutors {
					if interlocutor == il {
						il.roomManager.interlocutors = append(il.roomManager.interlocutors[:i], il.roomManager.interlocutors[i+1:]...)
					}
				}
				fmt.Println(il.roomManager.interlocutors)
				il.roomManager.mu.Unlock()
				return
			}
			in := &inMessage{}
			json.Unmarshal(d, in)

			m := &model.Message{
				Timestamp: time.Now(),
				User:      il.user,
				Text:      in.Message,
			}

			err = il.room.NewMessage(roomID, m)
			if err != nil {
				fmt.Println("Error saving message.", err)
				continue
			}
			for _, i := range il.roomManager.interlocutors {
				i.sendMessage(m)
			}

			users, err := il.room.GetRoomUsers(roomID)

			for _, user := range users {
				if user.ID != il.user.ID {
					updMsg := &msg{User: user, LastMessage: &m.Timestamp}
					il.guest.sendMessage(updMsg)
					g, ok := s.guests[user.ID]
					if ok {
						g.sendMessage(&msg{User: il.user, Status: "online", LastMessage: &m.Timestamp, NewMessage: true})
					}
				}
			}
		}

	}()

	// go func() {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			for i, interlocutor := range il.roomManager.interlocutors {
	// 				if interlocutor == il {
	// 					il.roomManager.interlocutors = append(il.roomManager.interlocutors[:i], il.roomManager.interlocutors[i+1:]...)
	// 				}
	// 			}
	// 			return
	// 		}
	// 	}
	// }()

}

//MyPrint ...
func MyPrint(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
	fmt.Println()
}
