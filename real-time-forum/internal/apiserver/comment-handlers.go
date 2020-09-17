package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var tempID int = 25

func (s *server) makeComments() {
	s.comments = make(map[int][]*model.Comment)
	s.comments[1] = []*model.Comment{
		{
			ID: 1,
			Author: &model.User{
				ID:       3,
				Nickname: "nickfury",
			},
			Text: "Avengers Assemble",
		},
		{
			ID: 2,
			Author: &model.User{
				ID:       1,
				Nickname: "kusbek",
			},
			Text: "Debich",
		},
	}

	s.comments[2] = []*model.Comment{
		{
			ID: 3,
			Author: &model.User{
				ID:       4,
				Nickname: "gavnojui",
			},
			Text: "TEXT TEXT TEXT",
		},
		{
			ID: 4,
			Author: &model.User{
				ID:       1,
				Nickname: "kusbek",
			},
			Text: "Debich",
		},
		{
			ID: 5,
			Author: &model.User{
				ID:       1,
				Nickname: "kusbek",
			},
			Text: "Debich",
		},
	}
}

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

	res, _ := s.comments[params.postID]
	s.respond(w, http.StatusOK, map[string]interface{}{
		"comments": res,
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
		s.error(w, http.StatusUnauthorized, errors.New("No cookie"))
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
	res := &model.Comment{
		ID: tempID,
		Author: &model.User{
			ID:       user.ID,
			Nickname: user.Nickname,
		},
		Text: d.Text,
	}
	tempID++

	if c, ok := s.comments[d.PostID]; ok {
		c = append(c, res)
		s.comments[d.PostID] = c
	} else {
		s.comments[d.PostID] = make([]*model.Comment, 0)
		s.comments[d.PostID] = append(s.comments[d.PostID], res)
	}
	s.posts[d.PostID].Comments = len(s.comments[d.PostID])

	s.respond(w, http.StatusOK, map[string]interface{}{
		"comment": res,
	})
}
