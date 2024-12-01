// Logger

package glog

import "fmt"

// Logger
type Logger struct {
	// Logger configuration
	Config LoggerConfiguration

	// Prefix for logs
	Prefix string

	// Log function
	LogFunc func(level int, message string)
}

// Creates a root logger
//
// Parameters:
// - config: The logger configuration
// - logFunc: The function to log the messages. By default, use glog.StandardLogFunction
//
// Returns a Logger instance
func CreateRootLogger(config LoggerConfiguration, logFunc func(level int, message string)) *Logger {
	return &Logger{
		Config:  config,
		Prefix:  "",
		LogFunc: logFunc,
	}
}

// Creates child logger, adding a prefix to the logs of the parent logger
//
// Parameters:
// - prefix: The prefix for the logs
//
// Returns a Logger instance
func (logger *Logger) CreateChildLogger(prefix string) *Logger {
	return &Logger{
		Config:  logger.Config,
		Prefix:  logger.Prefix + prefix,
		LogFunc: logger.LogFunc,
	}
}

// Logs a message
//
// Parameters:
// - level: Log level
// - message: The message
func (logger *Logger) Log(level int, message string) {
	if logger.LogFunc != nil {
		logger.LogFunc(level, logger.Prefix+message)
	}
}

// Logs an ERROR message
//
// Parameters:
// - message: The message
func (logger *Logger) Error(message string) {
	if !logger.Config.ErrorEnabled {
		return
	}

	logger.Log(ERROR, message)
}

// Logs an ERROR message (with format)
//
// Parameters:
// - format: The message format
// - a: The parameters to be used in the format
//
// Example:
//  logger.Errorf("An error happened in the request to %v | %v", url, error.Error())
func (logger *Logger) Errorf(format string, a ...any) {
	logger.Error(fmt.Sprintf(format, a...))
}

// Logs an WARNING message
//
// Parameters:
// - message: The message
func (logger *Logger) Warning(message string) {
	if !logger.Config.WarningEnabled {
		return
	}

	logger.Log(WARNING, message)
}

// Logs an WARNING message (with format)
//
// Parameters:
// - format: The message format
// - a: The parameters to be used in the format
//
// Example:
//  logger.Warningf("Unknown value: %v", value)
func (logger *Logger) Warningf(format string, a ...any) {
	logger.Warning(fmt.Sprintf(format, a...))
}

// Logs an INFO message
//
// Parameters:
// - message: The message
func (logger *Logger) Info(message string) {
	if !logger.Config.InfoEnabled {
		return
	}

	logger.Log(INFO, message)
}

// Logs an INFO message (with format)
//
// Parameters:
// - format: The message format
// - a: The parameters to be used in the format
//
// Example:
//  logger.Infof("Received request from: %v", ip)
func (logger *Logger) Infof(format string, a ...any) {
	logger.Info(fmt.Sprintf(format, a...))
}

// Logs an DEBUG message
//
// Parameters:
// - message: The message
func (logger *Logger) Debug(message string) {
	if !logger.Config.DebugEnabled {
		return
	}

	logger.Log(DEBUG, message)
}

// Logs an DEBUG message (with format)
//
// Parameters:
// - format: The message format
// - a: The parameters to be used in the format
//
// Example:
//  logger.Debugf("The value of the intermediate value is: %v", value)
func (logger *Logger) Debugf(format string, a ...any) {
	logger.Debug(fmt.Sprintf(format, a...))
}

// Logs an TRACE message
//
// Parameters:
// - message: The message
func (logger *Logger) Trace(message string) {
	if !logger.Config.TraceEnabled {
		return
	}

	logger.Log(TRACE, message)
}

// Logs an TRACE message (with format)
//
// Parameters:
// - format: The message format
// - a: The parameters to be used in the format
//
// Example:
//  logger.Tracef("DB Query: %v", query)
func (logger *Logger) Tracef(format string, a ...any) {
	logger.Trace(fmt.Sprintf(format, a...))
}
