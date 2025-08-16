package error

import "errors"

const (
	ErrUserNotFound = "user not found"
)

func UserNotFound() error {
	return errors.New(ErrUserNotFound)
}
