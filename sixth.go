// for loop
package main

import (
	"fmt"
)

func main() {

	// assignment, condition, update
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// just condition
	i := 1
	for i <= 5 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	// only condition and update
	i = 1
	for ; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Empty
	i = 1
	for {
		if i <= 5 {
			fmt.Printf("%d ", i)
		} else {
			break
		}
		i++
	}

	fmt.Println()
	// inner loops
outer:
	for x := 0; x < 5; x++ {
		for j := 0; j < 5; j++ {
			if i+j == 10 {
				break outer
			}
			fmt.Printf("%d ", i+j)
		}
	}
	fmt.Println()
	array := []int{1, 2, 3, 4, 5}

	// just print the numbers
	for _, i := range array {
		fmt.Printf("%d ", i)
	}

	fmt.Println()
	// print the index also
	for idx, val := range array {
		fmt.Printf("a[%d]=%d ", idx, val)
	}
}
