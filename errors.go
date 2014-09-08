package config

import (
	"errors"
)

var (
	ErrConfigFileNotFound = errors.New(`Configuration file was not found.`)
)
