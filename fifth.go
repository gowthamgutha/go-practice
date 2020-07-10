// routines
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go loop(5)
	time.Sleep(500) // bad way - sleep for some time to ensure that the goroutine runs

	wg := sync.WaitGroup{}
	wg.Add(1)
	go loopWg(5, &wg)
	wg.Wait()
}

func loop(x int) {
	for i := 0; i < x; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func loopWg(x int, wg *sync.WaitGroup) {
	fmt.Println("Executing loop with wait group")
	loop(5)

	wg.Done()
}
