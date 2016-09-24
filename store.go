package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const ISO8601 = "2006-01-02 15:04:05.000 MST"

type Element struct {
	Id    int
	Name  string
	Start time.Time
	End   time.Time
}

func (e Element) Duration() time.Duration {
	if e.End.IsZero() {
		return time.Since(e.Start)
	}
	return e.End.Sub(e.Start)
}

type FocusStore struct {
	path string
}

func NewFocusStore(path string) (FocusStore, error) {
	fs := FocusStore{path}
	err := fs.initDB()
	return fs, err
}

func (fs FocusStore) initDB() error {
	db, err := fs.open()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStmt := `
			CREATE TABLE IF NOT EXISTS Elements (
				Id INTEGER NOT NULL PRIMARY KEY, 
				Name TEXT,
				Start TEXT DEFAULT "",
				End TEXT DEFAULT ""
			);
	`
	_, err = db.Exec(sqlStmt)
	return err
}

func (fs FocusStore) ListToday() ([]Element, error) {
	result := []Element{}
	db, err := fs.open()
	if err != nil {
		return result, err
	}
	defer db.Close()
	today := time.Now().Format("2006-01-02")
	rows, err := db.Query("SELECT Id, Name, Start, End FROM Elements WHERE Start > ?", today)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		e := Element{}
		var start, end string
		if err := rows.Scan(
			&e.Id,
			&e.Name,
			&start,
			&end,
		); err != nil {
			return result, err
		}
		if start != "" {
			e.Start, err = time.Parse(ISO8601, start)
			if err != nil {
				log.Println("Failed to parse time: ", err.Error())
			}
		}
		if end != "" {
			e.End, err = time.Parse(ISO8601, end)
			if err != nil {
				log.Println("Failed to parse time: ", err.Error())
			}
		}
		result = append(result, e)

	}

	return result, nil
}

func (fs FocusStore) Now() (Element, error) {
	e := Element{}
	db, err := fs.open()
	if err != nil {
		return e, err
	}
	defer db.Close()
	var start string
	row := db.QueryRow("SELECT Id, Name, Start FROM Elements WHERE End=''")
	if err = row.Scan(&e.Id, &e.Name, &start); err != nil {
		return e, err
	}
	if start != "" {
		e.Start, err = time.Parse(ISO8601, start)
		if err != nil {
			log.Println("Failed to parse time: ", err.Error())
		}
	}
	return e, nil
}

func (fs FocusStore) List() ([]Element, error) {
	result := []Element{}
	db, err := fs.open()
	if err != nil {
		return result, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT Id, Name, Start, End FROM Elements")
	if err != nil {
		return result, err
	}

	for rows.Next() {
		e := Element{}
		var start, end string
		if err := rows.Scan(
			&e.Id,
			&e.Name,
			&start,
			&end,
		); err != nil {
			return result, err
		}
		if start != "" {
			e.Start, err = time.Parse(ISO8601, start)
			if err != nil {
				log.Println("Failed to parse time: ", err.Error())
			}
		}
		if end != "" {
			e.End, err = time.Parse(ISO8601, end)
			if err != nil {
				log.Println("Failed to parse time: ", err.Error())
			}
		}
		result = append(result, e)

	}

	return result, nil
}

func (fs FocusStore) Add(name string) error {
	db, err := fs.open()
	if err != nil {
		return err
	}
	defer db.Close()

	now := time.Now()

	// Close previous activity
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE Elements SET End=? WHERE END=''", now.Format(ISO8601))
	if err != nil {
		log.Println("Failed to update DB.", err.Error())
		err = tx.Rollback()
		return err
	}
	// Start new activity
	_, err = tx.Exec(
		"INSERT INTO Elements(Name, Start) VALUES(?, ?)",
		name,
		now.Format(ISO8601))
	if err != nil {
		log.Println("Failed to update DB.", err.Error())
		err = tx.Rollback()
		return err
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		log.Println("Failed to commit transaction", err.Error())
		err = tx.Rollback()
		return err
	}
	return nil

}

func (fs FocusStore) open() (*sql.DB, error) {
	return sql.Open("sqlite3", fs.path)
}
