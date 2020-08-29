package apiserver

import (
	"DIV-01/real-time-forum/internal/model"
	"encoding/json"
	"net/http"
)

func (s *server) signUpHandler(w http.ResponseWriter, r *http.Request) {
	type body struct {
		Nickname  string `json:"nickname,omitempty"`
		Age       uint8  `json:"age,omitempty"`
		Gender    string `json:"gender,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Email     string `json:"email,omitempty"`
		Password  string `json:"password"`
	}

	b := &body{}
	err := json.NewDecoder(r.Body).Decode(b)
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}

	user := &model.User{
		Nickname:  b.Nickname,
		Age:       b.Age,
		Gender:    b.Gender,
		FirstName: b.FirstName,
		LastName:  b.LastName,
		Email:     b.Email,
		Password:  b.Password,
	}

	err = user.Validate()
	if err != nil {
		s.error(w, http.StatusBadRequest, err)
		return
	}

	// s.error(w, http.StatusUnauthorized, errors.New("Unothorized"))
	s.respond(w, http.StatusOK, map[string]string{
		"data": "It's a me, Mario!!!",
	})
}
