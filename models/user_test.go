package models

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestUserPasswordIsHashed(t *testing.T) {

	user := &User{
		Password: "12345678",
	}

	user.HashPassword()

	assert.NotEqual(t, "12345678", user.Password)

}
