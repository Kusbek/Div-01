package sqlstore

import (
	"DIV-01/real-time-forum/internal/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PostRepository(t *testing.T) {
	st, err := Start(sqlOpt)
	if err != nil {
		log.Fatal(err)
	}

	posts, err := st.Post().GetAll()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(posts))

	post := model.TestPost(posts[0].Author.ID)
	err = st.Post().Create(post)
	assert.NoError(t, err)

	err = st.Post().Delete(post.ID)
	assert.NoError(t, err)

	afterCreatePosts, err := st.Post().GetAll()
	assert.NoError(t, err)
	assert.Equal(t, len(afterCreatePosts), len(posts))
}
