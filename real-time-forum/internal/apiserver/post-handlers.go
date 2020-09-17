package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	temoPostID = 4
)

func (s *server) makePosts() {
	s.posts = make(map[int]*model.Post)
	s.posts[1] = &model.Post{
		ID:    1,
		Title: "TITLE HEADING 1",
		Text: `Some text..
		Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.`,
		Comments: 2,
		Author: &model.User{
			ID:       1,
			Nickname: "kusbek",
		},
	}
	s.posts[2] = &model.Post{
		ID:    2,
		Title: "TITLE HEADING 1",
		Text: `Some text..
		Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.`,
		Comments: 3,
		Author: &model.User{
			ID:       2,
			Nickname: "postAuthorNickname",
		},
	}
	s.posts[3] = &model.Post{
		ID:    3,
		Title: "TITLE HEADING 1",
		Text: `Some text..
		Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.`,
		Comments: 0,
		Author: &model.User{
			ID:       1,
			Nickname: "kusbek",
		},
	}
}
func (s *server) handlePosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Posts")
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
	res := make([]*model.Post, 0)

	for _, v := range s.posts {
		res = append(res, v)
	}

	s.respond(w, http.StatusOK, map[string]interface{}{
		"posts": res,
	})
}

//CreatePostParams ...
type CreatePostParams struct {
	Title string `json:"title"`
	Text  string `json:"text"`
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

	d := &createCommentParams{}
	err = d.getParams(r)
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}
	res := &model.Post{
		ID: temoPostID,
		Author: &model.User{
			ID:       user.ID,
			Nickname: user.Nickname,
		},
		Comments: func() int {
			s.comments[temoPostID] = make([]*model.Comment, 0)
			return len(s.comments[temoPostID])
		}(),
		Text: d.Text,
	}

	s.posts[temoPostID] = res
	temoPostID++
	s.respond(w, http.StatusOK, map[string]interface{}{
		"post": res,
	})
}
