// Test

package glog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	var lastLog string
	var lastLogLevel int
	var logger *Logger

	// ERROR

	logger = CreateRootLogger(LoggerConfiguration{
		ErrorEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Error("test message")

	assert.Equal(t, lastLog, "test message")
	assert.Equal(t, lastLogLevel, ERROR)

	logger = CreateRootLogger(LoggerConfiguration{
		ErrorEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Error("test message")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// ERROR (with format)

	logger = CreateRootLogger(LoggerConfiguration{
		ErrorEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Errorf("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, fmt.Sprintf("test message: %v, %v", 1, "test"))
	assert.Equal(t, lastLogLevel, ERROR)

	logger = CreateRootLogger(LoggerConfiguration{
		ErrorEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Errorf("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// WARNING

	logger = CreateRootLogger(LoggerConfiguration{
		WarningEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Warning("test message")

	assert.Equal(t, lastLog, "test message")
	assert.Equal(t, lastLogLevel, WARNING)

	logger = CreateRootLogger(LoggerConfiguration{
		WarningEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Warning("test message")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// WARNING (with format)

	logger = CreateRootLogger(LoggerConfiguration{
		WarningEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Warningf("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, fmt.Sprintf("test message: %v, %v", 1, "test"))
	assert.Equal(t, lastLogLevel, WARNING)

	logger = CreateRootLogger(LoggerConfiguration{
		WarningEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Warningf("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// INFO

	logger = CreateRootLogger(LoggerConfiguration{
		InfoEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Info("test message")

	assert.Equal(t, lastLog, "test message")
	assert.Equal(t, lastLogLevel, INFO)

	logger = CreateRootLogger(LoggerConfiguration{
		InfoEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Info("test message")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// INFO (with format)

	logger = CreateRootLogger(LoggerConfiguration{
		InfoEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Infof("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, fmt.Sprintf("test message: %v, %v", 1, "test"))
	assert.Equal(t, lastLogLevel, INFO)

	logger = CreateRootLogger(LoggerConfiguration{
		InfoEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Infof("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// DEBUG

	logger = CreateRootLogger(LoggerConfiguration{
		DebugEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Debug("test message")

	assert.Equal(t, lastLog, "test message")
	assert.Equal(t, lastLogLevel, DEBUG)

	logger = CreateRootLogger(LoggerConfiguration{
		DebugEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Debug("test message")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// DEBUG (with format)

	logger = CreateRootLogger(LoggerConfiguration{
		DebugEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Debugf("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, fmt.Sprintf("test message: %v, %v", 1, "test"))
	assert.Equal(t, lastLogLevel, DEBUG)

	logger = CreateRootLogger(LoggerConfiguration{
		DebugEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Debugf("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// TRACE

	logger = CreateRootLogger(LoggerConfiguration{
		TraceEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Trace("test message")

	assert.Equal(t, lastLog, "test message")
	assert.Equal(t, lastLogLevel, TRACE)

	logger = CreateRootLogger(LoggerConfiguration{
		TraceEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Trace("test message")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)

	// TRACE (with format)

	logger = CreateRootLogger(LoggerConfiguration{
		TraceEnabled: true,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Tracef("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, fmt.Sprintf("test message: %v, %v", 1, "test"))
	assert.Equal(t, lastLogLevel, TRACE)

	logger = CreateRootLogger(LoggerConfiguration{
		TraceEnabled: false,
	}, func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	lastLog = ""
	lastLogLevel = -1

	logger.Tracef("test message: %v, %v", 1, "test")

	assert.Equal(t, lastLog, "")
	assert.Equal(t, lastLogLevel, -1)
}

func TestCreateChildLogger(t *testing.T) {
	var lastLog string
	var lastLogLevel int

	logger := CreateRootLogger(CreateLoggerConfigurationFromLevel(TRACE), func(level int, msg string) {
		lastLog = msg
		lastLogLevel = level
	})

	childLogger := logger.CreateChildLogger("[PREFIX] ")

	lastLog = ""
	lastLogLevel = -1

	childLogger.Info("test message")

	assert.Equal(t, lastLog, "[PREFIX] test message")
	assert.Equal(t, lastLogLevel, INFO)

	grandChildLogger := childLogger.CreateChildLogger("[P2] ")

	lastLog = ""
	lastLogLevel = -1

	grandChildLogger.Info("test message")

	assert.Equal(t, lastLog, "[PREFIX] [P2] test message")
	assert.Equal(t, lastLogLevel, INFO)
}
