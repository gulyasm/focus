package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/apcera/termtables"
)

func usage(err string) {
	fmt.Println("Commands: add, list, today")
	fmt.Println("Error: " + err)
}

func exit(code int) {
	os.Exit(code)
}

func cmdToday(fs FocusStore) error {
	results, err := fs.ListToday()
	if err != nil {
		return err
	}
	printElements(results)
	return nil
}

func cmdNow(fs FocusStore) error {
	result, err := fs.Now()
	if err != nil {
		return err
	}
	fmt.Println(result.Name)
	return nil
}

func cmdList(fs FocusStore) error {
	results, err := fs.List()
	if err != nil {
		return err
	}
	printElements(results)
	return nil
}

func printElements(elements []Element) {
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Duration")
	var td time.Duration

	for _, r := range elements {
		d := r.Duration()
		table.AddRow(r.Name, d)
		td += d
	}
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Total duration:", td)

	fmt.Println(table.Render())
}
func main() {
	if len(os.Args) < 2 {
		usage("Missing cmd")
		exit(1)
	}
	fs, err := NewFocusStore("test.db")
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[1]
	switch cmd {
	case "add":
		name := os.Args[2]
		err = fs.Add(name)
	case "list":
		err = cmdList(fs)
	case "today":
		err = cmdToday(fs)
	case "now":
		err = cmdNow(fs)
	default:
		usage("Command not found: " + cmd)
		exit(2)
	}
	if err != nil {
		usage(err.Error())
		exit(3)
	}
}
