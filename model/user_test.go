package model

import "testing"

func TestUserMigrate(t *testing.T) {
	UserMigrate()
}

func TestCreateUser(t *testing.T) {
	CreateUser("nagaoka", "aaa@gmail.com", "pass")
}

func TestLoginCheck(t *testing.T) {
	LoginCheck("aaa@gmail.com", "sass")
}
