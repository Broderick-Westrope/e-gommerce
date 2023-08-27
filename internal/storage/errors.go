package storage

type ErrNotFound struct {
	Operation string
}

func (e *ErrNotFound) Error() string {
	return "Not found: " + e.Operation
}
