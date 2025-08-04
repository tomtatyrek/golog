package main

import (
	"os"
	"time"

	"github.com/tomtatyrek/golog"
)

func main() {

	// Creates a custom clock with one custom timestamp layout and one
	// predefined in time package
	clock := golog.NewClock("15-04-05", time.RFC1123)

	// Defines a slice of files to which the log will be written.
	// (In this case only one file is provided.)
	files := []*os.File{os.Stdout}

	// Defines which logging levels will be logged by applying bitwise or to them
	allowedLogLevels := golog.FATAL | golog.ERROR | golog.INFO

	// Creates a custom Logger using above-defined variables
	logger := golog.NewLogger(files, clock, allowedLogLevels)

	// It is recommended to always defer closing the logger in the same place
	// where it was created, especially if it writes to actual files.
	defer logger.Close()

	logger.Error("This message will be shown")
	logger.Trace("This message won't be shown")

}
