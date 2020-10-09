package sqlstore

import (
	"DIV-01/real-time-forum/internal/model"
	"DIV-01/real-time-forum/internal/store"
)

//PostRepository ...
type PostRepository struct {
	store *Store
}

//Post ...
func (s *Store) Post() store.PostRepository {
	if s.postRepository != nil {
		return s.postRepository
	}
	s.postRepository = &PostRepository{
		store: s,
	}
	return s.postRepository
}

//GetAll ...
func (pr *PostRepository) GetAll() ([]*model.Post, error) {
	rows, err := pr.store.db.Query(`
		SELECT posts.id, users.id, users.nickname, posts.title, posts.post_text, posts.comments FROM posts
		LEFT JOIN users ON posts.author_id = users.id`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	posts := make([]*model.Post, 0)
	for rows.Next() {
		user := &model.User{}
		post := &model.Post{}
		err := rows.Scan(
			&post.ID,
			&user.ID,
			&user.Nickname,
			&post.Title,
			&post.Text,
			&post.Comments,
		)
		if err != nil {
			return nil, err
		}
		post.Author = user
		posts = append(posts, post)
	}
	return posts, nil
}

//Get ...
func (pr *PostRepository) Get(category string) ([]*model.Post, error) {
	rows, err := pr.store.db.Query(`
		SELECT posts.id, users.id, users.nickname, posts.title, posts.post_text, posts.comments FROM posts
		LEFT JOIN users ON posts.author_id = users.id
		WHERE posts.category = $1
		`, category)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	posts := make([]*model.Post, 0)
	for rows.Next() {
		user := &model.User{}
		post := &model.Post{}
		err := rows.Scan(
			&post.ID,
			&user.ID,
			&user.Nickname,
			&post.Title,
			&post.Text,
			&post.Comments,
		)
		if err != nil {
			return nil, err
		}
		post.Author = user
		posts = append(posts, post)
	}
	return posts, nil
}

//Create ...
func (pr *PostRepository) Create(post *model.Post) error {
	res, err := pr.store.db.Exec(
		`INSERT INTO posts (author_id, title, post_text, category) VALUES ($1, $2, $3, $4)`,
		post.Author.ID,
		post.Title,
		post.Text,
		post.Category,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	post.ID = int(id)
	return nil
}

//Delete ...
func (pr *PostRepository) Delete(id int) error {
	_, err := pr.store.db.Exec(`DELETE FROM posts WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
