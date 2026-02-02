package main

import (
	"fmt"
	"sync"
	"time"
)

/*
MUTEX IN GO:

- A Mutex ensures that only one goroutine can access a piece of code at a time.
- Protects shared resources like variables, maps, slices.
- Methods:
    - Lock()   -> acquire the lock
    - Unlock() -> release the lock
*/

func main() {
	// Shared counter variable
	counter := 0

	// Mutex to protect counter
	var mutex sync.Mutex

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Number of goroutines
	numGoroutines := 5

	fmt.Println("Starting goroutines to increment counter safely using mutex...")

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for j := 0; j < 3; j++ {
				// Lock the mutex before accessing shared variable
				mutex.Lock()
				counter++
				fmt.Printf("Goroutine %d incremented counter to %d\n", id, counter)
				// Unlock after modifying
				mutex.Unlock()

				// simulate some work
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}

