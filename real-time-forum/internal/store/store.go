package store

import "DIV-01/real-time-forum/internal/model"

//Store ...
type Store interface {
	Close()
	User() UserRepository
	Post() PostRepository
	Comment() CommentRepository
	Room() RoomRepository
}

//RoomRepository ...
type RoomRepository interface {
	GetRoom(userID1, userID2 int) (*model.Room, error)
	CreateRoom(userID1, userID2 int) (*model.Room, error)
	DeleteRoom(id int) error
	NewMessage(roomID int, m *model.Message) error
	GetMessages(roomID int, from int) ([]*model.Message, error)
}

//CommentRepository ...
type CommentRepository interface {
	Get(postID int) ([]*model.Comment, error)
	Create(comment *model.Comment) error
}

//PostRepository ...
type PostRepository interface {
	GetAll() ([]*model.Post, error)
	Create(post *model.Post) error
	Delete(id int) error
	Get(category string) ([]*model.Post, error)
}

//UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	Find(nickmail string) (*model.User, error)
	Delete(id int) error
	GetByID(id int) (*model.User, error)
	Exists(nickname, email string) (bool, error)
}
