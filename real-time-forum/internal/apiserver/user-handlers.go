package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func (s *server) authHandler(w http.ResponseWriter, r *http.Request) {
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

	s.respond(w, http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (s *server) signInHandler(w http.ResponseWriter, r *http.Request) {
	type body struct {
		Creds    string `json:"creds"`
		Password string `json:"password"`
	}
	if r.Method != "POST" {
		s.error(w, http.StatusMethodNotAllowed, errors.New("Wrong Method"))
		return
	}
	b := &body{}
	err := json.NewDecoder(r.Body).Decode(b)
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}

	user, err := s.store.User().Find(b.Creds)
	if err != nil {
		if err == sql.ErrNoRows {
			s.error(w, http.StatusForbidden, errors.New("No such user"))
			return
		}
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	err = user.ComparePasswords(b.Password)
	if err != nil {
		s.error(w, http.StatusForbidden, err)
		return
	}

	s.setCookies(w, user)
	// s.error(w, http.StatusUnauthorized, errors.New("Unothorized"))
	s.respond(w, http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (s *server) signOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.error(w, http.StatusMethodNotAllowed, errors.New("Wrong Method"))
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		s.respond(w, http.StatusOK, map[string]interface{}{
			"success": true,
		})
		return
	}

	s.cookies.Delete(session.Value)
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	s.respond(w, http.StatusOK, map[string]interface{}{
		"success": true,
	})
}

func (s *server) signUpHandler(w http.ResponseWriter, r *http.Request) {
	type body struct {
		Nickname  string `json:"nickname,omitempty"`
		Age       uint8  `json:"age,omitempty"`
		Gender    string `json:"gender,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Email     string `json:"email,omitempty"`
		Password  string `json:"password"`
	}

	if r.Method != "POST" {
		s.error(w, http.StatusMethodNotAllowed, errors.New("Wrong Method"))
		return
	}

	b := &body{}
	err := json.NewDecoder(r.Body).Decode(b)
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}

	user := &model.User{
		Nickname:  b.Nickname,
		Age:       b.Age,
		Gender:    b.Gender,
		FirstName: b.FirstName,
		LastName:  b.LastName,
		Email:     b.Email,
		Password:  b.Password,
	}

	err = user.Validate()
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}

	exists, err := s.store.User().Exists(user.Nickname, user.Email)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	if exists {
		s.error(w, http.StatusConflict, errors.New("User already exists"))
		return
	}

	err = s.store.User().Create(user)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.setCookies(w, user)
	// s.error(w, http.StatusUnauthorized, errors.New("Unothorized"))
	s.respond(w, http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
