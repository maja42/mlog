package mlog

import (
	"log"
)

// RedirectStdLogger bridges log messages from go's standard log-package
// to the provided logger. Overwrites any previous redirections.
func RedirectStdLogger(logger Logger) {
	log.SetOutput(logger.WriterLevel(WarnLevel))
}
