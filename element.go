package main

import "time"

// Element represents an task.
type Element struct {
	ID    int
	Name  string
	Start time.Time
	End   time.Time
}

// Duration returns the the time duration for the
// given Element. If it's still running, the time elapsed
// since the start.
func (e Element) Duration() time.Duration {
	if e.End.IsZero() {
		return time.Since(e.Start)
	}
	return e.End.Sub(e.Start)
}

// IsRunning returns true if the element is still running, false
// otherwise.
func (e Element) IsRunning() bool {
	return e.End.IsZero()
}
