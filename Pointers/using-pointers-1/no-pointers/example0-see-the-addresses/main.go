package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x) // address in func zero
	fmt.Println(&x)        // addres sin func zero
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) //address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}

/*
same code in example0, but showing how the memroy addresses are different
hence why the value of x doesn't change in main
*/
