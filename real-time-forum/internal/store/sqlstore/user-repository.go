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
		`INSERT INTO users (nickname, email, gender, first_name, last_name, password, age)
		 VALUES($1,$2,$3,$4,$5,$6)`,
		user.Nickname,
		user.Email,
		user.Gender,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Age,
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
			&user.Age,
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
			&user.Age,
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

//GetAll ...
func (ur *UserRepository) GetAll() ([]*model.User, error) {
	rows, err := ur.store.db.Query(`SELECT id, nickname FROM users`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, 0)
	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.Nickname)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

//GetUsers ...
func (ur *UserRepository) GetUsers(id int) ([]*model.User, error) {
	users := make([]*model.User, 0)
	rows, err := ur.store.db.Query(`
	SELECT users.id, users.nickname from users 
	JOIN 
		(SELECT user_id, room_id FROM room_participants WHERE room_id IN 
			(SELECT room_id FROM room_participants WHERE user_id = $1 GROUP BY room_id) AND user_id != $1) 
	as open_rooms ON users.id = open_rooms.user_id 
	LEFT JOIN
		(SELEct * FROM (SELECT room_id,message_timestamp from messages WHERE room_id in 
			(SELECT room_id FROM room_participants WHERE room_id IN 
				(SELECT room_id FROM room_participants WHERE user_id = $1 GROUP BY room_id) AND user_id != $1) ORDER BY message_timestamp DESC) 
				GROUP BY room_id) 
	as temp_table ON open_rooms.room_id = temp_table.room_id
	WHERE id != $1 ORDER BY temp_table.message_timestamp DESC
	`, id)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.Nickname)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	rows, err = ur.store.db.Query(`
	SELECT id, nickname FROM users WHERE id NOT IN (SELECT user_id FROM room_participants WHERE room_id IN (SELECT room_id FROM room_participants WHERE user_id = $1 GROUP BY room_id) AND user_id != $1) AND id != $1 ORDER BY nickname ASC;
	`, id)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.Nickname)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
