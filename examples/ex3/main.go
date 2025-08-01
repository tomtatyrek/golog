package main

import (
	"time"

	"github.com/tomtatyrek/golog"
)

func main() {

	start := time.Now()

	// Creating a logger with default configuration
	var l1 *golog.Logger = golog.NewDefaultLogger()
	defer l1.Close() // <-- THIS IS IMPORTANT

	l1.Debug("Hello World")
	time.Sleep(1 * time.Second)
	l1.Infof("Are you still here? It's been %v milliseconds since we've started", time.Since(start).Milliseconds())

	start2 := time.Now()

	thousandFactorial := factorial(20)

	l1.Warnf("Wow I'm starting to get bored. I've just counted twenty factorial and it only took me %v nanoseconds", time.Since(start2).Nanoseconds())

	time.Sleep(3 * time.Second)

	l1.Infof("It's %v btw...", thousandFactorial)

	time.Sleep(3 * time.Second)

	l1.Fatal("I'm critically bored and have to shut down D:")

	time.Sleep(2 * time.Second)

}

func factorial(n uint64) uint64 {
	if n < 2 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
