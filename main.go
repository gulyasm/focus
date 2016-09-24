package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Commands: add, list")
}

func exit(code int) {
	os.Exit(code)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		exit(1)
	}
	fs := NewFocusStore("test.db")
	cmd := os.Args[1]
	var err error
	switch cmd {
	case "add":
		err = fs.Add()
	case "list":
		err = fs.List()
	default:
		usage()
		exit(2)
	}
	if err != nil {
		usage()
		exit(3)
	}
}
