package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * The Worker Pool pattern is used to distribute tasks among multiple handlers.
 * It allows efficiently using resources by distributing tasks among multiple goroutines.
 * Usage example: image processing, where each image is processed by one of the workers.
 */

// Task represents a task for processing
type Task struct {
	ID       int    // Task identifier
	Filename string // Filename to process
}

// Process processes one task (simulating long work)
func Process(task Task) string {
	time.Sleep(1 * time.Second) // Simulate processing (1 second)

	// Return processing result
	return fmt.Sprintf("FileID: %d done - %s \n", task.ID, task.Filename)
}

// Worker reads tasks from channel and processes them
func Worker(taskCh <-chan Task, resCh chan<- string) {
	for task := range taskCh { // Read tasks until channel is closed
		resCh <- Process(task) // Process and send result
	}
}

func main() {
	// Constants for worker pool configuration
	const (
		numWorkers = 3  // Number of workers
		numTasks   = 10 // Number of tasks
	)

	// Create buffered channels for tasks and results
	taskCh := make(chan Task, numTasks)  // Task channel
	resCh := make(chan string, numTasks) // Result channel

	// WaitGroup to track completion of all workers
	wg := sync.WaitGroup{}
	wg.Add(numWorkers) // Add number of workers

	// Start worker pool
	for range numWorkers {
		go func() {
			defer wg.Done()       // Decrement counter on completion
			Worker(taskCh, resCh) // Worker processes tasks
		}()
	}

	// Goroutine to send tasks to pool
	go func() {
		for i := range numTasks { // Generate tasks
			// Create and send task to channel
			taskCh <- Task{ID: i, Filename: fmt.Sprintf("file_%d.jpg", i)}
		}
		defer close(taskCh) // Close task channel after sending all
	}()

	// Goroutine to close result channel after all workers complete
	go func() {
		wg.Wait()    // Wait for all workers to complete
		close(resCh) // Close result channel
	}()

	// Read and print results as they arrive
	for res := range resCh {
		fmt.Println(res) // Print processing result
	}
}
