package model

//Comment ...
type Comment struct {
	ID     int    `json:"id"`
	Author *User  `json:"author" validate:"required"`
	Text   string `json:"text"`
}
