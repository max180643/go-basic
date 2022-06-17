package main

import "fmt"

func main() {
	name := "max180643"
	age := 22

	// \n for new line
	// %v for basic value
	// %t for boolean
	// %T for type of a value
	// %d for standard, base-10 formatting
	// %f for basic decimal

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %d\n", age)

	// round float to 2 decimal places
	// fmt.Sprintf method use for display the value as a string
	convert := fmt.Sprintf("%.2f", 12.3456)
	fmt.Println(convert)
}
