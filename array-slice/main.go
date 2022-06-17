package main

import "fmt"

func main() {
	// array: fixed length
	// declare array without value
	var names [3]string
	names[0] = "Harry"
	names[1] = "Ron"
	names[2] = "Hermione"
	// declare array with value
	names1 := [3]string{"Harry", "Ron", "Hermione"}
	// slice: dynamic length
	var names2 []string
	names2 = append(names2, "Harry")
	names2 = append(names2, "Ron")
	names2 = append(names2, "Hermione")
	// short
	names3 := []string{"Harry", "Ron", "Hermione"}

	// for loop (simple)
	for index := 0; index < len(names1); index++ {
		fmt.Println(index, names1[index])
	}
	// for loop with range
	for index, name := range names3 {
		fmt.Println(index, name)
	}
}
