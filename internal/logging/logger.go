// internal/logging/logger.go
package logging

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

// InitLogger initializes the global logger
func InitLogger(level string, pretty bool) {
	InitLoggerWithFile(level, pretty, "")
}

// InitLoggerWithFile initializes the global logger with optional file output
func InitLoggerWithFile(level string, pretty bool, logFile string) {
	// Set appropriate log level
	var logLevel zerolog.Level
	switch level {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	// Configure output
	var output io.Writer = os.Stderr

	// If log file is specified, use it instead of stderr
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open log file %s: %v\n", logFile, err)
			output = os.Stderr
		} else {
			output = file
		}
	}

	// Use pretty logging for human-readable output if requested
	if pretty && logFile == "" {
		output = zerolog.ConsoleWriter{
			Out:        output,
			TimeFormat: time.RFC3339,
			FormatLevel: func(i interface{}) string {
				return fmt.Sprintf("| %-6s|", i)
			},
		}
	}

	// Create the logger
	zerolog.SetGlobalLevel(logLevel)
	logger = zerolog.New(output).With().Timestamp().Caller().Logger()

	log.Logger = logger

	logger.Debug().Msg("Logger initialized")
}

// GetLogger returns the configured logger
func GetLogger() zerolog.Logger {
	return logger
}
