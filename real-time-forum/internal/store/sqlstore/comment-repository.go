package sqlstore

import (
	"DIV-01/real-time-forum/internal/model"
	"DIV-01/real-time-forum/internal/store"
	"context"
)

//CommentRepository ...
type CommentRepository struct {
	store *Store
}

//Comment ...
func (s *Store) Comment() store.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}
	s.commentRepository = &CommentRepository{
		store: s,
	}
	return s.commentRepository
}

//Create ...
func (cr *CommentRepository) Create(comment *model.Comment) error {
	ctx := context.Background()
	tx, err := cr.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, `
		INSERT INTO comments (author_id, post_id, comment_text) values ($1, $2, $3)
	`, comment.Author.ID, comment.PostID, comment.Text)
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	comment.ID = int(id)

	_, err = tx.ExecContext(ctx, `
	UPDATE posts SET comments = comments + 1 WHERE id = $1
	`, comment.PostID)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

//Get ...
func (cr *CommentRepository) Get(postID int) ([]*model.Comment, error) {
	res, err := cr.store.db.Query(`
		SELECT comments.id, post_id, users.id, users.nickname, comment_text from comments
		LEFT JOIN users ON users.id = author_id
		WHERE post_id = $1
	`, postID)
	defer res.Close()
	if err != nil {
		return nil, err
	}
	comments := make([]*model.Comment, 0)
	for res.Next() {
		author := &model.User{}
		comment := &model.Comment{}
		err = res.Scan(
			&comment.ID,
			&comment.PostID,
			&author.ID,
			&author.Nickname,
			&comment.Text,
		)
		if err != nil {
			return nil, err
		}

		comment.Author = author
		comments = append(comments, comment)
	}

	return comments, nil
}
