// Log levels

package glog

import (
	"errors"
	"strings"
)

// Silent level (no logs)
const SILENT int = 0

// Error level
const ERROR int = 1

// Warning level
const WARNING int = 2

// Info level
const INFO int = 3

// Debug level
const DEBUG int = 4

// Trace level
const TRACE int = 5

// Turns level from int to string
//
// Parameters:
// - level: The level
//
// Returns the level as string
func LevelToString(level int) string {
	switch level {
	case SILENT:
		return "SILENT"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	default:
		return ""
	}
}

// Parses level from string (case insensitive)
//
// Parameters:
// - str: The string to parse
//
// Returns the parsed level, or an error if the value has an invalid format
func LevelFromString(str string) (int, error) {
	switch strings.ToUpper(str) {
	case "SILENT":
		return SILENT, nil
	case "ERROR":
		return ERROR, nil
	case "WARNING":
		return WARNING, nil
	case "INFO":
		return INFO, nil
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	default:
		return -1, errors.New("unrecognized log level")
	}
}
