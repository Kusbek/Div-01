package apiserver

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

func (s *server) roomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetRoom(w, r)
	default:
		s.error(w, http.StatusMethodNotAllowed, errors.New("No such method"))
	}
}

var roomCount = 0

func (s *server) handleGetRoom(w http.ResponseWriter, r *http.Request) {
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

	guestID, err := strconv.Atoi(r.URL.Query().Get("guest_id"))
	if err != nil {
		s.error(w, http.StatusUnauthorized, errors.New("guest-id is invalid"))
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	roomCount++

	slc := []int{user.ID, guestID}
	sort.Ints(slc)
	key := fmt.Sprintf("%d_%d", slc[0], slc[1])
	if roomID, ok := s.roomIds[key]; ok {
		s.respond(w, http.StatusOK, map[string]interface{}{
			"room": roomID,
		})
		return
	}
	s.roomIds[key] = roomCount
	s.rooms[roomCount] = &room{
		ID: roomCount,
		messages: func() []*message {
			msgs := make([]*message, 0)
			msgs = append(msgs, &message{
				Timestamp: time.Now(),
				User:      user,
				Text:      "Idi nahoi",
			})
			for _, g := range s.guests {
				if g.user.ID == guestID {
					msgs = append(msgs, &message{
						Timestamp: time.Now(),
						User:      g.user,
						Text:      "sam idi nahoi",
					})
					break
				}
			}

			return msgs
		}(),
		interlocutors: make([]*interlocutor, 0),
		mu:            &sync.Mutex{},
	}
	s.respond(w, http.StatusOK, map[string]interface{}{
		"room": roomCount,
	})
}
