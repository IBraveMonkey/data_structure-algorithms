package genarator

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
 * The Generator pattern is used to create functions that return a channel and generate values for that channel.
 * It allows creating data streams that can be used by other parts of the program.
 * Usage example: generating a sequence of numbers that can be used for processing in other goroutines.
 */

// randomTimeWork simulates work that takes a random amount of time (up to 100 seconds)
func randomTimeWork() {
	// Generate a random number from 0 to 99 and multiply by a second
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

// PredictableTimeWork demonstrates how to limit the execution time of a function
// How to make a function work only for some time - 3s
func PredictableTimeWork() {
	// Create a channel for work completion signal
	ch := make(chan struct{})

	// Start a goroutine that will do the work and close the channel
	go func() {
		randomTimeWork() // Do the work (can be long)
		close(ch)        // Close channel, signaling completion
	}()

	// Wait for either work completion or timeout
	select {
	case <-ch:
		// Work finished before timeout
		fmt.Println("Done")
	case <-time.After(3 * time.Second):
		// 3 seconds passed, interrupt waiting
		fmt.Println("Cancel with timeout")
	}
}

// writer is a generator that creates a channel and returns it
// Micropattern - Generator that creates a channel and returns it
func writer() <-chan int {
	// Create channel for passing integers
	ch := make(chan int)
	// Create WaitGroup to wait for all goroutines
	wg := &sync.WaitGroup{}

	// Add 2 goroutines to counter
	wg.Add(2)

	// First goroutine: sends numbers from 1 to 5
	// Should be non-blocking, must be done in a separate goroutine
	go func() {
		defer wg.Done()    // Decrement counter on function completion
		for i := range 5 { // Iterate from 0 to 4
			ch <- i + 1 // Send numbers 1, 2, 3, 4, 5 to channel
		}
	}()

	// Second goroutine: sends numbers from 11 to 15
	go func() {
		defer wg.Done()    // Decrement counter on function completion
		for i := range 5 { // Iterate from 0 to 4
			ch <- i + 11 // Send numbers 11, 12, 13, 14, 15 to channel
		}
	}()

	// Third goroutine: waits for completion of all goroutines and closes channel
	go func() {
		wg.Wait() // Blocks until counter becomes 0
		close(ch) // Close channel after all data is sent
	}()

	// Return read-only channel (safety)
	return ch
}

// Generator demonstrates usage of the Generator pattern
func Generator() {
	// Get channel from generator
	ch := writer()

	// Read all values from channel until it is closed
	for v := range ch {
		fmt.Println("v =", v) // Print each received value
	}
}
