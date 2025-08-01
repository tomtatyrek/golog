package golog

import (
	"fmt"
	"log"
	"os"
)

// Logger is used to log messages according to its configuration
type Logger struct {
	files            []*os.File
	clock            *Clock
	allowedLogLevels LogLevel
}

// NewLogger returns a new logger created from [os.File] slice and a [Clock].
// The files have to already be open
func NewLogger(files []*os.File, clock *Clock, allowedLogLevels LogLevel) *Logger {
	logger := Logger{files, clock, allowedLogLevels}
	// Default starting message
	m := fmt.Sprintf("==== New Logger created on %s at %s ====\n\n", clock.nowDate(), clock.nowTime())
	logger.log(m)
	return &logger
}

// NeDefaultLogger returns a logger that logs to [os.Stdout],
// uses default time and date formats and shows default log levels
func NewDefaultLogger() *Logger {
	return NewLogger([]*os.File{os.Stdout}, NewDefaultClock(), DefaultLogLevels)
}

func (l *Logger) log(s string) {
	for i := range len(l.files) {
		_, err := l.files[i].WriteString(s)
		check(err)
	}
}

func (l *Logger) Close() {
	for i := range len(l.files) {
		err := l.files[i].Close()
		check(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
