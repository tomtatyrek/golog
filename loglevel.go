package golog

// A LogLevel is used both to specify the logging level of a message
// and to list logging levels a [Logger] is allowed to log
type LogLevel uint8

// Golog supports these 6 logging levels which are represented by
// a uint8 number which, when written in binary, has only one of its
// bits occupied by 1 and the others are zero. They can, therefore,
// be combined using the bitwise or (|) operator.
//
//  1. FATAL = 0b_00000001
//  2. ERROR = 0b_00000010
//  3. WARN  = 0b_00000100
//  4. INFO  = 0b_00001000
//  5. DEBUG = 0b_00010000
//  6. TRACE = 0b_00100000
//
// For more info about when should each of these levels be used refer
// to [this article]
//
// [this article]: https://sematext.com/blog/logging-levels/.
const (
	FATAL LogLevel = 0b_00000001
	ERROR LogLevel = 0b_00000010
	WARN  LogLevel = 0b_00000100
	INFO  LogLevel = 0b_00001000
	DEBUG LogLevel = 0b_00010000
	TRACE LogLevel = 0b_00100000
	// RES1  LogLevel = 0b_01000000
	// RES2  LogLevel = 0b_10000000
)

var logLevelNames = map[LogLevel]string{
	FATAL: "FATAL",
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
	TRACE: "DEBUG",
	// RES1:  "RES1",
	// RES2:  "RES2",
}

const DefaultLogLevels LogLevel = FATAL | ERROR | WARN | INFO | DEBUG
const AllLogLevels LogLevel = FATAL | ERROR | WARN | INFO | DEBUG | TRACE

func (ll LogLevel) isAllowed(allowed LogLevel) bool {
	return (ll & (allowed ^ 0b_11111111)) == 0
}
