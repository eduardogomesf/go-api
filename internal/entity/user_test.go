package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUse(t *testing.T) {
	user, err := NewUser("John Doe", "j.d@mail.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j.d@mail.com", user.Email)
}

func TestPasswordValidation(t *testing.T) {
	user, err := NewUser("John Doe", "j.d@mail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("123457"))
	assert.NotEqual(t, "123457", user.Password)
}
