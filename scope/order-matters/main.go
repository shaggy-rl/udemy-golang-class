package main

import "fmt"

// this will not compile
// jjx is declared after it is called
// y will work though due to y being declared in the package level scope
func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}

var y = 42
