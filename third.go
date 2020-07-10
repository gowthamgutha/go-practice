// functions example
package main

import (
	"fmt"
)

func main() {
	print(10)
	printAll(10, 20, 30)
	fmt.Println(double(10))
	fmt.Println(doubleTriple(10))
	fmt.Println(swap(10, 20))
	x := 9
	increment(&x)
	fmt.Println(x)
	nextInt := sequence() // initialize i=0 and store the function in nextInt variable
	// since nextInt is of type func, use () to invoke it
	// since the variable is initialized only once, we will see nextInt() returning incremented values on every call
	fmt.Println(nextInt(), nextInt(), nextInt())

	// anonymous functions
	func(x int) {
		fmt.Println("In anonymous function", x)
	}(nextInt())

	anon := func(x int) {
		fmt.Println("In anonymous function assigned to variable", x)
	}
	anon(nextInt())

	anon1 := func(x func() int) {
		fmt.Println("In anonymous function that takes another function as arg", x())
	}
	anon1(nextInt)
}

func print(x int) {
	fmt.Println(x)
}

func printAll(x ...int) {
	for index, num := range x {
		fmt.Printf("%d => %d\n", index, num)
	}
}

func double(x int) int {
	return x * 2
}

func doubleTriple(x int) (int, int) {
	return x * 2, x * 3
}

func swap(x int, y int) (int, int) {
	return y, x
}

func increment(x *int) {
	*x++
}

// closure
func sequence() func() int {
	i := 0 // declare the variable here, initialized at the time of calling sequence()

	// this return value is stored in a variable (function can be stored in a variable)
	return func() int {
		i++
		return i // This returns the incremented value
	}
}
