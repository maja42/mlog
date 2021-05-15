package mlog

import (
	"io"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var newLine string // newLine on the current platform
func init() {
	switch runtime.GOOS {
	case "windows":
		newLine = "\r\n"
	default:
		newLine = "\n"
	}
}

type LoggerBuilder struct {
	logger *StdLogger
}

func NewLoggerBuilder(name ...string) *LoggerBuilder {
	return &LoggerBuilder{
		logger: NewStdLogger(name...),
	}
}

func (f *LoggerBuilder) WithOutput(w io.Writer) *LoggerBuilder {
	f.logger.rawLogger().Out = w
	return f
}

func (f *LoggerBuilder) WithLevel(level Level) *LoggerBuilder {
	f.logger.SetLevel(level)
	return f
}

// The Formatter converts a log message into a string suitable for console output.
type Formatter func(timestamp time.Time, level Level, name string, msg string) ([]byte, error)

func (f *LoggerBuilder) WithConsoleFormatter(formatter Formatter) *LoggerBuilder {
	f.logger.rawLogger().SetFormatter(&logrusFormatterAdapter{
		formatter: formatter,
	})
	return f
}

// Hook is called for outgoing log messages.
// Note: hooks are not called asynchronously and should therefore be non-blocking.
type Hook func(timestamp time.Time, level Level, name string, msg string) error

func (f *LoggerBuilder) WithHook(levels []Level, hook Hook) *LoggerBuilder {
	adapter := newLogrusHookAdapter(levels, hook)
	f.logger.rawLogger().AddHook(adapter)
	return f
}

func (f *LoggerBuilder) Create() Logger {
	return f.logger.Clone()
}

type logrusFormatterAdapter struct {
	formatter Formatter
}

func (a *logrusFormatterAdapter) Format(entry *logrus.Entry) ([]byte, error) {
	name := entry.Data["name"].(string)
	return a.formatter(entry.Time, Level(entry.Level), name, entry.Message)
}

type logrusHookAdapter struct {
	levels []logrus.Level
	hook   Hook
}

func newLogrusHookAdapter(levels []Level, hook Hook) *logrusHookAdapter {
	logrusLevels := make([]logrus.Level, len(levels))
	for i, lvl := range levels {
		logrusLevels[i] = logrus.Level(lvl)
	}
	return &logrusHookAdapter{
		levels: logrusLevels,
		hook:   hook,
	}
}

func (a *logrusHookAdapter) Levels() []logrus.Level {
	return a.levels
}

func (a *logrusHookAdapter) Fire(entry *logrus.Entry) error {
	name := entry.Data["name"].(string)
	return a.hook(entry.Time, Level(entry.Level), name, entry.Message)
}
