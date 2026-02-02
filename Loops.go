package main

import "fmt"

func main() {
	// Classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// for as while
	x := 0
	for x < 3 {
		fmt.Println("x =", x)
		x++
	}

	// for range over slice
	names := []string{"Alice", "Bob", "Eve"}
	for i, name := range names {
		fmt.Println(i, name)
	}

	// for range over map
	ages := map[string]int{"Alice": 25, "Bob": 30}
	for name, age := range ages {
		fmt.Println(name, age)
	}
}

