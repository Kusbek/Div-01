package model

//Post ...
type Post struct {
	ID       int    `json:"id,omitempty"`
	Author   *User  `json:"author"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Comments int    `json:"comments"`
}
