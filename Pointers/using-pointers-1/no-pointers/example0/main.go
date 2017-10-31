package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}

/*
zero takes in an int
we apss the VALUE of x into zero, hence x not changing
if we passed the memory addr of x in main it would be set to 0 and print 0
*/
