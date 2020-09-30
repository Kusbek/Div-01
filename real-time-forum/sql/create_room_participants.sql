CREATE TABLE room_participants (
  user_id INTEGER NOT NULL,
  room_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (room_id) REFERENCES rooms (id)
);