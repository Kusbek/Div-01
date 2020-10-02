package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"DIV-01/real-time-forum/internal/session"
	"DIV-01/real-time-forum/internal/store"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

//Server ...
type Server interface {
	Run() error
}

type server struct {
	mux             *http.ServeMux
	store           store.Store
	cookies         session.Cookie
	rooms           map[int]*roomManager
	guests          []*guest
	deleteGuestChan chan *guest
	mu              *sync.Mutex
}

//NewServer ...
func NewServer(st store.Store) Server {
	s := &server{
		cookies:         session.New(),
		store:           st,
		guests:          make([]*guest, 0),
		deleteGuestChan: make(chan *guest, 10),
		rooms:           make(map[int]*roomManager),
		mu:              &sync.Mutex{},
	}

	s.newMux()
	go s.monitorDeleteGuestChan()
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
	mux.HandleFunc("/comment", s.handleComments)
	mux.HandleFunc("/chat", s.chatWsHandler)
	mux.HandleFunc("/room", s.roomHandler)
	mux.HandleFunc("/message", s.messageWsHandler)
	s.mux = mux
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
