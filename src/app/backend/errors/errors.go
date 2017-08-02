package errors

import (
	"errors"
)

// New return error
func New(str string) error {
	return errors.New(str)
}
