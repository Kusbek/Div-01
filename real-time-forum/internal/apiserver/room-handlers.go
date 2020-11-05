package apiserver

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"sync"
)

type roomManager struct {
	mu            *sync.Mutex
	ID            int
	interlocutors []*interlocutor
}

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
		s.error(w, http.StatusUnauthorized, errors.New("Not Authorized"))
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

	room, err := s.store.Room().GetRoom(user.ID, guestID)
	if err != nil {
		if err == sql.ErrNoRows {
			room, err = s.store.Room().CreateRoom(user.ID, guestID)
			if err != nil {
				s.error(w, http.StatusInternalServerError, errors.New("Failed to create room"))
				return
			}
		} else {
			s.error(w, http.StatusInternalServerError, errors.New("Failed to get room"))
			return
		}

	}
	s.addRoomManager(room.ID)
	s.respond(w, http.StatusOK, map[string]interface{}{
		"room": room,
	})
}

func (s *server) addRoomManager(id int) {
	s.mu.Lock()
	if _, ok := s.rooms[id]; !ok {
		s.rooms[id] = &roomManager{ID: id, interlocutors: make([]*interlocutor, 0), mu: &sync.Mutex{}}
	}
	s.mu.Unlock()
}
