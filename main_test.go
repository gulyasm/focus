package main

import (
	"os/user"
	"testing"
)

func TestDBPath(t *testing.T) {
	u, err := user.Current()
	homeDir := u.HomeDir
	if err != nil {
		t.Error(err)
	}
	s, err := getDBPath()
	if err != nil {
		t.Error(err)
	}
	e := homeDir + "/.focus.db"
	if s != e {
		t.Errorf("getDBPath should be %v WAS: %v", e, s)
	}
}
