package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // memroy address of a

	var b *int = &a
	fmt.Println(b)  // memory address
	fmt.Println(*b) // dereferencing 43

	*b = 42        // b says, "The value at this adress, change it ot 42"
	fmt.Println(a) // 42
}

/*
this is very useful
we can pass a memory address instead of a bunch of values (pass a reference)
and then we can still change the value of wahtever is stored at that mem addr
this makes programs way more performant
this allows us not to have to pass around large amounts of data

everything is PASS BY VALUE in GO
when we pass a mem address, we are passing a value
*/
