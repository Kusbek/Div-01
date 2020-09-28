package model

import "github.com/go-playground/validator"

//Post ...
type Post struct {
	ID       int    `json:"id,omitempty"`
	Author   *User  `json:"author" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Text     string `json:"text" validate:"required" validate:"contains=news|contains=game|contains=movie"`
	Comments int    `json:"comments"`
	Category string `json:"category" validate:"required"`
}

//Validate ...
func (p *Post) Validate() error {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(p)
	return err
}

//TestPost ...
func TestPost(userID int) *Post {
	return &Post{
		Author:   &User{ID: userID},
		Title:    "My Test Post",
		Text:     "This is a test post",
		Category: "test",
	}
}
