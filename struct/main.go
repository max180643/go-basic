package main

import "fmt"

type student struct {
	name string
	year int
}

// function with struct
func (s student) greeting() {
	if s.name == "Steve Ronin" {
		fmt.Printf("%v: Bello!\n", s.name)
	} else {
		fmt.Printf("%v: Hello!\n", s.name)
	}
}

func main() {
	student1 := student{
		name: "Steve Ronin",
		year: 3,
	}

	student2 := student{
		name: "Robert Cane",
		year: 4,
	}

	student1.greeting()
	student2.greeting()
}
