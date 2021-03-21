package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals
		fmt.Println(sig)
		done <- struct{}{}
	}()

	fmt.Println("awaiting signal")
	<-done
}

func test() func(args ...interface{}) {
	return func(args ...interface{}) {
		go func(args ...interface{}) {}(args)
	}
}
