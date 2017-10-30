package main

import "fmt"

// can specify type
const p string = "death & taxes"

func main() {
	// can also do it without specifing type
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

// a CONSTANT is a simple unchaning value
