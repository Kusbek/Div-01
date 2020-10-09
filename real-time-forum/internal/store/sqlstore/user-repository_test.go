package sqlstore

import (
	"DIV-01/real-time-forum/internal/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sqlOpt = &Options{
		Address: "../../../local.db",
	}
)

func Test_UserRepository(t *testing.T) {
	st, err := Start(sqlOpt)
	if err != nil {
		log.Fatal(err)
	}
	pass := "lopata"
	user := model.TestUser("kusbek_test", pass)

	exists, err := st.User().Exists(user.Nickname, user.Email)
	assert.NoError(t, err)
	assert.Equal(t, false, exists)

	err = st.User().Create(user)
	assert.NoError(t, err)
	assert.NotEqual(t, user.Password, pass)
	assert.NotNil(t, user.ID)

	nickUser, err := st.User().Find(user.Nickname)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, nickUser.Email)
	assert.NoError(t, nickUser.ComparePasswords(pass))

	eUser, err := st.User().Find(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.Nickname, eUser.Nickname)

	exists, err = st.User().Exists(user.Nickname, user.Email)
	assert.NoError(t, err)
	assert.Equal(t, true, exists)

	assert.NoError(t, st.User().Delete(user.ID))
}
