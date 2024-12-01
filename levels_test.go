// Test

package glog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelToString(t *testing.T) {
	assert.Equal(t, LevelToString(SILENT), "SILENT")
	assert.Equal(t, LevelToString(ERROR), "ERROR")
	assert.Equal(t, LevelToString(WARNING), "WARNING")
	assert.Equal(t, LevelToString(INFO), "INFO")
	assert.Equal(t, LevelToString(DEBUG), "DEBUG")
	assert.Equal(t, LevelToString(TRACE), "TRACE")
	assert.Equal(t, LevelToString(-1), "")
}

func TestLevelFromString(t *testing.T) {
	l, err := LevelFromString("SILENT")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, l, SILENT)

	l, err = LevelFromString("ERROR")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, l, ERROR)

	l, err = LevelFromString("error")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, l, ERROR)

	l, err = LevelFromString("WARNING")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, l, WARNING)

	l, err = LevelFromString("INFO")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, l, INFO)

	l, err = LevelFromString("DEBUG")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, l, DEBUG)

	l, err = LevelFromString("TRACE")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, l, TRACE)

	_, err = LevelFromString("invalid")
	assert.NotEqual(t, err, nil)
}
