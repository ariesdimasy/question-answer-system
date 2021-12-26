package helper

import "errors"

var (
	ErrNotFound = errors.New("data not found")
	ErrDbServer = errors.New("db server error")
)
