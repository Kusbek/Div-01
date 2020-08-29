package apiserver

import (
	"errors"
	"net/http"
)

func (s *server) signInHandler(w http.ResponseWriter, r *http.Request) {
	s.error(w, http.StatusUnauthorized, errors.New("Unothorized"))
	// s.respond(w, http.StatusOK, map[string]string{
	// 	"data": "It's a me, Mario!!!",
	// })
}
