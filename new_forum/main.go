package main

import (
	"DIV-01/new_forum/comment"
	"DIV-01/new_forum/like"
	"DIV-01/new_forum/post"
	sqlite "DIV-01/new_forum/sqlite"
	"DIV-01/new_forum/user"
	"html/template"
	"log"
	"net/http"
)

func main() {
	sqlite.Initialize()
	db := sqlite.GetDB()
	defer db.Close()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	user.NewCookies()

	http.HandleFunc("/", HandleMain)
	//Posts handling
	http.HandleFunc("/post", post.HandlePosts)
	http.HandleFunc("/det_post", post.HandleDetailedPosts)
	//Likes Handling
	http.HandleFunc("/like", like.HandleLikes)

	//Comment Handling
	http.HandleFunc("/comment", comment.HandleComments)

	//Comment Handling
	http.HandleFunc("/signup", user.HandleUserCreate)
	http.HandleFunc("/signin", user.HandleUserLogin)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func HandleMain(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
