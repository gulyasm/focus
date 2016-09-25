package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/apcera/termtables"
)

const DEFAULT_DB_NAME = "focus.db"

func getDBPath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	s := u.HomeDir + "/" + DEFAULT_DB_NAME
	return s, err
}

func usage(err string) {
	fmt.Println("Commands: start, stop, list, today, yesterday")
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
	if err != nil && err != ErrNoElement {
		return err
	}
	if err != ErrNoElement {
		fmt.Println(result.Name)
	}
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

func cmdStop(fs FocusStore) error {
	return fs.Stop()
}

func printElements(elements []Element) {
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Duration")
	var td time.Duration

	for _, r := range elements {
		d := r.Duration()
		var tag string
		if r.IsRunning() {
			tag = " (R)"
		}
		table.AddRow(r.Name+tag, d)
		td += d
	}
	fmt.Println(table.Render())
	fmt.Println("Total duration:", td)

}
func main() {
	if len(os.Args) < 2 {
		usage("Missing cmd")
		exit(1)
	}
	path, err := getDBPath()
	if err != nil {
		log.Fatal(err)
	}
	fs, err := NewFocusStore(path)
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[1]
	switch cmd {
	case "start":
		name := os.Args[2]
		err = fs.Start(name)
	case "list":
		err = cmdList(fs)
	case "today":
		err = cmdToday(fs)
	case "now":
		err = cmdNow(fs)
	case "stop":
		err = cmdStop(fs)
	default:
		usage("Command not found: " + cmd)
		exit(2)
	}
	if err != nil {
		usage(err.Error())
		exit(3)
	}
}
