// interface
package main

import (
	"fmt"
)

type Animal interface {
	Move() int // move using legs
	Talk() string
}

type Human struct {
	legs int
}
type Dog struct {
	legs int
}

func (d Dog) Move() int {
	return d.legs
}

func (h Human) Move() int {
	return h.legs
}

func (h Human) Talk() string {
	return "hello"
}

func main() {
	var animal Animal = Human{legs: 2}
	fmt.Println(animal.Move())
	fmt.Println(animal.Talk())

	/*
		Dog 'here' is not an Animal, because an Animal we 'declared' can Move(), Talk()
		But only Move() method is written
	*/
	fmt.Println(Dog{legs: 4}.Move())
}
