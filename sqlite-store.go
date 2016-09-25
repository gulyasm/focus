package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	path string
}

func NewSQLiteStore(path string) (Store, error) {
	fs := SQLiteStore{path}
	err := fs.initDB()
	return fs, err
}

func (fs SQLiteStore) initDB() error {
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

func (fs SQLiteStore) ListDay(day time.Time) ([]Element, error) {
	result := []Element{}
	db, err := fs.open()
	if err != nil {
		return result, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT Id, Name, Start, End FROM Elements WHERE date(Start) = date(?)", day)
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

func (fs SQLiteStore) Now() (Element, error) {
	e := Element{}
	db, err := fs.open()
	if err != nil {
		return e, err
	}
	defer db.Close()
	var start string
	row := db.QueryRow("SELECT Id, Name, Start FROM Elements WHERE End=''")
	if err = row.Scan(&e.Id, &e.Name, &start); err != nil {
		if err == sql.ErrNoRows {
			return e, ErrNoElement
		}
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

func (fs SQLiteStore) List() ([]Element, error) {
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

func (fs SQLiteStore) Stop() error {
	db, err := fs.open()
	if err != nil {
		return err
	}
	defer db.Close()

	now := time.Now().UTC()

	_, err = db.Exec("UPDATE Elements SET End=? WHERE END=''", now.Format(ISO8601))
	if err != nil {
		log.Println("Failed to update DB.", err.Error())
	}
	return err
}

func (fs SQLiteStore) Start(name string) error {
	db, err := fs.open()
	if err != nil {
		return err
	}
	defer db.Close()

	now := time.Now().UTC()

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

func (fs SQLiteStore) open() (*sql.DB, error) {
	return sql.Open("sqlite3", fs.path)
}