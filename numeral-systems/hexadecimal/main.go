package main

import "fmt"

func main() {
	// %x is hexadecimal base 16
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)

	// %#x will add leading 0x
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)

	// %#X will add leading 0x and make letters capitalized
	fmt.Printf("%d - %b - %#X\n", 42, 42, 42)

	// \t escape code is a tab character
	fmt.Printf("%d \t %b \t %#X\n", 42, 42, 42)
}
