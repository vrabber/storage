package driver

import "errors"

var (
	ErrorFileExists    = errors.New("file exists")
	ErrorInvalidOffset = errors.New("invalid offset")
)
