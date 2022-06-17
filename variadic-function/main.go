package main

import "fmt"

func main() {
	variadicExample("Red", "Blue", "Green")
}

func variadicExample(s ...string) {
	fmt.Printf("First:\t%s\n", s[0])
	fmt.Printf("Second:\t%s\n", s[1])
	fmt.Printf("Third:\t%s\n", s[2])
}
