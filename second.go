// second
// different ways of declaring variables
package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var y, z int = 20, 30
	a := 1
	// x := "test" throws error since x is of rtype int previously
	b := "b"
	bool := true
	float := 26.14
	fmt.Println(a, b, bool, float, x, y, z)
}
