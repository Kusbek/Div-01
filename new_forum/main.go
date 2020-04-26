package main

import (
	"DIV-01/new_forum/comment"
	"DIV-01/new_forum/like"
	"DIV-01/new_forum/post"
	sqlite "DIV-01/new_forum/sqlite"
	"DIV-01/new_forum/user"
	"net/http"
)

func main() {
	sqlite.Initialize()
	db := sqlite.GetDB()
	defer db.Close()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	user.NewCookies()

	//Posts handling
	http.HandleFunc("/", post.HandlePosts)

	//Likes Handling
	http.HandleFunc("/like", like.HandleLikes)

	//Comment Handling
	http.HandleFunc("/comment", comment.HandleComments)

	//Comment Handling
	http.HandleFunc("/user/create", user.HandleUserCreate)
	http.HandleFunc("/user/login", user.HandleUserLogin)

	// fmt.Println(post.GetPosts(db, "all"))

	http.ListenAndServe(":8080", nil)
}
