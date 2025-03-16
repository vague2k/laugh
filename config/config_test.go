package config_test

import (
	"path/filepath"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"github.com/vague2k/laugh/config"
)

func TestConfig(t *testing.T) {
	t.Run("Create config dir if it doesn't exist", func(t *testing.T) {
		// first, assert that the temp dir does NOT exist
		temp := t.TempDir()
		expected := filepath.Join(temp, "laugh")
		assert.NoDirExists(t, expected)

		conf, err := config.LoadConfig(temp)
		assert.NoError(t, err)
		assert.NotNil(t, conf)

		// finally, assert that the temp dir DOES exist
		actual := conf.Dir()
		assert.DirExists(t, expected)
		assert.Equal(t, expected, actual)
	})

	t.Run("Create config file if it doesn't exist", func(t *testing.T) {
		// first, assert that the temp config file does NOT exist
		dir := t.TempDir()
		expected := filepath.Join(dir, "laugh", "laugh.toml")
		assert.NoFileExists(t, expected)

		conf, err := config.LoadConfig(dir)
		assert.NoError(t, err)
		assert.NotNil(t, conf)

		// finally, assert that the temp config file DOES exist
		actual := conf.Filename()
		assert.FileExists(t, expected)
		assert.Equal(t, expected, actual)
	})

	t.Run("Has expected default", func(t *testing.T) {
		dir := t.TempDir()
		conf, err := config.LoadConfig(dir)
		assert.NoError(t, err)
		assert.NotNil(t, conf)

		actual := config.Config{}
		expected := config.Config{
			Url: "the url to your student canvas calendar should go here",
		}

		// decode file contents into actual
		metadata, err := toml.DecodeFile(conf.Filename(), &actual)
		assert.NoError(t, err)
		assert.True(t, metadata.IsDefined("Url"))
		assert.Equal(t, expected, actual)
	})
}
