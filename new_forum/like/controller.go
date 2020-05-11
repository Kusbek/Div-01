package like

import (
	sqlite "DIV-01/new_forum/sqlite"
	"DIV-01/new_forum/user"
	"encoding/json"
	"fmt"
	"net/http"
)

type LikeRequest struct {
	PostID int `json:"post_id"`
	IsLike int `json:"is_like"`
}

func HandleLikes(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling like")
	l := &LikeRequest{}
	err := json.NewDecoder(req.Body).Decode(l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := user.Authenticate(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	db := sqlite.GetDB()
	switch req.Method {
	case "POST":
		fmt.Println("HANDLING LIKE POST")
		if exists, like := likeExists(db, u.ID, l.PostID); exists {
			if l.IsLike == like {
				fmt.Println("Deleting Like")
				DeleteLike(db, u.ID, l.PostID)
			} else {
				fmt.Println("Updating Like")
				UpdateLike(db, u.ID, l.PostID, l.IsLike)
			}
			return
		}
		fmt.Println("Inserting Like")
		InsertLike(db, u.ID, l.PostID, l.IsLike)
	default:
		w.Write([]byte("Takogo methoda net"))
	}
}
