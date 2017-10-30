package vis

import "fmt"

// PrintVar prints MyName andf yourName due to yourName being in the package scope
// Not visable in main.go
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
}
