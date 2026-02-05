package main

import "fmt"

/*
 * The Pipeline pattern is used to create a chain of data handlers, where each stage processes data and passes it to the next stage.
 * It allows breaking down complex processing into several simple steps.
 * Usage example: image processing, where each image passes through several stages of filtering and transformation.
 */

// gen — first stage of the pipeline: generates numbers from the passed slice
func gen(numbers ...int) <-chan int {
	// Create channel for outputting numbers
	out := make(chan int)

	// Start goroutine to generate numbers
	go func() {
		defer close(out) // Close channel after sending all numbers
		for _, number := range numbers {
			out <- number // Send each number to channel
		}
	}()

	// Return read-only channel
	return out
}

// mul — second stage of the pipeline: multiplies each number by itself (squaring)
func mul(in <-chan int) <-chan int {
	// Create output channel
	out := make(chan int)

	// Start goroutine to process incoming numbers
	go func() {
		defer close(out)         // Close output channel after processing all data
		for number := range in { // Read numbers from input channel
			out <- number * number // Square and send to output channel
		}
	}()

	// Return read-only output channel
	return out
}

func main() {
	// Create pipeline: gen() -> mul()
	ch := gen(1, 2, 3, 4, 5) // Generate numbers 1-5

	// Read results from the second stage of the pipeline
	for value := range mul(ch) { // mul receives numbers and returns their squares
		fmt.Println(value) // Will print: 1, 4, 9, 16, 25
	}
}
