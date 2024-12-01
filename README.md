# Utility library for logs (Go)

This a very simple utility library to create configurable loggers that can be nested.

[Documentation](https://pkg.go.dev/github.com/AgustinSRG/glog)

## Installation

To install the library in your project, run:

```sh
go get github.com/AgustinSRG/glog
```

## Usage

In the main function, create a root logger with the `CreateRootLogger(config)` function, passing the configuration as argument, which allows to enable or disable each specific level of log: `ERROR`, `WARNING`, `INFO`, `DEBUG`, `TRACE`.

You can then call the `CreateChildLogger(prefix)` on the root logger to create logger with the same configuration, but adding a prefix. For example, logs for the example function may have the `[ExampleFunction] ` prefix to better locate the logs.

Once you have the logger, you may call its methods to add logs.

Example:

```go
package main

import (
    // Import the module
    "github.com/AgustinSRG/glog"
)

func main() {
    // Create a root logger
	logger := glog.CreateRootLogger(glog.CreateLoggerConfigurationFromLevel(glog.INFO), glog.StandardLogFunction)

	// Log messages
	logger.Info("Example log message")

	// You can also log with format
	logger.Infof("Example log message: %v, %v, %v", 4, "example string", true)

	// If you log a level that is not enabled, no logs will be made
	logger.Debug("Example debug message")

	// You can also check the config to prevent running the function if disabled
	if logger.Config.DebugEnabled {
		logger.Debug("Example debug message")
	}

	// You can create child loggers, to include prefixes
	childLogger := logger.CreateChildLogger("[PREFIX] ")

	// This will log: [PREFIX] Example log message
	childLogger.Info("Example log message")
}
```

## Build the library

To install dependencies, run:

```sh
go get .
```

To build the code, run:

```sh
go build .
```

## Run linter

To run the code linter, run:

```sh
golangci-lint run
```

## Run tests

In order to run the tests for this library, run:

```sh
go test -v
```
