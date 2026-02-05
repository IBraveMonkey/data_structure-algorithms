package main

import (
	"fmt"
	"time"
)

/*
 * The Or-channel pattern is used to combine multiple channels into one,
 * which closes as soon as ANY of the input channels closes.
 * This is useful when you have multiple asynchronous tasks, and you only need
 * completion of any of them (e.g., multiple timeouts or data sources).
 */

// or recursively combines channels, returning a channel that closes when any input channel closes
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Base cases for recursion
	switch len(channels) {
	case 0:
		// If no channels, return nil
		return nil
	case 1:
		// If one channel, return it as is
		return channels[0]
	}

	// Create resulting channel
	orDone := make(chan interface{})

	// Start goroutine to monitor channels
	go func() {
		defer close(orDone) // Close resulting channel on exit

		switch len(channels) {
		case 2:
			// Optimization for two channels: wait for any to close
			select {
			case <-channels[0]: // First channel closed
			case <-channels[1]: // Second channel closed
			}
		default:
			// For 3+ channels: use recursion
			select {
			case <-channels[0]: // First channel closed
			case <-channels[1]: // Second channel closed
			case <-channels[2]: // Third channel closed
			case <-or(append(channels[3:], orDone)...): // Recursively process the rest
				// Combine remaining channels with orDone to interrupt recursion
			}
		}
	}()

	// Return read-only channel
	return orDone
}

func main() {
	// sig creates a channel that will close after specified time
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)    // Close channel after waiting
			time.Sleep(after) // Wait specified time
		}()
		return c
	}

	start := time.Now()

	// Wait for any of these channels to close
	// The fastest channel closes in 1 second
	<-or(
		sig(2*time.Hour),    // Closes in 2 hours
		sig(5*time.Minute),  // Closes in 5 minutes
		sig(1*time.Second),  // Closes in 1 second (fastest!)
		sig(1*time.Hour),    // Closes in 1 hour
		sig(10*time.Second), // Closes in 10 seconds
	)

	// Print execution time (should be ~1 second)
	fmt.Printf("Completed in %v\n", time.Since(start))
}
