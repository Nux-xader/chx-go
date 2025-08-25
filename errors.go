package chx

import "errors"

// ErrNotFound is returned when a key is not found.
var ErrNotFound = errors.New("key not found")

// ErrServer represents an error that occurred on the server.
// It wraps the original error for more detailed inspection.
type ErrServer struct {
	Err error
}

// Error returns the error message, prefixed with "server error:".
func (e *ErrServer) Error() string {
	return "server error: " + e.Err.Error()
}

// Unwrap returns the underlying error, allowing for error chain inspection.
func (e *ErrServer) Unwrap() error {
	return e.Err
}
