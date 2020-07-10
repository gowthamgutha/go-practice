// maps
package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages["rama"] = 21
	if _, ok := ages["sita"]; !ok {
		fmt.Println("No age found for entry sita")
	}
	fmt.Println(ages)
}
