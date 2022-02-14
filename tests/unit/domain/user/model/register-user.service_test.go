package model_test

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasswordLessThanRequired(t *testing.T) {
	assert := require.New(t)
	_, err := model.NewUser("JUan", "1234", "USER", "test@gmail.com")
	assert.Error(err, "The leng of the password is incorrect")
	assert.True(err.Error() == "The leng of the password is incorrect")
}

func TestUserInstanceCorrect(t *testing.T) {
	assert := require.New(t)
	var expectUser *model.User
	expectUser = &model.User{Name: "JUan", Password: "123ss3r4", Creation_date: time.Now(), Role: "USER", Email: "test@gmail.com"}
	usr, err := model.NewUser(expectUser.Name, expectUser.Password, expectUser.Role, expectUser.Email)
	assert.True(err == nil)
	assert.True(usr != nil)
	assert.Equal(expectUser, usr)
}
