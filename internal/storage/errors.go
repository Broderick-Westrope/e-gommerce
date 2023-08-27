package storage

type NotFoundError struct {
	Operation string
}

func (e *NotFoundError) Error() string {
	return "Not found: " + e.Operation
}
