package main

import "fmt"

func main() {
	num := 7

	// if / else
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// switch
	switch num {
	case 1:
		fmt.Println("One")
	case 7:
		fmt.Println("Lucky seven!")
	default:
		fmt.Println("Other number")
	}
}

