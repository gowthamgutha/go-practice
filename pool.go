package main

// create a pool of go routines
import (
	"fmt"
	"sync"
	"time"
)

func pool(max int) func(func()) {
	guard := make(chan struct{}, max)
	return func(function func()) {
		guard <- struct{}{}
		go func() {
			function()
			<-guard
		}()
	}
}

func main() {
	work := pool(3)

	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		work(func() {
			fmt.Println("Hello", i)
			time.Sleep(1 * time.Second)
			wg.Done()
		})
	}

	wg.Wait()
}
