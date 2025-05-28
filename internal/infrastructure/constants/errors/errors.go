package errors

import "errors"

type MessageError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	ErrNotFound = errors.New("document is nil")
)
