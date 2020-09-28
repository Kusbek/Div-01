package model

import (
	"errors"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

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

func encrypt(str string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

//EncryptPassword ...
func (u *User) EncryptPassword() error {
	encrypted, err := encrypt(u.Password)
	if err != nil {
		return err
	}

	u.Password = encrypted
	return nil
}

//ComparePasswords ...
func (u *User) ComparePasswords(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.New("Passwords are not equal")
	}
	return nil
}

//Validate ...
func (u *User) Validate() error {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(u)
	return err
}

var count int = 0

//TestUser ...
func TestUser(nick string, password string) *User {
	count++
	return &User{
		ID:        count,
		Nickname:  nick,
		Email:     nick + "@gmail.com",
		Gender:    "male",
		FirstName: "First Name",
		LastName:  "Last Name",
		Age:       20,
		Password:  password,
	}
}
