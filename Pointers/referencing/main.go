package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
}

/*
the above code makes b a pointer to the memory address where an int is stored
b is of type "int pointer"
*int -- the * is the part of the type -- b is of type *int
when you print out b it will be the address of a
it should be noted that Go will infere *int when declaring b so it should be
omitted to follow  Go best practices.
*/
