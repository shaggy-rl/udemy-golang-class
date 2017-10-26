package main

import "fmt"

func main() {
	// := is used to initialize i = 0
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x\n", i, i, i)
	}
}
