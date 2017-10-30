package main

import "fmt"

// func wrapper return type is a function that has a return type of an int

func wrapper() func() int {
	x := 0

	// fucn being returned
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

// will print 1 2
