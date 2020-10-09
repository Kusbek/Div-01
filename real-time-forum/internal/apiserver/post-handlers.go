package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *server) handlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleGetPosts(w, r)
	case "POST":
		s.handleCreatePost(w, r)
	default:
		s.error(w, http.StatusMethodNotAllowed, errors.New("No such method"))
	}
}

func (s *server) handleGetPosts(w http.ResponseWriter, r *http.Request) {

	category := r.URL.Query().Get("category")
	if category == "" {
		category = "all"
	}
	fmt.Println(category)
	if category == "all" {
		posts, err := s.store.Post().GetAll()
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
		}

		s.respond(w, http.StatusOK, map[string]interface{}{
			"posts": posts,
		})
		return
	}

	posts, err := s.store.Post().Get(category)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
	}

	s.respond(w, http.StatusOK, map[string]interface{}{
		"posts": posts,
	})
}

//CreatePostParams ...
type CreatePostParams struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

func (r *CreatePostParams) getParams(req *http.Request) error {
	d := json.NewDecoder(req.Body)
	defer req.Body.Close()
	err := d.Decode(r)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) handleCreatePost(w http.ResponseWriter, r *http.Request) {
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

	d := &CreatePostParams{}
	err = d.getParams(r)
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}
	newPost := &model.Post{
		Author: &model.User{
			ID:       user.ID,
			Nickname: user.Nickname,
		},
		Title:    d.Title,
		Text:     d.Text,
		Category: d.Category,
	}

	err = s.store.Post().Create(newPost)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, map[string]interface{}{
		"post": newPost,
	})
}
