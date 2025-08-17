package error

import "errors"

const (
	ErrCreatingCustomer = "an error occurred while trying to register the customer"
)

func CreatingCustomer() error {
	return errors.New(ErrCreatingCustomer)
}
