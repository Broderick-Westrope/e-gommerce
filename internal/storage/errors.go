package storage

import "fmt"

// NotFoundError is an error that is returned when a resource is not found.
type NotFoundError struct {
	Operation string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found: %s", e.Operation)
}
