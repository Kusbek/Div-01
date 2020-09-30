package model

import "time"

//Room ...
type Room struct {
	ID int `json:"id"`
}

//Message ...
type Message struct {
	Timestamp time.Time `json:"timestamp"`
	User      *User     `json:"user"`
	Text      string    `json:"text"`
}
