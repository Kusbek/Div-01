package store

import "DIV-01/real-time-forum/internal/model"

//Store ...
type Store interface {
	Close()
	User() UserRepository
	Post() PostRepository
}

//PostRepository ...
type PostRepository interface {
	GetAll() ([]*model.Post, error)
	Create(post *model.Post) error
	Delete(id int) error
}

//UserRepository ...
type UserRepository interface {
	Create(user *model.User) error
	Find(nickmail string) (*model.User, error)
	Delete(id int) error
	GetByID(id int) (*model.User, error)
}
