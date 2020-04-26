package like

import (
	"DIV-01/new_forum/common"
	"database/sql"
	"fmt"
)

var (
	insertLike = `INSERT INTO likes ( user_id, post_id, like ) VALUES (?,?,?)`
	deleteLike = `delete from likes where likes.user_id = ? and likes.post_id = ?`
	checkLike  = `select count(*) as likes_count, coalesce(likes.like,0) from likes where likes.user_id = ? and likes.post_id = ?`
	updateLike = `update likes set like = ? where user_id = ? and post_id = ?`
)

type Like struct {
	Count  int `sql:"likes_count"`
	IsLike int `sql:"like"`
}

func likeExists(db *sql.DB, userID, postID int) (bool, int) {
	sqlRead := checkLike
	rows, err := db.Query(sqlRead, userID, postID)
	common.AbortOnError(err, "Failed to check if like exists")
	defer rows.Close()
	result := Like{}
	for rows.Next() {
		err = rows.Scan(&result.Count, &result.IsLike)
		common.AbortOnError(err, "Does not follow model: ")
	}

	if result.Count != 0 {
		return true, result.IsLike
	}

	return false, 0
}

func UpdateLike(db *sql.DB, userID, postID, islike int) {
	fmt.Println("Updating like")
	sqlStatement := updateLike
	stmt, err := db.Prepare(sqlStatement)
	common.AbortOnError(err, "wrong query")
	defer stmt.Close()
	_, err = stmt.Exec(islike, userID, postID)
	common.AbortOnError(err, "Failed to update like")
}

func InsertLike(db *sql.DB, userID, postID, islike int) {
	fmt.Println("Inserting like")
	sqlStatement := insertLike
	stmt, err := db.Prepare(sqlStatement)
	common.AbortOnError(err, "wrong query")
	defer stmt.Close()
	_, err = stmt.Exec(userID, postID, islike)
	common.AbortOnError(err, "Failed to insert like")
}

func DeleteLike(db *sql.DB, userID, postID int) {
	fmt.Println("Deleting like")
	sqlStatement := deleteLike
	stmt, err := db.Prepare(sqlStatement)
	common.AbortOnError(err, "wrong query")
	defer stmt.Close()
	_, err = stmt.Exec(userID, postID)
	common.AbortOnError(err, "Failed to Delete Like")
}
