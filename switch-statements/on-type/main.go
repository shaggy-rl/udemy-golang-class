package main

import "fmt"

// switch on types
// -- normally we switch on value of variable
// -- go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}

// SwitchOnType takes in a unknow value and returns what type it is
// haven't done much with interfaces, allows us to pass in an unknow type value
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Shaggy")
	var t = contact{"Good to see you,", "Shaggy"}
	SwitchOnType(t)
}
