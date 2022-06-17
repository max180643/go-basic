package main

import (
	"fmt"
	"function/tpackage"
)

func main() {
	sayHi("Max180643", 22)
	tpackage.SayHi()
}

func sayHi(name string, age int) {
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
