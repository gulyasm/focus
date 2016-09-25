package main

import (
	"errors"
	"time"
)

const iso8601 = "2006-01-02 15:04:05.000Z"

// ErrNoElement signals that for the given request from the
// store there was no element to return
var ErrNoElement = errors.New("No element found")

// Store interface represent the Element store where elements are stored.
// Different implementations can persist the elements however it's not a
// requirement.
type Store interface {
	Now() (Element, error)
	Start(name string) error
	Stop() error
	List() ([]Element, error)
	ListDay(day time.Time) ([]Element, error)
}
