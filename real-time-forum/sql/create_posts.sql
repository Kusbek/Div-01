CREATE TABLE posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  author_id INTEGER NOT NULL,
  title VARCHAR NOT NULL,
  post_text VARCHAR,
  category VARCHAR,
  comments INTEGER DEFAULT 0,
  FOREIGN KEY (author_id) REFERENCES users (id)
 );