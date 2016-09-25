package main

import (
	"testing"
	"time"
)

func TestRunning(t *testing.T) {
	e := Element{
		ID:    1,
		Name:  "TestElement",
		Start: time.Now(),
	}

	if !e.IsRunning() {
		t.Error("Element should be running but IsRunning returned false")
	}
}

func TestNotRunning(t *testing.T) {
	e := Element{
		ID:    1,
		Name:  "TestElement",
		Start: time.Now(),
		End:   time.Now().Add(time.Minute * 5),
	}

	if e.IsRunning() {
		t.Error("Element has End time, but IsRunning returned true")
	}
}

func TestDuration(t *testing.T) {
	now := time.Now()
	e := Element{
		ID:    1,
		Name:  "TestElement",
		Start: now,
		End:   now.Add(time.Minute * 5),
	}

	if e.Duration() != time.Minute*5 {
		t.Error("Element should have 5 minute duration. Actual: ", e.Duration())
	}
}

func TestRunningDuration(t *testing.T) {
	now := time.Now()
	e := Element{
		ID:    1,
		Name:  "TestElement",
		Start: now,
	}
	d1 := e.Duration()
	time.Sleep(time.Millisecond * 50)
	d2 := e.Duration()
	if d2 <= d1 {
		t.Errorf("For running task Duration should be increase. D1: %v, D2: %v\n", d1, d2)
	}
}
