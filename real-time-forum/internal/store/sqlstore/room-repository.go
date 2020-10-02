package sqlstore

import (
	"DIV-01/real-time-forum/internal/model"
	"DIV-01/real-time-forum/internal/store"
	"context"
)

//RoomRepository ...
type RoomRepository struct {
	store *Store
}

//Room ...
func (s *Store) Room() store.RoomRepository {
	if s.roomRepository != nil {
		return s.roomRepository
	}
	s.roomRepository = &RoomRepository{
		store: s,
	}
	return s.roomRepository
}

//GetRoom ...
func (rr *RoomRepository) GetRoom(userID1, userID2 int) (*model.Room, error) {
	room := &model.Room{}
	row := rr.store.db.QueryRow(
		`SELECT room_id FROM room_participants WHERE user_id in ($1,$2) GROUP BY room_id HAVING COUNT(*)>1;`,
		userID1, userID2)
	err := row.Scan(&room.ID)
	if err != nil {
		return nil, err
	}

	return room, nil
}

//CreateRoom ...
func (rr *RoomRepository) CreateRoom(userID1, userID2 int) (*model.Room, error) {
	room := &model.Room{}
	ctx := context.Background()
	tx, err := rr.store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	res, err := tx.ExecContext(ctx, `INSERT INTO rooms (id) values (null)`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	roomID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	room.ID = int(roomID)

	_, err = tx.ExecContext(ctx, `
	INSERT INTO room_participants (user_id, room_id) VALUES ($1, $2)
	`, userID1, int(roomID))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.ExecContext(ctx, `
	INSERT INTO room_participants (user_id, room_id) VALUES ($1, $2)
	`, userID2, int(roomID))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return room, nil
}

//DeleteRoom ...
func (rr *RoomRepository) DeleteRoom(id int) error {
	ctx := context.Background()
	tx, err := rr.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `DELETE FROM rooms WHERE id=$1`, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `
	DELETE FROM room_participants WHERE room_id=$1
	`, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `
	DELETE FROM messages WHERE room_id=$1
	`, id)
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

//NewMessage ...
func (rr *RoomRepository) NewMessage(roomID int, m *model.Message) error {
	_, err := rr.store.db.Exec(
		`INSERT INTO messages (room_id, author_id, message_text, message_timestamp) VALUES ($1, $2, $3, $4)`,
		roomID,
		m.User.ID,
		m.Text,
		m.Timestamp,
	)
	if err != nil {
		return err
	}

	return nil
}

//GetMessages ...
func (rr *RoomRepository) GetMessages(roomID int, from int) ([]*model.Message, error) {
	rows, err := rr.store.db.Query(`
		SELECT message_timestamp, message_text, users.id, users.nickname FROM messages
		LEFT JOIN users ON messages.author_id = users.id
		WHERE room_id = $1
		LIMIT 10 OFFSET $2
		`, roomID, from)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	messages := make([]*model.Message, 0)
	for rows.Next() {
		user := &model.User{}
		message := &model.Message{}
		err := rows.Scan(
			&message.Timestamp,
			&message.Text,
			&user.ID,
			&user.Nickname,
		)
		if err != nil {
			return nil, err
		}
		message.User = user
		messages = append(messages, message)
	}
	return messages, nil
}
