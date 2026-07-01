// Package logging provides logging utilities for the Harness Go SDK.
//
// This package wraps the logrus logging library to provide consistent logging
// across all SDK packages with a standard format and configuration options.
//
// # Creating a Logger
//
// Create a new logger with default settings:
//
//	import "github.com/harness/harness-go-sdk/logging"
//
//	logger := logging.NewLogger()
//	logger.Info("Starting application")
//
// # Log Levels
//
// The logger supports standard log levels: Trace, Debug, Info, Warn, Error, Fatal.
// Check the current level:
//
//	if logging.IsDebugOrHigher(logger) {
//	    logger.Debug("Detailed debug information")
//	}
//
// # Leveled Logger Interface
//
// The package provides a LeveledLogger interface for dependency injection:
//
//	type LeveledLogger interface {
//	    Error(msg string, keysAndValues ...interface{})
//	    Warn(msg string, keysAndValues ...interface{})
//	    Info(msg string, keysAndValues ...interface{})
//	    Debug(msg string, keysAndValues ...interface{})
//	    IsDebugOrHigher() bool
//	}
//
// # HTTP Transport Logging
//
// Enable HTTP request/response logging for debugging API calls:
//
//	transport := logging.NewLoggingTransport(http.DefaultTransport, logger)
//	client := &http.Client{Transport: transport}
package logging
