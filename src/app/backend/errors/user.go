package errors

import "fmt"

// EmptyName ...
type EmptyName struct{}

func (err EmptyName) Error() string {
	return "Empty error"
}

// IsEemptyName return weheter name is empty
func IsEemptyName(err error) bool {
	_, ok := err.(EmptyName)
	return ok
}

// UserNotExist definition
type UserNotExist struct {
	UserID   int64
	UserName string
}

// IsUserNotExist return whether user is not exist
func IsUserNotExist(err error) bool {
	_, ok := err.(UserNotExist)
	return ok
}

// Error ...
func (err UserNotExist) Error() string {
	return "User not exist"
}

// NameReserved definition
type NameReserved struct {
	Name string
}

// IsNameReserved return wether username is reservedUsernames
func IsNameReserved(err error) bool {
	_, ok := err.(NameReserved)
	return ok
}

// Error ...
func (err NameReserved) Error() string {
	return fmt.Sprintf("name %s reserved", err.Name)
}

// ErrUserAlreadyExist ...
type ErrUserAlreadyExist struct {
	Name string
}

// IsErrUserAlreadyExist ...
func IsErrUserAlreadyExist(err error) bool {
	_, ok := err.(ErrUserAlreadyExist)
	return ok
}

// Error ...
func (err ErrUserAlreadyExist) Error() string {
	return fmt.Sprintf("user already exists [name: %s]", err.Name)
}

// ErrEmailAlreadyUsed ...
type ErrEmailAlreadyUsed struct {
	Email string
}

// IsErrEmailAlreadyUsed ...
func IsErrEmailAlreadyUsed(err error) bool {
	_, ok := err.(ErrEmailAlreadyUsed)
	return ok
}

// Error ...
func (err ErrEmailAlreadyUsed) Error() string {
	return fmt.Sprintf("e-mail has been used [email: %s]", err.Email)
}
