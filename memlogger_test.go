package mlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemLogger(t *testing.T) {
	logger := NewMemLogger()
	assert.Equal(t, "", logger.Name())
	assert.Equal(t, TraceLevel, logger.Level())

	sub := logger.New("sub")
	sub.SetLevel(ErrorLevel)
	assert.Equal(t, "", sub.Name())
	assert.Equal(t, TraceLevel, sub.Level())
}

func TestMemLogger_Log(t *testing.T) {
	logger := NewMemLogger()

	for _, lvl := range AllLevels {
		logger.Log(lvl, "msg")
		logger.Logf(lvl, "%s", "msg")
	}

	assert.Equal(t, len(AllLevels)*2, logger.LogCount())
	for _, lvl := range AllLevels {
		logs := logger.Logs(lvl)
		assert.Len(t, logs, 2)
		assert.Equal(t, "msg", logs[0])
		assert.Equal(t, "msg", logs[1])
	}
}

func TestMemLogger_Clear(t *testing.T) {
	logger := NewMemLogger()

	for _, lvl := range AllLevels {
		logger.Log(lvl, "msg")
	}

	logger.Clear()
	assert.Zero(t, logger.LogCount())

	for _, lvl := range AllLevels {
		logs := logger.Logs(lvl)
		assert.Empty(t, logs)
	}

	logger.AssertNoErrors(t)
	logger.AssertNoWarnings(t)
	logger.AssertNoLogs(t, AllLevels...)
}

func TestMemLogger_Assertions(t *testing.T) {
	logger := NewMemLogger()

	logger.Error("abc")
	logger.Error("def")
	logger.Error("ghi")

	logger.AssertAllMessages(t, ErrorLevel, "abc", "def", "ghi")
	logger.AssertAllSubMessages(t, ErrorLevel, "a", "e", "i")
	logger.AssertAnyMessage(t, ErrorLevel, "def")
	logger.AssertAnySubMessage(t, ErrorLevel, "e")
	logger.AssertNoLogs(t, WarnLevel, InfoLevel, DebugLevel, TraceLevel)
}
