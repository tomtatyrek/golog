package main

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
