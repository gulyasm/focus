package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type FocusStore struct {
	db sql.DB
}

func NewFocusStore(path string) (*FocusStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStmt := `
			CREATE TABLE Elements (
				Id INTEGER NOT NULL PRIMARY KEY, 
				Name TEXT,
				Start TEXT,
				End TEXT
			);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return &FocusStore{db}, err
}

func (fs *FocusStore) List() error {
	fmt.Println("List")
	return nil
}

func (fs *FocusStore) Add() error {
	fmt.Println("Add")
	return nil
}
