package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Method on struct
func (p Person) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
	p.Greet()

	// Struct pointer
	ptr := &p
	ptr.Age = 31
	ptr.Greet()
}

