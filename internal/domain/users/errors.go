package users

import "errors"

var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrUserNotFound     = errors.New("user not found")
	ErrUsernameTaken    = errors.New("username already taken")
	ErrCannotFollowSelf = errors.New("cannot follow self")
	ErrAlreadyFollowing = errors.New("already following user")
	ErrNotFollowing     = errors.New("not following user")
)
