package mlog

import (
	"log"
)

// RedirectStdLogger bridges log messages from go's standard log-package to the provided logger.
// Disables date/time output of standard logger. Overwrites any previous redirections.
func RedirectStdLogger(logger Logger) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime)) // Don't log date and time
	log.SetOutput(logger.WriterLevel(WarnLevel))
}
