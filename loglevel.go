package golog

// LogLevel is used both to sidentify the logging level of a message
// and to specify messages of which levels a [Logger] is allowed to log
type LogLevel uint8

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
