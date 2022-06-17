package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter year: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	// convert string to int
	year, error := strconv.Atoi(text)

	if error != nil {
		fmt.Println("Can't convert to int")
	}

	switch year {
	case 1997:
		fmt.Println("Year 1997")
	case 1998:
		fmt.Println("Year 1998")
	case 1999:
		fmt.Println("Year 1999")
	case 2000:
		fmt.Println("Year 2000")
	case 2001:
		fmt.Println("Year 2001")
	case 2002:
		fmt.Println("Year 2002")
	case 2003:
		fmt.Println("Year 2003")
	default:
		fmt.Println("This year is not exist!")
	}
}
