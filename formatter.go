package mlog

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ConsoleFormatter(timestamp time.Time, level Level, name string, msg string) ([]byte, error) {
	ts := timestamp.Format("02.01.2006 15:04:05")
	lvl := strings.ToUpper(level.String())[:4]

	ts = colorCode(darkGray, ts)
	lvl = colorCode(levelColor(level), lvl)
	sym := colorCode(levelColor(level), "▶")

	name = trimPadLeft(name, 40)
	line := ts + " " + lvl + " " + name + " " + sym + " " + msg + "\n"
	return []byte(line), nil
}

func trimPadLeft(msg string, length int) string {
	msg = fmt.Sprintf("%"+strconv.Itoa(length)+"s", msg)
	return msg[len(msg)-length:]
}

const (
	red      = "31"
	yellow   = "33"
	cyan     = "36"
	gray     = "37"
	darkGray = "90"
)

func levelColor(level Level) string {
	switch level {
	case TraceLevel, DebugLevel:
		return gray
	case InfoLevel:
		return cyan
	case WarnLevel:
		return yellow
	case ErrorLevel:
		return red
	default: // unknown
		return red
	}
}

func colorCode(color, str string) string {
	return "\x1b[" + color + "m" + str + "\u001B[0m"
}
