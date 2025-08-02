package main

import (
	"sync"
	"time"

	"github.com/tomtatyrek/golog"
)

var logger *golog.Logger
var wg sync.WaitGroup

func main() {
	logger = golog.NewDefaultLogger()
	defer logger.Close()
	wg.Add(3)
	go routine1()
	go routine2()
	go routine3()
	wg.Wait()
}

func routine1() {
	for range 5 {
		time.Sleep(500 * time.Millisecond)
		logger.Info("This is routine 1 calling")
	}
	wg.Done()
}

func routine2() {
	for range 5 {
		time.Sleep(500 * time.Millisecond)
		logger.Info("This is routine 2 calling")
	}
	wg.Done()
}

func routine3() {
	for range 5 {
		time.Sleep(500 * time.Millisecond)
		logger.Info("This is routine 3 calling")
	}
	wg.Done()
}
