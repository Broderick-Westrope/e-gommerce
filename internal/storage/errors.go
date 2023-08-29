package storage

import "fmt"

type NotFoundError struct {
	Operation string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found: %s", e.Operation)
}
