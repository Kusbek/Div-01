package user

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `sql:"id" json:"id,omitempty"`
	Username string `sql:"username" json:"username"`
	Email    string `sql:"email" json:"email"`
	Password string `sql:"password" json:"password"`
}

type Credentials struct {
	Credentials string `json:"credentials"`
	Password    string `json:"password"`
}

const (
	sqlInsertUser = `INSERT INTO users(username, email, password)VALUES (?,?,?)`
	sqlQueryUser  = `SELECT * FROM users WHERE username=? or email=?`
)

func InsertUser(db *sql.DB, username, email, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare(sqlInsertUser)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, email, string(hashedPass))
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *sql.DB, creds, password string) (*User, error) {
	user := &User{}
	result := db.QueryRow(sqlQueryUser, creds, creds)
	err := result.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
