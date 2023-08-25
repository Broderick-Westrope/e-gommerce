package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"text/tabwriter"
)

// Logger is an interface for logging.
type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
}

// Slog is a wrapper around slog.Logger.
type Slog struct {
	log *slog.Logger
}

// NewSlog returns a new Slog using slog.NewJSONHandler.
func NewSlog() *Slog {
	return &Slog{slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

// Debug logs a debug message.
func (s *Slog) Debug(msg string, args ...interface{}) {
	s.log.Debug(msg, args...)
}

// Info logs an info message.
func (s *Slog) Info(msg string, args ...interface{}) {
	s.log.Info(msg, args...)
}

// Warn logs a warning message.
func (s *Slog) Warn(msg string, args ...interface{}) {
	s.log.Warn(msg, args...)
}

// Error logs an error message.
func (s *Slog) Error(msg string, args ...interface{}) {
	s.log.Error(msg, args...)
}

// Log is a wrapper around log.Logger.
type Log struct {
	*log.Logger
}

// NewLog returns a new Log using log.New.
func NewLog() *Log {
	return &Log{log.Default()}
}

// Debug logs a debug message.
func (l *Log) Debug(msg string, args ...interface{}) {
	l.createLog("DEBUG", msg, args...)
}

// Info logs an info message.
func (l *Log) Info(msg string, args ...interface{}) {
	l.createLog("INFO", msg, args...)
}

// Warn logs a warning message.
func (l *Log) Warn(msg string, args ...interface{}) {
	l.createLog("WARN", msg, args...)
}

// Error logs an error message.
func (l *Log) Error(msg string, args ...interface{}) {
	l.createLog("ERR", msg, args...)
}

func (l *Log) createLog(level, msg string, args ...interface{}) {
	// Initialize a tabwriter for aligned columns.
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 6, 2, ' ', tabwriter.TabIndent)

	// Loop through all args and format them as key-value pairs.
	for i := 0; i < len(args); i += 2 {
		fmt.Fprintf(w, ",\t%v: %v", args[i], args[i+1])
	}

	// Flush the tabwriter.
	w.Flush()

	msg = fmt.Sprintf("[%s] %s%s", level, msg, buf.String())
	l.Println(msg)
}
