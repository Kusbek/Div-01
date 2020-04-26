package like

import (
	sqlite "DIV-01/new_forum/sqlite"
	"DIV-01/new_forum/user"
	"fmt"
	"net/http"
	"strconv"
)

func HandleLikes(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling like")
	u, err := user.Authenticate(req)
	if err != nil {
		w.Write([]byte("Failed to autheticate"))
	}
	postID, err := strconv.Atoi(req.URL.Query().Get("post_id"))
	if err != nil {
		w.Write([]byte("post_id should be integer"))
	}
	isLike, err := strconv.Atoi(req.URL.Query().Get("is_like"))
	if err != nil {
		w.Write([]byte("is_like should be integer"))
	}
	db := sqlite.GetDB()
	switch req.Method {
	case "POST":
		if exists, like := likeExists(db, u.ID, postID); exists {
			if isLike == like {
				DeleteLike(db, u.ID, postID)
			} else {
				UpdateLike(db, u.ID, postID, isLike)
			}
			return
		}
		InsertLike(db, u.ID, postID, isLike)
	default:
		w.Write([]byte("Takogo methoda net"))
	}
}
