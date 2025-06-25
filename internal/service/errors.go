package service

import "errors"

// ErrNotFound is returned when a requested entity is not found by the service.
var ErrNotFound = errors.New("service: entity not found")
