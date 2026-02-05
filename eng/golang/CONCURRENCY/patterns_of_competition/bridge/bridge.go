package main

import (
	"fmt"
)

/*
 * The Bridge-channel pattern is used for "flattening" a stream of channels.
 * If you have a channel from which other channels come (chan <-chan interface{}),
 * the Bridge pattern allows reading values directly from a single resulting channel.
 */

// bridge joins a stream of channels into a single channel of values
func bridge(done <-chan interface{}, chanStream <-chan (<-chan interface{})) <-chan interface{} {
	// Create the resulting channel for values
	valStream := make(chan interface{})

	// Start a goroutine to process the stream of channels
	go func() {
		defer close(valStream) // Close the output channel upon completion

		for {
			var stream <-chan interface{} // Current channel being processed

			// Wait for the next channel from the stream or a completion signal
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					// Channel stream is closed, finish work
					return
				}
				stream = maybeStream // Received a new channel for reading
			case <-done:
				// Completion signal received
				return
			}

			// Read all values from the current channel
			for val := range orDone(done, stream) {
				select {
				case valStream <- val: // Send value to the resulting channel
				case <-done:
					// Interrupt if completion signal received
					return
				}
			}
		}
	}()

	// Return read-only channel
	return valStream
}

// orDone - helper function for correctly closing channels considering done
func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	// Create output channel
	valStream := make(chan interface{})

	// Start goroutine for reading considering done
	go func() {
		defer close(valStream) // Close channel on exit

		for {
			select {
			case <-done:
				// Completion signal received
				return
			case v, ok := <-c:
				if !ok {
					// Input channel closed
					return
				}
				// Try to send value to output channel
				select {
				case valStream <- v: // Successfully sent
				case <-done:
					// Interrupted by completion signal
					return
				}
			}
		}
	}()

	return valStream
}

func main() {
	// Create channel for completion signal
	done := make(chan interface{})
	defer close(done) // Close on main completion

	// genChanStream - generator that creates a stream of channels
	genChanStream := func() <-chan (<-chan interface{}) {
		// Channel of channels (stream of channels)
		chanStream := make(chan (<-chan interface{}))

		go func() {
			defer close(chanStream) // Close channel stream on completion

			// Create 5 channels and send them to the stream
			for i := 0; i < 5; i++ {
				stream := make(chan interface{}) // Create new channel

				// Goroutine to send value to channel
				go func(v int) {
					defer close(stream) // Close channel after sending
					stream <- v         // Send value
				}(i)

				chanStream <- stream // Send channel to channel stream
			}
		}()

		return chanStream
	}

	// Read from all channels as from one thanks to bridge
	for val := range bridge(done, genChanStream()) {
		fmt.Printf("Received value: %v\n", val) // Will print numbers 0, 1, 2, 3, 4
	}
}
