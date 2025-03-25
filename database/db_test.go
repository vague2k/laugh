package database_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vague2k/laugh/database"
	"github.com/vague2k/laugh/parser"
)

func TestNew(t *testing.T) {
	temp := t.TempDir()
	db, err := database.New(temp)
	assert.NoError(t, err)

	expectedDir := filepath.Join(temp, "laugh")
	expectedDBFile := filepath.Join(expectedDir, "laugh.db")

	assert.NotNil(t, db)
	assert.NotEmpty(t, db.Dir())
	assert.NotEmpty(t, db.Filename())
	assert.Equal(t, expectedDir, db.Dir())
	assert.Equal(t, expectedDBFile, db.Filename())
}

func TestAddEvent(t *testing.T) {
	db := mem(t)

	actual := &parser.CalendarEvent{
		Id:          1,
		DueDate:     "duedate",
		DueHour:     "duehour",
		Summary:     "summary",
		Description: "description",
		Course:      "course",
		Done:        false,
	}

	t.Run("Adds event with no error", func(t *testing.T) {
		err := db.AddEvent(actual)
		assert.NoError(t, err)
	})
	t.Run("Assert added event exists", func(t *testing.T) {
		expected, err := db.SelectItem(1)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestEvents(t *testing.T) {
	db := mem(t)
	for i := range 10 {
		event := &parser.CalendarEvent{Id: i}
		err := db.AddEvent(event)
		assert.NoError(t, err)
	}

	t.Run("Returns all events in DB", func(t *testing.T) {
		events, err := db.Events()
		assert.NoError(t, err)
		assert.NotNil(t, events)
		assert.Equal(t, len(*events), 10)
	})

	t.Run("Assert fields are populated", func(t *testing.T) {
		for i := range 10 {
			event, err := db.SelectItem(i)
			assert.NoError(t, err)
			assert.NotNil(t, event)
			assert.Equal(t, event.Id, i)
		}
	})
}

// A helper function to spin up a db in memory for quick testing
//
// This function already includes a cleanup function where when the test
// completes, the database is closed
func mem(t *testing.T) *database.DB {
	db, err := database.New(":memory:")
	assert.NoError(t, err)
	t.Cleanup(func() {
		db.Close()
		db = nil
	})
	return db
}
