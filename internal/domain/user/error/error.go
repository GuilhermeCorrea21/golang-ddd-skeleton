package error

import "errors"

const (
	ErrUserNotFound = "user not found"
	ErrCreatingUser = "error when trying to create user"
)

func UserNotFound() error {
	return errors.New(ErrUserNotFound)
}

func CreatingUser() error {
	return errors.New(ErrCreatingUser)
}
