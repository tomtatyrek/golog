package golog

import (
	"fmt"
	"log"
	"os"
)

// A Logger is used to log messages in according to its configuration.
type Logger struct {
	files            []*os.File
	clock            *Clock
	allowedLogLevels LogLevel
}

// NewLogger returns a new [Logger] created from [os.File] slice, a [Clock] and
// a [LogLevel], is the result of applying bitwise or to all of the logging levels,
// which are allowed to be logged.
//
// The files have to already be open. It is to be followed by a defer call to
// [Close()] in order to close all of the Logger's files.
//
// [Close()]: https://pkg.go.dev/github.com/tomtatyrek/golog#Logger.Close
func NewLogger(files []*os.File, clock *Clock, allowedLogLevels LogLevel) *Logger {
	logger := Logger{files, clock, allowedLogLevels}
	// Default starting message
	m := fmt.Sprintf("==== New Logger created on %s at %s ====\n\n", clock.nowDate(), clock.nowTime())
	logger.log(m)
	return &logger
}

// NewDefaultLogger returns a [Logger] that logs to [os.Stdout],
// uses default time and date formats and shows default log levels (all except [TRACE]).
func NewDefaultLogger() *Logger {
	return NewLogger([]*os.File{os.Stdout}, NewDefaultClock(), DefaultLogLevels)
}

func (l *Logger) log(s string) {
	for i := range len(l.files) {
		_, err := l.files[i].WriteString(s)
		check(err)
	}
}

// Close closes all the files that belong to a [Logger].
// It is to be always used with defer when creating a new [Logger].
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
