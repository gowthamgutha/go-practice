// time example
package main

import (
	"fmt"
	"time"
)
func main() {
	// Print the time
	fmt.Println(time.Now())

	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)

	time := time.Now()
	fmt.Println(time.Hour(), time.Minute(), time.Second())

	hour, min, sec := time.Clock()
	fmt.Printf("%v hour, %v min, %v sec", hour, min, sec)
}