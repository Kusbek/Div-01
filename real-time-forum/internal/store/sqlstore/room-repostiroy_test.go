package sqlstore

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RoomRepository(t *testing.T) {
	st, err := Start(sqlOpt)
	if err != nil {
		log.Fatal(err)
	}
	user1, user2 := 1, 2

	// createdRoom, err := st.Room().CreateRoom(user1, user2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// assert.NoError(t, err)

	room, err := st.Room().GetRoom(user1, user2)
	assert.NoError(t, err)

	// assert.Equal(t, createdRoom.ID, room.ID)
	// for i := 0; i < 100; i++ {
	// 	err = st.Room().NewMessage(room.ID, &model.Message{
	// 		Timestamp: time.Now(),
	// 		Text:      "Test Message" + strconv.Itoa(i),
	// 		User: &model.User{
	// 			ID: func() int {
	// 				if i%2 == 0 {
	// 					return user1
	// 				}
	// 				return user2
	// 			}(),
	// 		},
	// 	})
	// }

	messages, err := st.Room().GetMessages(room.ID, 0)
	fmt.Println(messages[0])
	assert.NoError(t, err)
	err = st.Room().DeleteRoom(room.ID)
	assert.NoError(t, err)

	// _, err = st.Room().GetRoomID(1, 10000)
	// assert.Equal(t, sql.ErrNoRows, err)
}

func Test_GetLastMessage(t *testing.T) {
	st, err := Start(sqlOpt)
	if err != nil {
		log.Fatal(err)
	}
	user1, user2 := 1, 2

	room, err := st.Room().GetRoom(user1, user2)
	assert.NoError(t, err)

	timestamp, err := st.Room().GetLastMessageTimestamp(room.ID)
	fmt.Println(timestamp)
	assert.NoError(t, err)
}
