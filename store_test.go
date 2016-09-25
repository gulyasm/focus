package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	s, err := NewFocusStore(tmpfile.Name())
	if err != nil {
		t.Error(err)
	}
	err = s.Add("Hello")
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
		t.Error("Added element doesn't match with expected. Expected: %s, actual: %s\n",
			"Hello", r[0].Name)
	}

}