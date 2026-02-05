package main

import (
	"fmt"
)

/*
 * The Tee-channel pattern is used to split one incoming channel
 * into two (or more) outgoing channels. It is similar to a tee in plumbing:
 * data from one source is duplicated into multiple independent streams.
 * Important: reading from outgoing channels must happen in parallel,
 * otherwise blocking one channel will block the entire stream.
 */

// tee splits one incoming channel into two outgoing ones
func tee(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	// Create two output channels
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	// Start goroutine to duplicate data
	go func() {
		defer close(out1) // Close first channel on exit
		defer close(out2) // Close second channel on exit

		// Read values from input channel
		for val := range in {
			// Create local copies of channels for safe iteration
			var out1, out2 = out1, out2

			// Send value to both channels (need to send twice)
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					// Completion signal received, exit
					return
				case out1 <- val:
					// Successfully sent to first channel
					out1 = nil // Nullify to avoid resending in this iteration
				case out2 <- val:
					// Successfully sent to second channel
					out2 = nil // Nullify to avoid resending in this iteration
				}
			}
		}
	}()

	// Return both channels as read-only
	return out1, out2
}

func main() {
	// Create channel for completion signal
	done := make(chan interface{})
	defer close(done)

	// Create input channel
	in := make(chan interface{})

	// Goroutine to generate data into input channel
	go func() {
		defer close(in) // Close channel after sending all data
		for i := 1; i <= 5; i++ {
			in <- i // Send numbers from 1 to 5
		}
	}()

	// Split input channel into two output channels
	out1, out2 := tee(done, in)

	// Read from two channels simultaneously (important!)
	for val1 := range out1 {
		fmt.Printf("Stream 1 received: %v\n", val1)
		fmt.Printf("Stream 2 received: %v\n", <-out2) // Read from second channel
	}
}
