package post

import (
	"DIV-01/new_forum/common"
	"database/sql"
)

// type Post struct {
// 	ID       int    `sql:"id"`
// 	userID   int    `sql:"user_id"`
// 	Title    string `sql:"title"`
// 	Text     string `sql:"text"`
// 	Game     int    `sql:"game"`
// 	Movie    int    `sql:"movie"`
// 	Cosplay  int    `sql:"cosplay"`
// 	Likes    int    `sql:"likes"`
// 	Comments int    `sql:"comments"`
// }

type Post struct {
	ID       int    `sql:"id" json:"id,omitempty"`
	Title    string `sql:"title" json:"title"`
	Text     string `sql:"text" json:"text"`
	Author   string `sql:"username" json:"author,omitempty"`
	Likes    int    `sql:"likes_count" json:"likes,omitempty"`
	Dislikes int    `sql:"dislikes_count" json:"dislikes,omitempty"`
	Comments int    `sql:"comments_count" json:"comments,omitempty"`
}

type Posts struct {
	Posts []Post
}

var (
	SinglePost = `
	select posts.id, posts.title, posts.text, users.username, coalesce(likers.likes_count,0), coalesce(dislikers.dislikes_count,0), coalesce(commenters.comments_count,0) from posts
    left join (select likes.post_id, count(likes.user_id) as likes_count from likes where (likes.like=1) group by likes.post_id) as likers on likers.post_id = posts.id
    left join (select likes.post_id, count(likes.user_id) as dislikes_count from likes where (likes.like=0) group by likes.post_id) as dislikers on dislikers.post_id = posts.id
    left join (select comments.post_id, count(comments.user_id) as comments_count from comments group by comments.post_id) as commenters on commenters.post_id = posts.id
	left join users on users.id = posts.user_id
	where posts.id = ?
	`
	AllPosts = `
	select posts.id, posts.title, posts.text, users.username, coalesce(likers.likes_count,0), coalesce(dislikers.dislikes_count,0), coalesce(commenters.comments_count,0) from posts
    left join (select likes.post_id, count(likes.user_id) as likes_count from likes where (likes.like=1) group by likes.post_id) as likers on likers.post_id = posts.id
    left join (select likes.post_id, count(likes.user_id) as dislikes_count from likes where (likes.like=0) group by likes.post_id) as dislikers on dislikers.post_id = posts.id
    left join (select comments.post_id, count(comments.user_id) as comments_count from comments group by comments.post_id) as commenters on commenters.post_id = posts.id
    left join users on users.id = posts.user_id
	`
	CreatedPosts = `
	select posts.id, posts.title, posts.text, users.username, coalesce(likers.likes_count,0), coalesce(dislikers.dislikes_count,0), coalesce(commenters.comments_count,0)
	from posts
		left join (select likes.post_id, count(likes.user_id) as likes_count from likes where (likes.like=1) group by likes.post_id) as likers on likers.post_id = posts.id
		left join (select likes.post_id, count(likes.user_id) as dislikes_count from likes where (likes.like=0) group by likes.post_id) as dislikers on dislikers.post_id = posts.id
		left join (select comments.post_id, count(comments.user_id) as comments_count from comments group by comments.post_id) as commenters on commenters.post_id = posts.id
		left join users on users.id = posts.user_id
	where (users.id = ?)
	`
	LikedPosts = `
	select posts.id, posts.title, posts.text, users.username, coalesce(likers.likes_count,0), coalesce(dislikers.dislikes_count,0), coalesce(commenters.comments_count,0)
	from posts
    left join (select likes.post_id, count(likes.user_id) as likes_count from likes where (likes.like=1) group by likes.post_id) as likers on likers.post_id = posts.id
	left join (select likes.post_id, count(likes.user_id) as dislikes_count from likes where (likes.like=0) group by likes.post_id) as dislikers on dislikers.post_id = posts.id
	left join (select comments.post_id, count(comments.user_id) as comments_count from comments group by comments.post_id) as commenters on commenters.post_id = posts.id
    left join users on users.id = posts.user_id
	join (select likes.post_id from likes where (user_id=$1)) as temp on temp.post_id = posts.id
	`
	GamePosts = `
	select posts.id, posts.title, posts.text, users.username, coalesce(likers.likes_count,0), coalesce(dislikers.dislikes_count,0), coalesce(commenters.comments_count,0)
	from posts
    left join (select likes.post_id, count(likes.user_id) as likes_count from likes where (likes.like=1) group by likes.post_id) as likers on likers.post_id = posts.id
	left join (select likes.post_id, count(likes.user_id) as dislikes_count from likes where (likes.like=0) group by likes.post_id) as dislikers on dislikers.post_id = posts.id
	left join (select comments.post_id, count(comments.user_id) as comments_count from comments group by comments.post_id) as commenters on commenters.post_id = posts.id
    left join users on users.id = posts.user_id
	where (posts.game = 1)
	`

	CosplayPosts = `
	select posts.id, posts.title, posts.text, users.username, coalesce(likers.likes_count,0), coalesce(dislikers.dislikes_count,0), coalesce(commenters.comments_count,0)
	from posts
    left join (select likes.post_id, count(likes.user_id) as likes_count from likes where (likes.like=1) group by likes.post_id) as likers on likers.post_id = posts.id
	left join (select likes.post_id, count(likes.user_id) as dislikes_count from likes where (likes.like=0) group by likes.post_id) as dislikers on dislikers.post_id = posts.id
	left join (select comments.post_id, count(comments.user_id) as comments_count from comments group by comments.post_id) as commenters on commenters.post_id = posts.id
    left join users on users.id = posts.user_id
	where (posts.cosplay = 1)
	`

	MoviePosts = `
	select posts.id, posts.title, posts.text, users.username, coalesce(likers.likes_count,0), coalesce(dislikers.dislikes_count,0), coalesce(commenters.comments_count,0)
	from posts
    left join (select likes.post_id, count(likes.user_id) as likes_count from likes where (likes.like=1) group by likes.post_id) as likers on likers.post_id = posts.id
	left join (select likes.post_id, count(likes.user_id) as dislikes_count from likes where (likes.like=0) group by likes.post_id) as dislikers on dislikers.post_id = posts.id
	left join (select comments.post_id, count(comments.user_id) as comments_count from comments group by comments.post_id) as commenters on commenters.post_id = posts.id
    left join users on users.id = posts.user_id
	where (posts.movie = 1)
	`

	insertPost = `INSERT INTO posts ( user_id, title, text, game, movie, cosplay) VALUES (?,?,?,?,?,?);`
)

func getQuery(category string) string {
	if category == "created" {
		return CreatedPosts
	}

	if category == "liked" {
		return LikedPosts
	}

	if category == "game" {
		return GamePosts
	}

	if category == "cosplay" {
		return CosplayPosts
	}

	if category == "movie" {
		return MoviePosts
	}

	return AllPosts
}

func GetUserPosts(db *sql.DB, category string, userID int) []Post {
	sqlRead := getQuery(category)
	rows, err := db.Query(sqlRead, userID)
	common.AbortOnError(err, "Error when reading from database")
	defer rows.Close()
	result := []Post{}

	for rows.Next() {
		item := Post{}
		err = rows.Scan(&item.ID, &item.Title, &item.Text, &item.Author, &item.Likes, &item.Dislikes, &item.Comments)
		common.AbortOnError(err, "Does not follow model: ")
		result = append(result, item)
	}
	return result
}

func GetPosts(db *sql.DB, category string) []Post {
	sqlRead := getQuery(category)
	rows, err := db.Query(sqlRead)
	common.AbortOnError(err, "Error when reading from database")
	defer rows.Close()
	result := []Post{}

	for rows.Next() {
		item := Post{}
		err = rows.Scan(&item.ID, &item.Title, &item.Text, &item.Author, &item.Likes, &item.Dislikes, &item.Comments)
		common.AbortOnError(err, "Does not follow model: ")
		result = append(result, item)
	}
	return result
}

func GetPost(db *sql.DB, id int) Post {
	sqlRead := SinglePost
	rows, err := db.Query(sqlRead, id)
	common.AbortOnError(err, "Error when reading from database")
	defer rows.Close()
	result := Post{}
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Title, &result.Text, &result.Author, &result.Likes, &result.Dislikes, &result.Comments)
		common.AbortOnError(err, "Does not follow model: ")
	}

	return result
}

func CreatePost(db *sql.DB, userID int, Game, Movie, Cosplay bool, title, text string) error {
	sqlAdditem := insertPost
	stmt, err := db.Prepare(sqlAdditem)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, title, text, Game, Movie, Cosplay)
	if err != nil {
		return err
	}
	return nil
}

// func CreatePostTable(db *sql.DB) {
// 	sqlTable := `
// 	CREATE TABLE posts(
// 		Id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
// 		Title TEXT,
// 		Text TEXT,
// 		CreateDate DATETIME
// 	);
// 	`

// 	_, err := db.Exec(sqlTable)
// 	common.AbortOnError(err, "Failed to create table")
// }
