package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}

/*
x is set to 5 in main. x memory address is passed by value over to zero. then
zero derefences the address pased to it and sets the value to 0. when main
prints out x the value of 0 that is stored now in the memory address is printed
*/
