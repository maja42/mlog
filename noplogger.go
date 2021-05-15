package mlog

import "io"

type nopLogger struct {
}

// NOPLogger is a stub that does perform any logging operations.
// Useful for testing to disable logging completely.
var NOPLogger nopLogger

func (nopLogger) New(...string) Logger {
	return nopLogger{}
}

func (nopLogger) Name() string {
	return ""
}

func (nopLogger) Level() Level {
	var zeroLevel Level
	return zeroLevel
}

func (nopLogger) SetLevel(Level) {}

func (nopLogger) WriterLevel(Level) io.WriteCloser {
	return nopCloserWriter{}
}

type nopCloserWriter struct{}

func (nopCloserWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (nopCloserWriter) Close() error {
	return nil
}

func (nopLogger) Log(Level, ...interface{})          {}
func (nopLogger) Logf(Level, string, ...interface{}) {}

func (nopLogger) Trace(...interface{})          {}
func (nopLogger) Tracef(string, ...interface{}) {}

func (nopLogger) Debug(...interface{})          {}
func (nopLogger) Debugf(string, ...interface{}) {}

func (nopLogger) Info(...interface{})          {}
func (nopLogger) Infof(string, ...interface{}) {}

func (nopLogger) Warn(...interface{})          {}
func (nopLogger) Warnf(string, ...interface{}) {}

func (nopLogger) Error(...interface{})          {}
func (nopLogger) Errorf(string, ...interface{}) {}
