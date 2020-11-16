package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (s *server) handleComments(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		s.handleGetComments(w, r)
	case "POST":
		s.handleCreateComment(w, r)
	default:
		s.error(w, http.StatusMethodNotAllowed, errors.New("No such method"))
	}
}

type getCommentParams struct {
	postID int
}

func (r *getCommentParams) getParams(req *http.Request) error {
	postID, err := strconv.Atoi(req.URL.Query().Get("post_id"))
	if err != nil {
		return errors.New("post_id should not be empty and int")
	}
	r.postID = postID
	return nil
}

func (s *server) handleGetComments(w http.ResponseWriter, r *http.Request) {
	params := &getCommentParams{}
	err := params.getParams(r)
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}
	comments, err := s.store.Comment().Get(params.postID)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}
	s.respond(w, http.StatusOK, map[string]interface{}{
		"comments": comments,
	})
}

type createCommentParams struct {
	PostID int    `json:"post_id" form:"post_id"`
	Text   string `json:"text" form:"text"`
}

func (r *createCommentParams) getParams(req *http.Request) error {
	d := json.NewDecoder(req.Body)
	defer req.Body.Close()
	err := d.Decode(r)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) handleCreateComment(w http.ResponseWriter, r *http.Request) {
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

	d := &createCommentParams{}
	err = d.getParams(r)
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}

	comment := &model.Comment{
		PostID: d.PostID,
		Author: &model.User{
			ID:       user.ID,
			Nickname: user.Nickname,
		},
		Text: d.Text,
	}

	if strings.Trim(comment.Text, " ") == "" {
		s.error(w, http.StatusBadRequest, errors.New("no empty values"))
		return
	}

	err = s.store.Comment().Create(comment)
	if err != nil {
		s.error(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, map[string]interface{}{
		"comment": comment,
	})
}
