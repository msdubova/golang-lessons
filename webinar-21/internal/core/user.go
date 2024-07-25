package core

import "errors"

type User struct {
	ID       int
	Nickname string
	Password string
}

var (
	ErrUserNotFound = errors.New("user not found")
)
