package comment

import (
	sqlite "DIV-01/new_forum/sqlite"
	"DIV-01/new_forum/user"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HandleComments(w http.ResponseWriter, req *http.Request) {
	db := sqlite.GetDB()

	switch req.Method {
	case "GET":
		postID, err := strconv.Atoi(req.URL.Query().Get("post_id"))
		if err != nil {
			http.Error(w, "post_id should be integer", http.StatusBadRequest)
			return
		}
		comments := GetCommentsForPost(db, postID)
		for _, comment := range comments {
			fmt.Println(comment)
		}
	case "POST":
		u, err := user.Authenticate(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		comment := &Comment{}
		err = json.NewDecoder(req.Body).Decode(comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = InsertComment(db, u.ID, comment.PostID, comment.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		w.Write([]byte("Takogo methoda net"))
	}
}
