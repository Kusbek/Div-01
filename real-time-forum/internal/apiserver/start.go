package apiserver

import (
	"fmt"
	"net/http"
)

//Options ...
type Options struct {
	Address string
}

//Run ....
func (s *server) Run() error {
	server := http.Server{
		Addr:    ":8082",
		Handler: s.mux,
	}
	fmt.Println("starting server at", server.Addr)
	return server.ListenAndServe()
}
