package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type msg struct {
	Action string      `json:"action"`
	User   *model.User `json:"user"`
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
				// s.guests = append(s.guests[:i], s.guests[i+1:]...)
				delete(s.guests, userID)
				continue
			}
			guest.ch <- &msg{"offline", dGuest.user}
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

	allusers, err := s.store.User().GetUsers(user.ID)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(allusers)
	s.mu.Lock()

	for _, u := range allusers {
		if gu, ok := s.guests[u.ID]; ok {
			gu.ch <- &msg{"online", g.user}
			g.sendMessage(&msg{"online", u})
		} else {
			g.sendMessage(&msg{"offline", u})
		}
	}
	s.guests[g.user.ID] = g
	s.mu.Unlock()
	// s.mu.Lock()
	// for _, gu := range s.guests {
	// 	// g.sendMessage(&msg{"add", gu.user})
	// 	gu.ch <- &msg{"add", g.user}
	// }

	// s.mu.Unlock()

	go g.monitorClient(s.deleteGuestChan)
}
