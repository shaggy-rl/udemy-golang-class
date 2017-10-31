package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // address of a in hex

	var b *int = &a
	fmt.Println(b)  // address of a in hex
	fmt.Println(*b) // 43
}

/*
b is an int pointer
b point to the memory address where an int is stored
to see the value in that memory address, add a * in front of the variable
this is know and dereferencing
the * is an operator in this case
*/
