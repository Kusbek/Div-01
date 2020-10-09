package sqlstore

import (
	"DIV-01/real-time-forum/internal/model"
	"DIV-01/real-time-forum/internal/store"
	"database/sql"
)

//User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create ...
func (ur *UserRepository) Create(user *model.User) error {
	err := user.EncryptPassword()
	if err != nil {
		return err
	}

	res, err := ur.store.db.Exec(
		`INSERT INTO users (nickname, email, gender, first_name, last_name, password)
		 VALUES($1,$2,$3,$4,$5,$6)`,
		user.Nickname,
		user.Email,
		user.Gender,
		user.FirstName,
		user.LastName,
		user.Password,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

//Delete ...
func (ur *UserRepository) Delete(id int) error {
	_, err := ur.store.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

//Find ...
func (ur *UserRepository) Find(nickmail string) (*model.User, error) {
	user := &model.User{}
	err := ur.store.db.QueryRow(
		`SELECT * FROM users WHERE nickname=$1 OR email=$1`,
		nickmail).
		Scan(
			&user.ID,
			&user.Nickname,
			&user.Email,
			&user.Gender,
			&user.FirstName,
			&user.LastName,
			&user.Password,
		)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//GetByID ...
func (ur *UserRepository) GetByID(id int) (*model.User, error) {
	user := &model.User{}
	err := ur.store.db.QueryRow(
		`SELECT * FROM users WHERE id=$1`,
		id).
		Scan(
			&user.ID,
			&user.Nickname,
			&user.Email,
			&user.Gender,
			&user.FirstName,
			&user.LastName,
			&user.Password,
		)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//Exists ...
func (ur *UserRepository) Exists(nickname, email string) (bool, error) {
	var exists bool
	err := ur.store.db.QueryRow(
		`SELECT 1 FROM users WHERE nickname IN ($1,$2) OR email IN ($1,$2)`,
		nickname, email).
		Scan(
			&exists,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return exists, nil
}
