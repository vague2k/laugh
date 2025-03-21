package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vague2k/laugh/utils"
)

type DB struct {
	sql      *sql.DB
	version  string
	dir      string
	filename string
}

// Initialize a new laugh db, returning a pointer to the instance.
//
// [New] also makes sure the "events" table exists.
//
// NOTE: private fields are *usually* meant for internal use, usually tests.
func New(path string) (*DB, error) {
	if path == "" {
		dir, err := utils.UserDataHome()
		if err != nil {
			return nil, err
		}
		path = dir
	}

	var dir string
	var file string
	if path == ":memory:" {
		file = ":memory:"
	} else {
		dir = filepath.Join(path, "laugh")
		file = filepath.Join(dir, "laugh.db")
		err := os.MkdirAll(dir, 0o777)
		if err != nil {
			return nil, fmt.Errorf("could not create db dir: \n%s", err)
		}
	}

	database, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("could not init rummage db: \n%s", err)
	}

	_, err = database.Exec(`
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		summary TEXT NOT NULL,
		description TEXT NOT NULL,
		course TEXT NOT NULL,
		due_date TEXT NOT NULL,
		due_hour TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0
	)

    `)
	if err != nil {
		return nil, fmt.Errorf("could not create 'items' table in rummage db: \n%s", err)
	}

	instance := DB{
		sql:      database,
		dir:      dir,
		filename: file,
		version:  "v0.1.0",
	}

	return &instance, nil
}

func (db *DB) Dir() string {
	return db.dir
}

func (db *DB) Filename() string {
	return db.filename
}

func (db *DB) Version() string {
	return db.version
}

func (db *DB) Close() error {
	return db.sql.Close()
}
