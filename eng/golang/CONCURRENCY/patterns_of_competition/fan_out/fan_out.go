package fan_out

import (
	"fmt"
	"sync"
	"time"
)

/*
 * The Fan-Out pattern is used to distribute data from one channel to multiple other channels.
 * It allows processing data by multiple goroutines in parallel.
 * Usage example: processing tasks by multiple workers, where one task queue is distributed among multiple handlers.
 */

// worker processes tasks from the input channel
func worker(
	id int, // Worker ID for identification in logs
	input <-chan int, // Channel for reading tasks (read-only)
	wg *sync.WaitGroup, // WaitGroup for completion synchronization
) {
	defer wg.Done()          // Decrement counter on worker completion
	for num := range input { // Read tasks until channel is closed
		fmt.Printf("Worker %d processing %d\n", id, num)
		time.Sleep(500 * time.Millisecond) // Simulate processing
		fmt.Printf("Worker %d finished %d, result: %d\n", id, num, num*2)
	}
}

// FanOut demonstrates the Fan-Out pattern
func FanOut() {
	const numWorkers = 3                     // Number of workers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8} // Tasks to process

	// Create a separate channel for each worker
	inputs := make([]chan int, numWorkers)
	var wg sync.WaitGroup

	// Start all workers
	for i := 0; i < numWorkers; i++ {
		inputs[i] = make(chan int)     // Create channel for worker
		wg.Add(1)                      // Add to counter
		go worker(i+1, inputs[i], &wg) // Start worker (ID from 1)
	}

	// Goroutine to distribute tasks among workers
	go func() {
		for i, num := range numbers {
			// Distribute tasks round-robin
			inputs[i%numWorkers] <- num // Send task to worker channel
		}

		// Close all channels after distributing all tasks
		for _, in := range inputs {
			close(in) // Signal worker to finish
		}
	}()

	wg.Wait() // Wait for all workers to complete
	fmt.Println("All workers done")
}

// JoinChannels combines multiple channels into one (reverse operation to Fan-Out)
func JoinChannels(chs ...<-chan int) <-chan int {
	// Create output channel for merged data
	mergedCh := make(chan int)

	// Start goroutine to merge channels
	go func() {
		wg := &sync.WaitGroup{}

		// Add goroutine for each input channel
		wg.Add(len(chs))

		// Read from each channel in a separate goroutine
		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()      // Decrement counter on completion
				for id := range ch { // Read all values from channel
					mergedCh <- id // Redirect to common channel
				}
			}(ch, wg)

			wg.Wait()       // Wait for all goroutines to complete
			close(mergedCh) // Close output channel
		}
	}()

	// Return read-only channel
	return mergedCh
}
