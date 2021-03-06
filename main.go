package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/apcera/termtables"
)

const defaultDBPath = ".focus.db"

func getDBPath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	s := u.HomeDir + "/" + defaultDBPath
	return s, err
}

func usage(err string) {
	fmt.Println("Error: " + err)
	helpMessage := `
Usage:
focus COMMAND OPTIONS

COMMANDS
	now
		Prints the currently running task

		EXAMPLE:
		focus now

	start
		Start a new task and stops any running. The start time will be automatically the
		time when the task is added.

		OPTIONS:
		The name of the task.

		EXAMPLE:
		focus start "Coding Focus interface"

	stop
		Stops the currently runnnig task.

		EXAMPLE:
		focus stop

	list
		List ALL tasks.

		EXAMPLE:
		focus list

	today
		List tasks started today with duration.

		EXAMPLE:
		focus today

	yesterday
		List tasks started yesterday with duration.

		EXAMPLE:
		focus yesterday
`
	fmt.Println(helpMessage)
}

func exit(code int) {
	os.Exit(code)
}

func cmdYesterday(s Store) error {
	yesterday := time.Now().AddDate(0, 0, -1)
	results, err := s.ListDay(yesterday)
	if err != nil {
		return err
	}
	printElements(results)
	return nil
}

func cmdToday(s Store) error {
	today := time.Now()
	results, err := s.ListDay(today)
	if err != nil {
		return err
	}
	printElements(results)
	return nil
}

func cmdNow(s Store) error {
	result, err := s.Now()
	if err != nil && err != ErrNoElement {
		return err
	}
	if err != ErrNoElement {
		fmt.Println(result.Name)
	}
	return nil
}

func cmdList(s Store) error {
	results, err := s.List()
	if err != nil {
		return err
	}
	printElements(results)
	return nil
}

func cmdStop(s Store) error {
	return s.Stop()
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
		usage("Missing command")
		exit(1)
	}
	path, err := getDBPath()
	if err != nil {
		log.Fatal(err)
	}
	fs, err := NewSQLiteStore(path)
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
	case "yesterday":
		err = cmdYesterday(fs)
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
