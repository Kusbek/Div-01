package user

import (
	sqlite "DIV-01/new_forum/sqlite"
	"encoding/json"
	"net/http"
	"time"
)

func HandleUserCreate(w http.ResponseWriter, req *http.Request) {
	db := sqlite.GetDB()
	switch req.Method {
	case "GET":
		w.Write([]byte("GET methoda net"))
		return
	case "POST":
		user := &User{}
		err := json.NewDecoder(req.Body).Decode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = InsertUser(db, user.Username, user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("Created"))
	default:
		w.Write([]byte("Takogo methoda net"))
	}
}

func HandleUserLogin(w http.ResponseWriter, req *http.Request) {
	db := sqlite.GetDB()
	switch req.Method {
	case "GET":
		w.Write([]byte("GET methoda net"))
		return
	case "POST":
		userCreds := &Credentials{}
		err := json.NewDecoder(req.Body).Decode(userCreds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u, err := GetUser(db, userCreds.Credentials, userCreds.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		c := GetCookies()
		uuid := c.Insert(u)

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   uuid,
			Expires: time.Now().Add(COOKIEEXPIRETIME * time.Second),
			Path:    "/",
		})
	default:
		w.Write([]byte("Takogo methoda net"))
	}
}
