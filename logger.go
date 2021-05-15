package mlog

import (
	"io"
	"strings"
)

const nameSeparator = "."

// Logger receives and processes log messages.
type Logger interface {
	// New returns a new sub-logger with the given name components appended
	New(name ...string) Logger

	// Name returns the logger's full name
	Name() string

	// Level returns the logger's log level
	Level() Level

	// SetLevel changes the logger's log level.
	// Does not modify the level of parent- or sub-loggers.
	SetLevel(level Level)

	// WriterLevel returns a new writer.
	// All lines written to it will be logged with the given log level.
	WriterLevel(level Level) io.WriteCloser

	Log(level Level, args ...interface{})
	Logf(level Level, format string, args ...interface{})

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

func appendLoggerNameComponents(name string, newNames ...string) string {
	if len(newNames) == 0 {
		return name
	}
	if len(name) > 0 {
		name = name + nameSeparator
	}
	return name + strings.Join(newNames, nameSeparator)
}
