package main

import "fmt"

func main() {
	a := 42

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
}

// & is operator that gets the variables memory address
