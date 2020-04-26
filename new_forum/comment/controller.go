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
	postID, err := strconv.Atoi(req.URL.Query().Get("post_id"))
	if err != nil {
		w.Write([]byte("post_id should be integer"))
	}
	switch req.Method {
	case "GET":
		comments := GetCommentsForPost(db, postID)
		for _, comment := range comments {
			fmt.Println(comment)
		}
	case "POST":
		u, err := user.Authenticate(req)
		if err != nil {
			w.Write([]byte("Failed to autheticate"))
		}
		comment := &Comment{}
		err = json.NewDecoder(req.Body).Decode(comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = InsertComment(db, u.ID, postID, comment.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		w.Write([]byte("Takogo methoda net"))
	}
}
