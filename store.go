package main

import (
	"errors"
	"time"
)

const ISO8601 = "2006-01-02 15:04:05.000Z"

var ErrNoElement = errors.New("No element found")

type Store interface {
	Now() (Element, error)
	Start(name string) error
	Stop() error
	List() ([]Element, error)
	ListDay(day time.Time) ([]Element, error)
}
