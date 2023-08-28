package server

import "errors"

var (
	ErrBadRequest     = errors.New("invalid_request")
	ErrNotImplemented = errors.New("not_implemented")
	ErrConflict       = errors.New("conflict")
)
