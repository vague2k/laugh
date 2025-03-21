package database_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vague2k/laugh/database"
)

func TestInit(t *testing.T) {
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

// A helper function to spin up a db in memory for quick testing
//
// This function already includes a cleanup function where when the test
// completes, the database is closed
// func mem(t *testing.T) *database.DB {
// 	db, err := database.Init(":memory:")
// 	assert.NoError(t, err)
// 	t.Cleanup(func() {
// 		db.Close()
// 		db = nil
// 	})
// 	return db
// }
