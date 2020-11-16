package apiserver

import (
	"fmt"
	"net/http"
	"os"
)

//Options ...
type Options struct {
	Address string
}

//Run ....
func (s *server) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := http.Server{
		Addr:    ":" + port,
		Handler: s.mux,
	}
	fmt.Println("starting server at", server.Addr)
	return server.ListenAndServe()
}
