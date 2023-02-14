package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	now := time.Now()
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
	go helloWorld()

	select {
	case <-sigChannel:

		fmt.Println("Stopped by user after", time.Since(now), "second")
		os.Exit(0)
	}
}
func helloWorld() {
	fmt.Println("hello World")
	time.Sleep(10 * time.Second)
	fmt.Println("Goodbye world")
	os.Exit(0)
}
