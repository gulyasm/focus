package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStore stores Elements in an SQLite database persisted to
// the disk.
type SQLiteStore struct {
	path string
}

// NewSQLiteStore returns an SQLiteStore with a backed SQLite DB on the
// provided path. The DB is initialized if it didn't exist before.
// Already existing DB is not touched.
func NewSQLiteStore(path string) (Store, error) {
	if path == "" {
		return SQLiteStore{}, errors.New("Empty path provided to SQLiteStore")
	}
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

// ListDay returns the elements for a given day. The day parameter shouldn't have to be
// midnight or any other specific time. Any time on the given date is fine.
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
			&e.ID,
			&e.Name,
			&start,
			&end,
		); err != nil {
			return result, err
		}
		if start != "" {
			e.Start, err = time.Parse(iso8601, start)
			if err != nil {
				log.Println("Failed to parse time: ", err.Error())
			}
		}
		if end != "" {
			e.End, err = time.Parse(iso8601, end)
			if err != nil {
				log.Println("Failed to parse time: ", err.Error())
			}
		}
		result = append(result, e)

	}

	return result, nil
}

// Now returns the running task. If there is no running task the ErrNoElement error is returned.
func (fs SQLiteStore) Now() (Element, error) {
	e := Element{}
	db, err := fs.open()
	if err != nil {
		return e, err
	}
	defer db.Close()
	var start string
	row := db.QueryRow("SELECT Id, Name, Start FROM Elements WHERE End=''")
	if err = row.Scan(&e.ID, &e.Name, &start); err != nil {
		if err == sql.ErrNoRows {
			return e, ErrNoElement
		}
		return e, err
	}
	if start != "" {
		e.Start, err = time.Parse(iso8601, start)
		if err != nil {
			log.Println("Failed to parse time: ", err.Error())
		}
	}
	return e, nil
}

// List returns all elements in the store.
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
			&e.ID,
			&e.Name,
			&start,
			&end,
		); err != nil {
			return result, err
		}
		if start != "" {
			e.Start, err = time.Parse(iso8601, start)
			if err != nil {
				log.Println("Failed to parse time: ", err.Error())
			}
		}
		if end != "" {
			e.End, err = time.Parse(iso8601, end)
			if err != nil {
				log.Println("Failed to parse time: ", err.Error())
			}
		}
		result = append(result, e)

	}

	return result, nil
}

// Stop stops the currently running task
func (fs SQLiteStore) Stop() error {
	db, err := fs.open()
	if err != nil {
		return err
	}
	defer db.Close()

	now := time.Now().UTC()

	_, err = db.Exec("UPDATE Elements SET End=? WHERE END=''", now.Format(iso8601))
	if err != nil {
		log.Println("Failed to update DB.", err.Error())
	}
	return err
}

// Start adds a new Element and start it.
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
	_, err = tx.Exec("UPDATE Elements SET End=? WHERE END=''", now.Format(iso8601))
	if err != nil {
		log.Println("Failed to update DB.", err.Error())
		err = tx.Rollback()
		return err
	}
	// Start new activity
	_, err = tx.Exec(
		"INSERT INTO Elements(Name, Start) VALUES(?, ?)",
		name,
		now.Format(iso8601))
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
