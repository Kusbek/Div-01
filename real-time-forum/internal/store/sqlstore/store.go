package sqlstore

import (
	"DIV-01/real-time-forum/internal/store"
	"database/sql"

	//go-sqlite3 is important
	_ "github.com/mattn/go-sqlite3"
)

//Options ...
type Options struct {
	Address string
}

//Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	postRepository *PostRepository
}

//Start ...
func Start(opts *Options) (store.Store, error) {
	db, err := sql.Open("sqlite3", opts.Address)
	if err != nil {
		return nil, err
	}

	s := &Store{
		db: db,
	}
	return s, nil
}

//Close ...
func (s *Store) Close() {
	s.db.Close()
}
