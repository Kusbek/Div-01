package comment

import (
	"DIV-01/new_forum/common"
	"database/sql"
)

type Comment struct {
	ID     int    `sql:"id" json:"id,omitempty"`
	Author string `sql:"username" json:"username,omitempty"`
	Text   string `sql:"text" json:"text,omitempty"`
}

var (
	commentsSQLQuery = `select comments.id, users.username, comments.text from comments
	left join users on users.id = comments.user_id
	where comments.post_id = ?;`

	commentSQLInsert = `INSERT INTO comments ( user_id, post_id, text ) VALUES (?,?,?);`
)

func GetCommentsForPost(db *sql.DB, postID int) []Comment {
	sqlRead := commentsSQLQuery
	rows, err := db.Query(sqlRead, postID)
	common.AbortOnError(err, "Error when reading from database")
	defer rows.Close()
	result := []Comment{}

	for rows.Next() {
		item := Comment{}
		err = rows.Scan(&item.ID, &item.Text, &item.Author)
		common.AbortOnError(err, "Does not follow model: ")
		result = append(result, item)
	}
	return result
}

func InsertComment(db *sql.DB, userID, postID int, text string) error {
	sqlInsert := commentSQLInsert
	stmt, err := db.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	common.AbortOnError(err, "Error when reading from database")
	defer stmt.Close()
	_, err = stmt.Exec(userID, postID, text)
	if err != nil {
		return err
	}
	common.AbortOnError(err, "Failed to insert comment")
	return nil
}
