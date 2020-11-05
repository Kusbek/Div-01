package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"DIV-01/real-time-forum/internal/store"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

//ChatUserInfo ...
type ChatUserInfo struct {
	Room        *model.Room `json:"room"`
	User        *model.User `json:"user"`
	LastMessage *time.Time  `json:"last_message"`
}
type msg struct {
	Status      string      `json:"status"`
	Room        *model.Room `json:"room"`
	User        *model.User `json:"user"`
	LastMessage *time.Time  `json:"last_message"`
	NewMessage  bool        `json:"new_message"`
}

type guest struct {
	user *model.User
	conn *websocket.Conn
	ch   chan *msg
}

func (s *server) monitorDeleteGuestChan() {
	for {
		dGuest := <-s.deleteGuestChan
		s.mu.Lock()
		for userID, guest := range s.guests {
			if guest == dGuest {
				delete(s.guests, userID)
				continue
			}
			guest.ch <- &msg{Status: "offline", User: dGuest.user}
		}
		s.mu.Unlock()
	}
}

func (g *guest) monitorClient(del chan *guest) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Closing monitoring routine!!!!")
				return
			case msg := <-g.ch:
				g.sendMessage(msg)
			}

		}
	}()
	go func() {
		for {
			_, _, err := g.conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message.", err)
				del <- g
				cancel()
				return
			}
		}
	}()
}

func (g *guest) sendMessage(m *msg) {
	if err := g.conn.WriteJSON(m); err != nil {
		fmt.Println(err)
	}
}

func (s *server) chatWsHandler(w http.ResponseWriter, r *http.Request) {
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

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		s.error(w, http.StatusInternalServerError, errors.New("Could not open websocket connection"))
		return
	}

	g := &guest{
		user: user,
		conn: conn,
		ch:   make(chan *msg, 10),
	}

	chatUserInfos, err := getAllUsers(user.ID, s.store)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.mu.Lock()
	for _, cui := range chatUserInfos {
		if gu, ok := s.guests[cui.User.ID]; ok {
			gu.ch <- &msg{Status: "online", User: g.user}
			g.sendMessage(&msg{Status: "online", User: cui.User, Room: cui.Room, LastMessage: cui.LastMessage})
		} else {
			g.sendMessage(&msg{Status: "offline", User: cui.User, Room: cui.Room, LastMessage: cui.LastMessage})
		}
	}
	s.guests[g.user.ID] = g
	s.mu.Unlock()
	go g.monitorClient(s.deleteGuestChan)
}

func getAllUsers(userID int, st store.Store) ([]*ChatUserInfo, error) {
	users, err := st.User().GetAll(userID)
	if err != nil {
		return nil, err
	}
	chatUserInfos := make([]*ChatUserInfo, 0)
	for _, user := range users {
		cui := &ChatUserInfo{
			User: user,
		}
		chatUserInfos = append(chatUserInfos, cui)
		room, err := st.Room().GetRoom(userID, user.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			} else {
				return nil, err
			}
		}
		cui.Room = room
		lastMessageTimestamp, err := st.Room().GetLastMessageTimestamp(room.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			} else {
				return nil, err
			}
		}
		cui.LastMessage = lastMessageTimestamp

	}

	return chatUserInfos, nil
}
