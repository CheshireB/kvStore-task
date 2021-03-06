package sbercloid_test_task

import "fmt"

type (
	keyNotExistError struct{ key interface{} }

	keyAlreadyExistError struct{ key interface{} }

	keyIsNotComparableError struct{ key interface{} }
)

const (
	keyNotExistErrorTemplate        = "key %q not exist"
	keyAlreadyExistErrorTemplate    = "key %q already exist"
	keyIsNotComparableErrorTemplate = "key %q is not comparable"
)

func NewKeyNotExistError(key interface{}) error {
	return &keyNotExistError{key: key}
}

func NewKeyAlreadyExistError(key interface{}) error {
	return &keyAlreadyExistError{key: key}
}

func NewKeyIsNotComparableError(key interface{}) error {
	return &keyIsNotComparableError{key: key}
}

func (e *keyNotExistError) Error() string {
	return fmt.Sprintf(keyNotExistErrorTemplate, fmt.Sprintf("%v", e.key))
}

func (e *keyAlreadyExistError) Error() string {
	return fmt.Sprintf(keyAlreadyExistErrorTemplate, fmt.Sprintf("%v", e.key))
}

func (e *keyIsNotComparableError) Error() string {
	return fmt.Sprintf(keyIsNotComparableErrorTemplate, fmt.Sprintf("%v", e.key))
}

func IsKeyNotExistError(err error) bool {
	_, ok := err.(*keyNotExistError)
	return ok
}

func IsKeyAlreadyExistError(err error) bool {
	_, ok := err.(*keyAlreadyExistError)
	return ok
}

func IsKeyIsNotComparableError(err error) bool {
	_, ok := err.(*keyIsNotComparableError)
	return ok
}
