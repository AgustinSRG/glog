// Configuration

package glog

// Logger configuration
type LoggerConfiguration struct {
	// True to enable ERROR messages
	ErrorEnabled bool

	// True to enable WARNING messages
	WarningEnabled bool

	// True to enable INFO messages
	InfoEnabled bool

	// True to enable DEBUG messages
	DebugEnabled bool

	// True to enable TRACE messages
	TraceEnabled bool
}

// Creates a configuration from a level
//
// Parameters:
// - level: The lowest level to log
//
// Returns a configuration with all the levels lower then the specified disabled
func CreateLoggerConfigurationFromLevel(level int) LoggerConfiguration {
	return LoggerConfiguration{
		ErrorEnabled:   level >= ERROR,
		WarningEnabled: level >= WARNING,
		InfoEnabled:    level >= INFO,
		DebugEnabled:   level >= DEBUG,
		TraceEnabled:   level >= TRACE,
	}
}
