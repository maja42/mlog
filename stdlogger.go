package mlog

import (
	"io"

	"github.com/sirupsen/logrus"
)

type StdLogger struct {
	name     string
	rawEntry *logrus.Entry
}

// NewStdLogger returns a new standard logger.
func NewStdLogger(name ...string) *StdLogger {
	loggerName := appendLoggerNameComponents("", name...)
	logger := &StdLogger{
		name:     loggerName,
		rawEntry: newLogrusEntry(logrus.New(), loggerName),
	}
	return logger
}

func newLogrusEntry(logger *logrus.Logger, loggerName string) *logrus.Entry {
	return &logrus.Entry{
		Logger: logger,
		Data: map[string]interface{}{
			"name": loggerName,
		},
	}
}

func (l *StdLogger) rawLogger() *logrus.Logger {
	return l.rawEntry.Logger
}

// New returns a new sub-logger.
func (l *StdLogger) New(name ...string) Logger {
	loggerName := appendLoggerNameComponents(l.name, name...)

	rawLogger := copyLogrusLogger(l.rawLogger())

	newLogger := &StdLogger{
		name:     loggerName,
		rawEntry: newLogrusEntry(rawLogger, loggerName),
	}
	return newLogger
}

func copyLogrusLogger(old *logrus.Logger) *logrus.Logger {
	clone := logrus.New()
	clone.Out = old.Out
	clone.Hooks = old.Hooks
	clone.Formatter = old.Formatter
	clone.Level = old.Level
	return clone
}

// Clone returns a copy of this logger
func (l *StdLogger) Clone() Logger {
	return l.New()
}

func (l *StdLogger) Name() string {
	return l.name
}

func (l *StdLogger) Level() Level {
	return Level(l.rawLogger().Level)
}

func (l *StdLogger) IsLevelEnabled(level Level) bool {
	return l.rawLogger().IsLevelEnabled(logrus.Level(level))
}

func (l *StdLogger) SetLevel(level Level) {
	l.rawLogger().Level = logrus.Level(level)
}

func (l *StdLogger) WriterLevel(level Level) io.WriteCloser {
	return l.rawEntry.WriterLevel(logrus.Level(level))
}

func (l *StdLogger) Log(level Level, args ...interface{}) {
	l.rawEntry.Log(logrus.Level(level), args...)
}

func (l *StdLogger) Logf(level Level, format string, args ...interface{}) {
	l.rawEntry.Logf(logrus.Level(level), format, args...)
}

func (l *StdLogger) Trace(args ...interface{}) {
	l.rawEntry.Trace(args...)
}

func (l *StdLogger) Tracef(format string, args ...interface{}) {
	l.rawEntry.Tracef(format, args...)
}

func (l *StdLogger) Debug(args ...interface{}) {
	l.rawEntry.Debug(args...)
}

func (l *StdLogger) Debugf(format string, args ...interface{}) {
	l.rawEntry.Debugf(format, args...)
}

func (l *StdLogger) Info(args ...interface{}) {
	l.rawEntry.Info(args...)
}

func (l *StdLogger) Infof(format string, args ...interface{}) {
	l.rawEntry.Infof(format, args...)
}

func (l *StdLogger) Warn(args ...interface{}) {
	l.rawEntry.Warn(args...)
}

func (l *StdLogger) Warnf(format string, args ...interface{}) {
	l.rawEntry.Warnf(format, args...)
}

func (l *StdLogger) Error(args ...interface{}) {
	l.rawEntry.Error(args...)
}

func (l *StdLogger) Errorf(format string, args ...interface{}) {
	l.rawEntry.Errorf(format, args...)
}

func (l *StdLogger) AddHook(levels []Level, hook Hook) *StdLogger {
	adapter := newLogrusHookAdapter(levels, hook)
	l.rawLogger().AddHook(adapter)
	return l
}
