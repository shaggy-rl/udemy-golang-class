package main

import "fmt"

// x is available to everything within the package
var x = 42

// anythin between {} is block level scope
func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
