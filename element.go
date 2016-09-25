package main

import "time"

type Element struct {
	Id    int
	Name  string
	Start time.Time
	End   time.Time
}

func (e Element) Duration() time.Duration {
	if e.End.IsZero() {
		return time.Since(e.Start)
	}
	return e.End.Sub(e.Start)
}

func (e Element) IsRunning() bool {
	return e.End.IsZero()
}
