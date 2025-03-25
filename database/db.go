package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vague2k/laugh/parser"
	"github.com/vague2k/laugh/utils"
)

type DBInterface interface {
	AddEvent(e *parser.CalendarEvent) error
	SelectItem(id int) (*parser.CalendarEvent, error)
	Events() (*[]parser.CalendarEvent, error)
}

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
		id INTEGER PRIMARY KEY,
		summary TEXT NOT NULL,
		description TEXT NOT NULL,
		course TEXT NOT NULL,
		dueDate TEXT NOT NULL,
		dueHour TEXT NOT NULL,
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

// Adds a [parser.CalendarEvent] into the database with a unique ID. Events
// cannot have duplicate IDs.
//
// It is the caller's responsibility to make sure the event that's passed in has
// all fields populated correctly.
func (db *DB) AddEvent(e *parser.CalendarEvent) error {
	_, err := db.sql.Exec(`
	INSERT INTO events (id, summary, description, course, dueDate, dueHour, completed)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		e.Id, e.Summary, e.Description, e.Course, e.DueDate, e.DueHour, e.Done,
	)
	if err != nil {
		return err
	}

	return nil
}

// Selects a specific event by it's ID, and returns a pointer to it.
// If the item does not exist, it is treated as an error
func (db *DB) SelectItem(id int) (*parser.CalendarEvent, error) {
	var event parser.CalendarEvent

	row := db.sql.QueryRow(`SELECT * FROM events WHERE id = ? LIMIT 1`, id)
	err := row.Scan(&event.Id, &event.Summary, &event.Description, &event.Course, &event.DueDate, &event.DueHour, &event.Done)
	if err != nil && err == sql.ErrNoRows {
		return nil, fmt.Errorf("the event with id %d does not exist", id)
	}

	return &event, nil
}

func (db *DB) Events() (*[]parser.CalendarEvent, error) {
	rows, err := db.sql.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []parser.CalendarEvent
	for rows.Next() {
		var event parser.CalendarEvent
		err := rows.Scan(&event.Id, &event.Summary, &event.Description, &event.Course, &event.DueDate, &event.DueHour, &event.Done)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}

	return &events, nil
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
