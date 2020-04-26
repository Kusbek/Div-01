package post

import (
	sqlite "DIV-01/new_forum/sqlite"
	"DIV-01/new_forum/user"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func HandlePosts(w http.ResponseWriter, req *http.Request) {
	db := sqlite.GetDB()
	switch req.Method {
	case "GET":
		HandleGetPosts(w, req)
	case "POST":
		u, err := user.Authenticate(req)
		if err != nil {
			w.Write([]byte("Failed to autheticate"))
			return
		}
		game, err := strconv.ParseBool(req.URL.Query().Get("game"))
		cosplay, err := strconv.ParseBool(req.URL.Query().Get("cosplay"))
		movie, err := strconv.ParseBool(req.URL.Query().Get("movie"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(game, cosplay, movie)
		post := &Post{}
		err = json.NewDecoder(req.Body).Decode(post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = CreatePost(db, u.ID, game, cosplay, movie, post.Title, post.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		w.Write([]byte("Takogo methoda net"))
	}
}

func HandleGetPosts(w http.ResponseWriter, req *http.Request) {
	db := sqlite.GetDB()
	categories := map[string]string{
		"all":     "all",
		"created": "created",
		"liked":   "liked",
		"game":    "game",
		"movie":   "movie",
		"cosplay": "cosplay",
	}
	category, ok := categories[req.URL.Query().Get("category")]
	if !ok {
		category = "all"
	}
	var posts *Posts
	// fmt.Println(category)
	if category == "created" || category == "liked" {
		u, err := user.Authenticate(req)
		if err != nil {
			w.Write([]byte("Failed to autheticate"))
			return
		}
		posts = &Posts{Posts: GetUserPosts(db, category, u.ID)}
	} else {
		posts = &Posts{Posts: GetPosts(db, category)}
	}

	tmpl := template.Must(template.ParseFiles("./post/posts.html"))
	tmpl.ExecuteTemplate(w, "posts.html", posts)
}
