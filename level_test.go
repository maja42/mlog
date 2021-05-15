package mlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LessSevereThan(t *testing.T) {
	assert.True(t, TraceLevel.LessSevereThan(DebugLevel))
	assert.True(t, DebugLevel.LessSevereThan(ErrorLevel))

	assert.False(t, WarnLevel.LessSevereThan(WarnLevel))
	assert.False(t, WarnLevel.LessSevereThan(InfoLevel))
}

func Test_ParseLevel(t *testing.T) {
	assert.Equal(t, TraceLevel, mustParseLevel(t, "trace"))
	assert.Equal(t, DebugLevel, mustParseLevel(t, "debug"))
	assert.Equal(t, InfoLevel, mustParseLevel(t, "info"))
	assert.Equal(t, WarnLevel, mustParseLevel(t, "warn"))
	assert.Equal(t, WarnLevel, mustParseLevel(t, "warning"))
	assert.Equal(t, ErrorLevel, mustParseLevel(t, "error"))

	assert.Equal(t, ErrorLevel, mustParseLevel(t, "ErRoR"))

	_, err := ParseLevel("")
	assert.Error(t, err)

	_, err = ParseLevel(" error")
	assert.Error(t, err)

	_, err = ParseLevel("e")
	assert.Error(t, err)
}

func mustParseLevel(t *testing.T, level string) Level {
	lvl, err := ParseLevel(level)
	assert.NoError(t, err)
	return lvl
}
