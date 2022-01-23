package sbercloid_test_task

import "fmt"

type (
	keyNotExistError struct{ key interface{} }

	keyAlreadyExistError struct{ key interface{} }
)

const (
	keyNotExistErrorTemplate     = "key %q not exist"
	keyAlreadyExistErrorTemplate = "key %q already exist"
)

func NewKeyNotExistError(key interface{}) error {
	return &keyNotExistError{key: key}
}

func NewAlreadyExistError(key interface{}) error {
	return &keyAlreadyExistError{key: key}
}

func (e *keyNotExistError) Error() string {
	return fmt.Sprintf(keyNotExistErrorTemplate, e.key)
}

func (e *keyAlreadyExistError) Error() string {
	return fmt.Sprintf(keyAlreadyExistErrorTemplate, e.key)
}

func IsKeyNotExistError(err error) bool {
	_, ok := err.(*keyNotExistError)
	return ok
}

func IsKeyAlreadyExistError(err error) bool {
	_, ok := err.(*keyAlreadyExistError)
	return ok
}
