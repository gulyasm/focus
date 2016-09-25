package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestStop(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	s, err := NewSQLiteStore(tmpfile.Name())
	if err != nil {
		t.Error(err)
	}
	err = s.Start("Hello")
	if err != nil {
		t.Error(err)
	}
	r, err := s.List()
	if err != nil {
		t.Error(err)
	}
	if len(r) != 1 {
		t.Error(err)
	}
	if !r[0].IsRunning() {
		t.Error("IsRunning returned false for an expected running element")
	}
	s.Stop()
	r, err = s.List()
	if err != nil {
		t.Error(err)
	}
	if len(r) != 1 {
		t.Error(err)
	}
	if r[0].IsRunning() {
		t.Error("IsRunning returned false for an expected running element")
	}
}

func TestStart(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	s, err := NewSQLiteStore(tmpfile.Name())
	if err != nil {
		t.Error(err)
	}
	err = s.Start("Hello")
	if err != nil {
		t.Error(err)
	}
	r, err := s.List()
	if err != nil {
		t.Error(err)
	}
	if len(r) != 1 {
		t.Error(err)
	}
	if r[0].Name != "Hello" {
		t.Errorf("Added element doesn't match with expected. Expected: %s, actual: %s\n",
			"Hello", r[0].Name)
	}
	if !r[0].IsRunning() {
		t.Error("IsRunning returned false for an expected running element")
	}
}

func TestToday(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	s, err := NewSQLiteStore(tmpfile.Name())
	if err != nil {
		t.Error(err)
	}
	err = s.Start("Hello")
	err = s.Start("Hello2")
	if err != nil {
		t.Error(err)
	}
	r, err := s.ListDay(time.Now())
	if err != nil {
		t.Error(err)
	}
	if len(r) != 2 {
		t.Error(err)
	}
	if !r[1].IsRunning() {
		t.Error("IsRunning returned false for an expected running element")
	}
	if r[0].IsRunning() {
		t.Error("IsRunning returned true for an expected finished element")
	}
}

func TestNow(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	s, err := NewSQLiteStore(tmpfile.Name())
	if err != nil {
		t.Error(err)
	}
	err = s.Start("Hello")
	err = s.Start("Hello2")
	if err != nil {
		t.Error(err)
	}
	r, err := s.ListDay(time.Now())
	if err != nil {
		t.Error(err)
	}
	if len(r) != 2 {
		t.Error(err)
	}
	if !r[1].IsRunning() {
		t.Error("IsRunning returned false for an expected running element")
	}
	if r[0].IsRunning() {
		t.Error("IsRunning returned true for an expected finished element")
	}

	// Test Now with running task
	e, err := s.Now()
	if err != nil {
		t.Error("Now() returned err but it should be nil", err)

	}
	if !e.IsRunning() {
		t.Error("Now() returned a task that is not Running", t)
	}
	if e.Name != "Hello2" {
		t.Error("Hello2 should be the running task as that was the last started", e)
	}

	// Test Now without running task
	s.Stop()
	_, err = s.Now()
	if err == nil {
		t.Error("Now() returned nil err but it should be ErrNoElement", err)
	}
	if err != ErrNoElement {
		t.Error("Now() returned err but which should have been ErrNoElement but wasn't", err)
	}
}
