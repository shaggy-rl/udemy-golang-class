package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Wassup Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushatn")
	}
}
