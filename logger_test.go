package mlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_appendLoggerNameComponents(t *testing.T) {
	assert.Equal(t, "", appendLoggerNameComponents(""))
	assert.Equal(t, "a", appendLoggerNameComponents("", "a"))
	assert.Equal(t, "a.b", appendLoggerNameComponents("", "a", "b"))
	assert.Equal(t, "a.b", appendLoggerNameComponents("a", "b"))
	assert.Equal(t, "a.b.c", appendLoggerNameComponents("a.b", "c"))
	assert.Equal(t, "a.b.c.d", appendLoggerNameComponents("a.b", "c", "d"))
}
