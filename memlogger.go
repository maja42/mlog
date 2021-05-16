package mlog

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/maja42/gotils"
	"github.com/maja42/gotils/compare"
	"github.com/maja42/gotils/testutil"
	"github.com/stretchr/testify/assert"
)

// MemLogger collects all log-messages instead of processing them.
// Intended for testing purposes to easily verify that certain messages were logged.
type MemLogger struct {
	m    sync.Mutex
	logs map[Level][]string
}

// NewMemLogger returns a new memory logger.
func NewMemLogger() *MemLogger {
	return &MemLogger{
		logs: make(map[Level][]string),
	}
}

func (l *MemLogger) New(...string) Logger {
	return l
}

func (l *MemLogger) Name() string {
	return ""
}

func (l *MemLogger) Level() Level {
	return TraceLevel
}

func (l *MemLogger) SetLevel(Level) {}

func (l *MemLogger) IsLevelEnabled(Level) bool {
	return true
}

func (l *MemLogger) WriterLevel(Level) io.WriteCloser {
	panic("not implemented yet")
}

func (l *MemLogger) Log(level Level, args ...interface{}) {
	msg := fmt.Sprint(args...)

	l.m.Lock()
	defer l.m.Unlock()
	l.logs[level] = append(l.logs[level], msg)
}

func (l *MemLogger) Logf(level Level, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	l.m.Lock()
	defer l.m.Unlock()
	l.logs[level] = append(l.logs[level], msg)
}

func (l *MemLogger) Trace(args ...interface{}) {
	l.Log(TraceLevel, args...)
}

func (l *MemLogger) Tracef(format string, args ...interface{}) {
	l.Logf(TraceLevel, format, args...)
}

func (l *MemLogger) Debug(args ...interface{}) {
	l.Log(DebugLevel, args...)
}

func (l *MemLogger) Debugf(format string, args ...interface{}) {
	l.Logf(DebugLevel, format, args...)
}

func (l *MemLogger) Info(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

func (l *MemLogger) Infof(format string, args ...interface{}) {
	l.Logf(InfoLevel, format, args...)
}

func (l *MemLogger) Warn(args ...interface{}) {
	l.Log(WarnLevel, args...)
}

func (l *MemLogger) Warnf(format string, args ...interface{}) {
	l.Logf(WarnLevel, format, args...)
}

func (l *MemLogger) Error(args ...interface{}) {
	l.Log(ErrorLevel, args...)
}

func (l *MemLogger) Errorf(format string, args ...interface{}) {
	l.Logf(ErrorLevel, format, args...)
}

func (l *MemLogger) Clear() {
	l.m.Lock()
	defer l.m.Unlock()
	l.logs = make(map[Level][]string)
}

func (l *MemLogger) LogCount() int {
	l.m.Lock()
	defer l.m.Unlock()
	cnt := 0
	for _, msgs := range l.logs {
		cnt += len(msgs)
	}
	return cnt
}

// Logs returns all log messages of a given level.
func (l *MemLogger) Logs(level Level) []string {
	l.m.Lock()
	defer l.m.Unlock()

	logs := make([]string, len(l.logs[level]))
	copy(logs, l.logs[level])
	return logs
}

// TraceLogs returns a copy of all trace logs.
func (l *MemLogger) TraceLogs() []string {
	return l.Logs(TraceLevel)
}

// DebugLogs returns a copy of all debug logs.
func (l *MemLogger) DebugLogs() []string {
	return l.Logs(DebugLevel)
}

// InfoLogs returns a copy of all info logs.
func (l *MemLogger) InfoLogs() []string {
	return l.Logs(InfoLevel)
}

// WarnLogs returns a copy of all warn logs.
func (l *MemLogger) WarnLogs() []string {
	return l.Logs(WarnLevel)
}

// ErrorLogs returns a copy of all error logs.
func (l *MemLogger) ErrorLogs() []string {
	return l.Logs(ErrorLevel)
}

// AssertNoLogs verifies that there are no log messages with the given level(s).
// Requires at least one level. Pass AllLevels to verify that no log-statements were made.
func (l *MemLogger) AssertNoLogs(t testutil.TestingT, levels ...Level) bool {
	t.Helper()
	ok := true

	if len(levels) == 0 { // prevent wrong API usage
		panic("invalid usage: no levels provided")
	}

	l.m.Lock()
	defer l.m.Unlock()

	for _, level := range levels {
		logs := l.logs[level]
		if len(logs) == 0 {
			continue
		}

		ok = false
		t.Errorf("Expected no %s logs, but got %d message(s):\n%s",
			level, len(logs), gotils.Indent(strings.Join(logs, "\n")))
	}
	return ok
}

// AssertNoErrors verifies that there were no errors logged.
func (l *MemLogger) AssertNoErrors(t testutil.TestingT) bool {
	t.Helper()
	return l.AssertNoLogs(t, ErrorLevel)
}

// AssertNoWarnings verifies that there were no warnings logged.
func (l *MemLogger) AssertNoWarnings(t testutil.TestingT) bool {
	t.Helper()
	return l.AssertNoLogs(t, ErrorLevel)
}

// AssertAllMessages verifies that the given messages were logged.
// Reports an error if additional messages were logged, or some expected messages are missing.
// The message order is ignored.
func (l *MemLogger) AssertAllMessages(t testutil.TestingT, level Level, messages ...string) bool {
	t.Helper()
	l.m.Lock()
	defer l.m.Unlock()

	msg := fmt.Sprintf("Unexpected %s log messages", level)
	return testutil.AssertElementsMatch(t, msg, messages, l.logs[level], compare.String)
}

// AssertAllSubMessages verifies that the given messages were logged.
// Instead of matching the whole message, the expected subMessages must be a substring.
// Reports an error if additional messages were logged, or some expected messages are missing.
// The message order is ignored.
func (l *MemLogger) AssertAllSubMessages(t testutil.TestingT, level Level, subMessages ...string) bool {
	t.Helper()
	l.m.Lock()
	defer l.m.Unlock()

	msg := fmt.Sprintf("Unexpected %s log messages", level)
	return testutil.AssertElementsMatch(t, msg, subMessages, l.logs[level], compare.SubString)
}

// AssertAnyMessage verifies that the given message was logged at least once.
func (l *MemLogger) AssertAnyMessage(t testutil.TestingT, level Level, message string) bool {
	t.Helper()
	l.m.Lock()
	defer l.m.Unlock()

	return assert.Contains(t, l.logs[level], message)
}

// AssertAnySubMessage verifies that the given subMessage was a sub-string of at least one log message.
func (l *MemLogger) AssertAnySubMessage(t testutil.TestingT, level Level, subMessage string) bool {
	t.Helper()
	l.m.Lock()
	defer l.m.Unlock()

	for _, msg := range l.logs[level] {
		if strings.Contains(msg, subMessage) {
			return true
		}
	}
	t.Errorf("%#v does not contain message with %q", l.logs[level], subMessage)
	return false
}
