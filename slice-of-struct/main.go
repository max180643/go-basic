package main

import "fmt"

type student struct {
	name string
	year int
}

func main() {
	students := []student{
		student{
			name: "Steve Ronin",
			year: 3,
		},
		student{
			name: "Robert Cane",
			year: 4,
		},
	}

	// append
	students = append(students, student{
		name: "Feb Dow",
		year: 5})

	for index, stud := range students {
		position := index + 1
		fmt.Printf("%d. Name: %s\t", position, stud.name)
		fmt.Printf("Year: %d\n", stud.year)
	}
}
