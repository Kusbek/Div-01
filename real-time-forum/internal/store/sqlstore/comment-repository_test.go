package sqlstore

import (
	"DIV-01/real-time-forum/internal/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CommentRepository(t *testing.T) {
	st, err := Start(sqlOpt)
	if err != nil {
		log.Fatal(err)
	}

	posts, err := st.Post().GetAll()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(posts))

	p := posts[0]

	newComment := &model.Comment{
		Author: p.Author,
		PostID: p.ID,
		Text:   "4e za hu",
	}

	err = st.Comment().Create(newComment)
	assert.NoError(t, err)

	aPosts, err := st.Post().GetAll()

	ap := aPosts[0]

	assert.NoError(t, err)
	assert.Equal(t, p.Comments+1, ap.Comments)

	comments, err := st.Comment().Get(p.ID)
	assert.NoError(t, err)
	assert.Equal(t, ap.Comments, len(comments))

}
