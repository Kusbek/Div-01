package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"DIV-01/real-time-forum/internal/session"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Server ...
type Server interface {
	ListenAndServe(string) error
}

type server struct {
	mux     *http.ServeMux
	cookies session.Cookie
}

func newServer() Server {
	s := &server{
		cookies: session.New(),
	}
	s.newMux()
	return s
}

func (s *server) newMux() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web")))
	mux.HandleFunc("/auth", s.authHandler)
	mux.HandleFunc("/signup", s.signUpHandler)
	mux.HandleFunc("/signin", s.signInHandler)
	mux.HandleFunc("/signout", s.signOutHandler)
	mux.HandleFunc("/post", s.handlePosts)
	s.mux = mux
}

func (s *server) ListenAndServe(addr string) error {
	server := http.Server{
		Addr:    addr,
		Handler: s.mux,
	}
	fmt.Println("starting server at", addr)
	return server.ListenAndServe()
}

func (s *server) error(w http.ResponseWriter, code int, err error) {
	fmt.Println("Error:", err)
	s.respond(w, code, map[string]interface{}{
		"error": err.Error(),
	})
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func (s *server) setCookies(w http.ResponseWriter, user *model.User) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   s.cookies.Insert(user),
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}
