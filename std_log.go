// Standard log output

package glog

import (
	"log"
)

// Standard log function to log messages using the "log" package
//
// Parameters:
// - level: The log level
// - message: The log message
func StandardLogFunction(level int, message string) {
	switch level {
	case ERROR:
		log.Println("[ERROR] " + message)
	case WARNING:
		log.Println("[WARNING] " + message)
	case INFO:
		log.Println("[INFO] " + message)
	case DEBUG:
		log.Println("[DEBUG] " + message)
	case TRACE:
		log.Println("[TRACE] " + message)
	}
}
