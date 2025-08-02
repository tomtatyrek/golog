/*
Package golog provides simple logging functionality.

The golog package offers a lightweight and easy-to-use logging interface for Go applications.
It supports logging messages at different levels (such as Info, Warning, Error), can be
configured to output logs to various destinations, such as standard output or files, and gives
the user the ability to change timestamp formats.

# Concurrency

From my testing Golog should be safe for concurrent use as it mainly depends on
[os.File.WriteString()] function, which according to [its documentation] is ok
to be used concurrently, as long as the system limit for writing to files,
should be quite hign, isn't exceeded.

# Examples

Typical usage:

	import "github.com/tomtatyrek/golog"

	func main() {
		// Creates a default logger
		logger := golog.NewDefaultLogger()
		defer logger.Close()

		// Logs a few messages with different logging levels
		logger.Warn("This is a warning")
		logger.Error("An error occurred")
		logger.Infof("The remainder of dividing 5 by 2 is %v", 5%2)
	}

This will result in the following output:

==== New Logger created on Fri, 01 Aug 2025 21:36:06 CEST at 21:36:06.687 ====

[21:36:06.687] [WARN] This is a warning

[21:36:06.687] [ERROR] An error occurred

[21:36:06.687] [INFO] The remainder of dividing 5 by 2 is 1

More advanced usage:

	import (
		"os"
		"time"

		"github.com/tomtatyrek/golog"
	)

	func main() {
		// Creates a custom clock with one custom timestamp layout and one
		// predefined in time package
		clock := golog.NewClock("15-04-05", time.RFC1123)

		// Defines a slice of files to which the log will be writen.
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

This will result in the following output:

==== New Logger created on Fri, 01 Aug 2025 21:25:54 CEST at 21-25-54 ====

[21-25-54] [ERROR] This message will be shown

# Logging levels

One of the ways to customize a golog logger is changing which levels of logging it shows
and which it doesn't. Golog supports these 6 logging levels which are represented by
a uint8 number which, when written in binary, has only one of its bits occupied by 1
and the others are zeros. They can, therefore, be combined using the bitwise or (|)
operator.

 1. FATAL = 0b_00000001
 2. ERROR = 0b_00000010
 3. WARN  = 0b_00000100
 4. INFO  = 0b_00001000
 5. DEBUG = 0b_00010000
 6. TRACE = 0b_00100000

For more info about when should each of these levels be used refer to [this article]

# Customising time formats

Golog allows you to customise the timestamp format for your log messages. You can use the NewClock
function to specify one or more layouts, either using Go's time package constants or your own custom format strings.

Example:

	clock := golog.NewClock("2006-01-02 15:04:05", time.RFC3339)
	logger := golog.NewLogger([]*os.File{os.Stdout}, clock, golog.INFO|golog.ERROR)

This will format timestamps according to the layouts provided. You can use any valid Go time layout string.

# Logging to different files

Golog supports logging to multiple files or destinations. When creating a logger, you can pass a slice
of [os.File] objects to specify where log messages should be written.

Example:

	// Creates aand opens a new writalble file
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// Creates a logger, which logs to os.Stdout and the created file
	files := []*os.File{os.Stdout, logFile}
	logger := golog.NewLogger(files, golog.NewClock(time.StampMilli), golog.INFO|golog.ERROR)

	// Closes the files
	defer logger.Close()

This will write log messages to both standard output and the specified log file. You can add as many
files as needed to the slice. Make sure to use the Close() method afterwards.

[this article]: https://sematext.com/blog/logging-levels/
[its documentation]: https://pkg.go.dev/os#hdr-Concurrency
*/
package golog
