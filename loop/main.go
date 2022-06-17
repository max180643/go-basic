package main

import "fmt"

func main() {
	// basic loop (print 0 - 9)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// while loop (print 0 - 9)
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	// while true (print 0 - 99)
	y := 0
	for {
		fmt.Println(y)
		y++
		if y > 99 {
			break
		}
	}
	// for loop with range
	// for index, name := range names {
	// 	fmt.Println(index, name)
	// }
}
