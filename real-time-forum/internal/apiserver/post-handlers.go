package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"errors"
	"fmt"
	"net/http"
)

var posts map[int]*model.Post = make(map[int]*model.Post)

func init() {
	posts[1] = &model.Post{
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
	posts[2] = &model.Post{
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
	posts[3] = &model.Post{
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

	for _, v := range posts {
		res = append(res, v)
	}

	s.respond(w, http.StatusOK, map[string]interface{}{
		"posts": res,
	})
}
func (s *server) handleCreatePost(w http.ResponseWriter, r *http.Request) {}
