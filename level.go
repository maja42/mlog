package mlog

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// Level represents the severity of a log message
type Level logrus.Level

const (
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = Level(logrus.ErrorLevel)
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = Level(logrus.WarnLevel)
	// InfoLevel level. General operational entries about what's going on inside the application.
	InfoLevel = Level(logrus.InfoLevel)
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = Level(logrus.DebugLevel)
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel = Level(logrus.TraceLevel)
)

// AllLevels exposes all logging levels
var AllLevels = []Level{
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

var levelStrings = map[Level]string{
	ErrorLevel: "error",
	WarnLevel:  "warn",
	InfoLevel:  "info",
	DebugLevel: "debug",
	TraceLevel: "trace",
}

func (l Level) String() string {
	str, ok := levelStrings[l]
	if !ok {
		return "unknown"
	}
	return str
}

// LessSevereThan returns true of the level is less severe than onother log level.
func (l Level) LessSevereThan(other Level) bool {
	return l > other
}

// ParseLevel takes a string and returns the matching log level.
func ParseLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	case "trace":
		return TraceLevel, nil
	}

	var zeroVal Level
	return zeroVal, fmt.Errorf("invalid log level: %q", level)
}
