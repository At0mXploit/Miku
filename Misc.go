package main

import (
	"fmt"
	"sync"
	"time"
)

/*
GO-SPECIFIC CONCEPTS / WORDS:

This file demonstrates Go features that are unique or uncommon in other languages:
- Goroutines
- Channels
- Select
- Defer
- Iota
- Blank identifier _
- Mutex / RWMutex
- Panic / Recover
- Type switches
- Struct embedding
- Constants (typed/untyped)
*/

// ----------------------
// 1. Goroutines
// ----------------------
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

// ----------------------
// 2. Channel + Goroutine communication
// ----------------------
func worker(id int, ch chan string) {
	time.Sleep(time.Second) // simulate work
	ch <- fmt.Sprintf("Worker %d done", id)
}

// ----------------------
// 3. Panic and Recover
// ----------------------
func doPanic() {
	defer func() {
		if r := recover(); r != nil { // recover from panic
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Something went wrong!") // triggers panic
}

// ----------------------
// 4. Struct embedding
// ----------------------
type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
	Animal // embedding Animal
	Breed  string
}

// ----------------------
// 5. Type switch example
// ----------------------
func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}

// ----------------------
// 6. Constants + iota
// ----------------------
const (
	Red = iota
	Green
	Blue
)

type Status int

const (
	Pending Status = iota
	InProgress
	Completed
)

func main() {
	fmt.Println("=== 1. Goroutines ===")
	go sayHello("Alice") // runs concurrently
	go sayHello("Bob")
	time.Sleep(time.Millisecond * 100) // small sleep to allow goroutines to print

	fmt.Println("\n=== 2. Channels ===")
	ch := make(chan string)
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}
	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch) // receive messages from workers
	}

	fmt.Println("\n=== 3. Select Statement ===")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 500)
		ch1 <- "Message from ch1"
	}()
	go func() {
		time.Sleep(time.Millisecond * 200)
		ch2 <- "Message from ch2"
	}()

	select { // waits for whichever channel is ready first
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No messages ready")
	}

	fmt.Println("\n=== 4. Defer ===")
	fmt.Println("Start")
	defer fmt.Println("Deferred: End of main") // executes at function exit
	fmt.Println("Doing work...")

	fmt.Println("\n=== 5. Panic and Recover ===")
	doPanic() // safe panic handling

	fmt.Println("\n=== 6. Mutex ===")
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()        // protect shared resource
			counter++
			fmt.Println("Goroutine", id, "incremented counter to", counter)
			mu.Unlock()
		}(i + 1)
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter)

	fmt.Println("\n=== 7. Blank Identifier _ ===")
	_, err := fmt.Println("This value will be ignored")
	fmt.Println("Error ignored:", err)

	fmt.Println("\n=== 8. Type Switch ===")
	printType(100)
	printType("Hello")
	printType(true)
	printType(3.14)

	fmt.Println("\n=== 9. Struct Embedding ===")
	d := Dog{Animal{Name: "Rex"}, "Labrador"}
	d.Speak() // inherited from embedded Animal
	fmt.Println("Breed:", d.Breed)

	fmt.Println("\n=== 10. Constants + iota ===")
	fmt.Println("Red:", Red, "Green:", Green, "Blue:", Blue)
	status := InProgress
	fmt.Println("Current Status:", status)
	switch status {
	case Pending:
		fmt.Println("Task Pending")
	case InProgress:
		fmt.Println("Task In Progress")
	case Completed:
		fmt.Println("Task Completed")
	}
}

