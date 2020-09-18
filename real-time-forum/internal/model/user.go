package model

import "github.com/go-playground/validator"

//User ...
type User struct {
	ID        int    `json:"id,omitempty" sql:"id"`
	Nickname  string `json:"nickname,omitempty" validate:"required,min=5,max=20"`
	Age       uint8  `json:"age,omitempty" validate:"gte=0,lte=120"`
	Gender    string `json:"gender,omitempty" validate:"contains=male|contains=female"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty" validate:"email"`
	Password  string `json:"-" validate:"min=5"`
}

//Validate ...
func (u *User) Validate() error {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(u)
	return err
}

var count int = 0

//TestUser ...
func TestUser(creds string, password string) *User {
	count++
	return &User{
		ID:        count,
		Nickname:  creds,
		Email:     "testemail@gmail.com",
		Gender:    "male",
		FirstName: "First Name",
		LastName:  "Last Name",
		Age:       20,
	}
}
