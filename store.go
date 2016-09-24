package main

import "fmt"

type FocusStore struct {
	path string
}

func NewFocusStore(path string) *FocusStore {
	return &FocusStore{path: path}
}

func (fs *FocusStore) List() error {
	fmt.Println("List")
	return nil
}

func (fs *FocusStore) Add() error {
	fmt.Println("Add")
	return nil
}
