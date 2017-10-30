package main

import "fmt"

// NEVER do the below example in real code; bad coding practice to shadow vars
func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max is now the result, not the function
}
