CREATE TABLE messages (
  room_id INTEGER NOT NULL,
  author_id INTEGER NOT NULL,
  message_text VARCHAR,
  message_timestamp DATETIME,
  FOREIGN KEY (author_id) REFERENCES users (id),
  FOREIGN KEY (room_id) REFERENCES rooms (id)
);