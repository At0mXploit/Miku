package main

import (
	"fmt"
	"time"
)

/*
CHANNELS IN GO:

- Channels are a way for goroutines (lightweight threads) to communicate safely.
- Think of a channel as a **pipe** through which data flows from one goroutine to another.
- Channels can be **unbuffered** (send blocks until receive) or **buffered** (send only blocks when full).
- Syntax:
    ch := make(chan Type)           // unbuffered channel
    ch := make(chan Type, capacity) // buffered channel

- Send:    ch <- value
- Receive: value := <-ch
- Close:   close(ch)
*/

func worker(id int, ch chan string) {
	/*
		This function simulates a task:
		- It sleeps for 1 second (simulate work)
		- Sends a message into the channel when done
	*/
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("Worker %d finished", id)
}

func main() {

	// =====================
	// 1. What is a channel
	// =====================
	fmt.Println("=== 1. Basic Channel Example ===")
	// Create an unbuffered channel of strings
	ch := make(chan string)

	// Start a goroutine that does some work
	go worker(1, ch)

	// Receive from the channel. This will **block** until the goroutine sends a value
	msg := <-ch
	fmt.Println("Received:", msg)
	// Explanation: Channels are blocking by default.
	// The main function waits here until the goroutine sends a value.

	// =====================
	// 2. Multiple goroutines with one channel
	// =====================
	fmt.Println("\n=== 2. Multiple Goroutines Example ===")

	done := make(chan string) // single channel to collect multiple results

	// Launch 3 workers concurrently
	for i := 1; i <= 3; i++ {
		go worker(i, done)
	}

	// Collect results from all workers
	// Note: the order may vary because goroutines run concurrently
	for i := 1; i <= 3; i++ {
		result := <-done
		fmt.Println("Received:", result)
	}

	// =====================
	// 3. Buffered channels
	// =====================
	fmt.Println("\n=== 3. Buffered Channel Example ===")

	// Buffered channel of capacity 2
	buffered := make(chan int, 2)

	// Send two values into the channel; does NOT block because buffer is not full
	buffered <- 10
	buffered <- 20
	fmt.Println("Sent two values into buffered channel")

	// Reading from buffered channel
	fmt.Println("Received:", <-buffered)
	fmt.Println("Received:", <-buffered)

	// =====================
	// 4. Closing channels
	// =====================
	fmt.Println("\n=== 4. Closing Channels ===")

	ch2 := make(chan int)

	// Goroutine sends multiple values then closes the channel
	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- i
		}
		close(ch2) // signal no more values will be sent
	}()

	// Range over channel receives until channel is closed
	for val := range ch2 {
		fmt.Println("Received from closed channel:", val)
	}

	// =====================
	// 5. Select statement
	// =====================
	fmt.Println("\n=== 5. Select Statement ===")

	chA := make(chan string)
	chB := make(chan string)

	// Send values from goroutines after different delays
	go func() {
		time.Sleep(1 * time.Second)
		chA <- "Message from chA"
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		chB <- "Message from chB"
	}()

	// Select waits on multiple channels and picks the one that’s ready first
	select {
	case msg := <-chA:
		fmt.Println("Received:", msg)
	case msg := <-chB:
		fmt.Println("Received:", msg)
	}

	// =====================
	// Summary Explanation
	// =====================
	fmt.Println("\n=== Channel Summary ===")
	fmt.Println("1. Channels are used to communicate between goroutines.")
	fmt.Println("2. Unbuffered channels block on send and receive until both sides are ready.")
	fmt.Println("3. Buffered channels allow sending up to 'capacity' items without blocking.")
	fmt.Println("4. Closing a channel signals no more values will be sent; reading after close returns zero value.")
	fmt.Println("5. Select allows waiting on multiple channels simultaneously.")
}

