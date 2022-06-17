package main

import "fmt"

func main() {
	// constant
	const pi = "3.14"

	// declare variable
	var num int
	num = 16

	var message string = "Text String"

	// type inferred
	x := 2

	// variable type
	// string type
	var1 := "hello world"

	// integer
	var2 := 10

	// float
	var3 := 1.55

	// boolean
	var4 := true

	// shorthand string array declaration
	var5 := []string{"foo", "bar", "baz"}

	// map is reference datatype
	var6 := map[int]string{100: "Ana", 101: "Lisa", 102: "Rob"}

	// complex64 and complex128
	// is basic datatype
	var7 := complex(9, 15)

	fmt.Println(num)
	fmt.Println(message)
	fmt.Println(x)

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)
	fmt.Println(var7)

	// default value
	// int, uint, float = 0
	// complex = (0+0i)
	// string = ""
	// bool = false
	// function = nil
}
