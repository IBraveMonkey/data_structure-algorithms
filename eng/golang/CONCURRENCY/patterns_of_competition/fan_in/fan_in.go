package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * The Fan-In pattern is used to combine multiple channels into a single channel.
 * It allows collecting results from multiple sources in one place.
 * Usage example: collecting results from several HTTP requests into one channel for further processing.
 */

// MergeChannels combines multiple input channels into one output channel
func MergeChannels(channels ...<-chan int) <-chan int {
	// Create resulting channel for combined data
	res := make(chan int)
	// WaitGroup to track completion of all reader goroutines
	wg := sync.WaitGroup{}

	// Add to counter the number of channels (one reader per channel)
	wg.Add(len(channels))

	// Start a separate goroutine for each input channel
	for _, channel := range channels {
		go func() {
			defer wg.Done()              // Decrement counter after reading is complete
			for value := range channel { // Read all values from channel
				res <- value // Redirect to common resulting channel
			}
		}()
	}

	// Separate goroutine to close the resulting channel
	go func() {
		wg.Wait()  // Wait until all channels are read
		close(res) // Close resulting channel
	}()

	// Return read-only channel
	return res
}

// Commented out alternative version of the same function
// func MergeChannels(channels ...<-chan int) <-chan int {
// 	wg := sync.WaitGroup{}
// 	wg.Add(len(channels))
//
// 	result := make(chan int)
// 	for _, channel := range channels {
// 		go func() {
// 			defer wg.Done()
// 			for value := range channel {
// 				result <- value
// 			}
// 		}()
// 	}
//
// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()
//
// 	return result
// }

func main() {
	// Create three channels for demonstration
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Start sender goroutine for data into three channels
	go func() {
		// defer guarantees closing of all channels on completion
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		// Generate 100 numbers, distributing them across three channels
		for i := 0; i < 100; i += 3 {
			ch1 <- i                           // Channel 1: 0, 3, 6, 9, ...
			ch2 <- i + 1                       // Channel 2: 1, 4, 7, 10, ...
			ch3 <- i + 2                       // Channel 3: 2, 5, 8, 11, ...
			time.Sleep(100 * time.Millisecond) // Delay for visibility
		}
	}()

	// Merge three channels into one and read from it
	for value := range MergeChannels(ch1, ch2, ch3) {
		fmt.Println(value) // Print values as they arrive
	}
}
