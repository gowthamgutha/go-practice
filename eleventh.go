// defer example
package main

import "fmt"

func main() {
	fmt.Println("First line")
	defer fmt.Println("Executed at last")         // last
	defer fmt.Println("Executed at last but one") // last but one
	defer fmt.Println("Executed at last but two") // last but two
	test()
	fmt.Println("Last line")
}

// this example is taken from https://golang.org/doc/effective_go.html
func trace(method string) string {
	fmt.Println("Entering", method)
	return method
}

func un(method string) {
	fmt.Println("Leaving", method)
}

func test() {
	// the arg to un() is evaluated first, so trace() is printed, but un() itself is executed at the enf of test()
	defer un(trace("Test"))
	fmt.Println("In test function")
}
