package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
using an anonymous function which is a function without a name
increment is set = to the anonymous function

func expression
assigning a func to a variable
*/
