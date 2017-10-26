package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 5.1
	d := true

	// %T is format verb which is Go-syntax representation of the typoe of the value
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}
