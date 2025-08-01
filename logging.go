package golog

import (
	"fmt"
)

// Log logs a string message with the specified logging level
func (l Logger) Log(level LogLevel, message string) {
	if level.isAllowed(l.allowedLogLevels) {
		l.log(fmt.Sprintf(
			"[%s] [%s] %s\n",
			l.clock.nowTime(),
			logLevelNames[level],
			message,
		))
	}
}

// Log logs a formatted message with the specified logging level
func (l Logger) Logf(level LogLevel, format string, a ...any) {
	if level.isAllowed(l.allowedLogLevels) {
		l.log(fmt.Sprintf(
			"[%s] [%s] %s\n",
			l.clock.nowTime(),
			logLevelNames[level],
			fmt.Sprintf(format, a...),
		))
	}
}

// Fatal logs a string message as a fatal error
func (l Logger) Fatal(message string) {
	l.Log(FATAL, message)
}

// Fatalf logs a formatted message as a fatal error
func (l Logger) Fatalf(format string, a ...any) {
	l.Log(FATAL, fmt.Sprintf(format, a...))
}

// Error logs a string message as an error
func (l Logger) Error(message string) {
	l.Log(ERROR, message)
}

// Error logs a formatted message as an error
func (l Logger) Errorf(format string, a ...any) {
	l.Log(ERROR, fmt.Sprintf(format, a...))
}

// Warn logs a string message as a warning
func (l Logger) Warn(message string) {
	l.Log(WARN, message)
}

// Warn logs a formatted message as a warning
func (l Logger) Warnf(format string, a ...any) {
	l.Log(WARN, fmt.Sprintf(format, a...))
}

// Info logs a string message as an information
func (l Logger) Info(message string) {
	l.Log(INFO, message)
}

// Info logs a formatted message as an information
func (l Logger) Infof(format string, a ...any) {
	l.Log(INFO, fmt.Sprintf(format, a...))
}

// Debug logs a string message as a debugging message
func (l Logger) Debug(message string) {
	l.Log(DEBUG, message)
}

// Debug logs a formatted message as a debugging message
func (l Logger) Debugf(format string, a ...any) {
	l.Log(DEBUG, fmt.Sprintf(format, a...))
}

// Trace logs a string message as a tracing message
func (l Logger) Trace(message string) {
	l.Log(TRACE, message)
}

// Trace logs a formatted message as a tracing message
func (l Logger) Tracef(format string, a ...any) {
	l.Log(TRACE, fmt.Sprintf(format, a...))
}
