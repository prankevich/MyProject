package errs

import "errors"

var (
	ErrNotfound              = errors.New("not found")
	ErrEmployeesNotfound     = errors.New("user not found")
	ErrInvalidEmployeesID    = errors.New("invalid user id")
	ErrInvalidRequestBody    = errors.New("invalid request body")
	ErrInvalidFieldValue     = errors.New("invalid field value")
	ErrUserNameAlreadyExists = errors.New("user name already exists")
)
