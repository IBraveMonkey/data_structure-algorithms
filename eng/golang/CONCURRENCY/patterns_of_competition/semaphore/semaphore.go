package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
 * The Semaphore pattern is used to limit the number of concurrently executing goroutines.
 * It allows controlling access to resources that have limited throughput.
 * Usage example: limiting the number of concurrent HTTP requests to an API to avoid exceeding limits.
 */

// downloadFile simulates file download with context for cancellation
func downloadFile(ctx context.Context, url string) {
	// Create channel for download completion signal
	ch := make(chan struct{})

	// Start goroutine to simulate download
	go func() {
		time.Sleep(1 * time.Second) // Simulate download (1 second)
		close(ch)                   // Signal completion
	}()

	// Wait for either download completion or context cancellation
	select {
	case <-ctx.Done():
		// Context was cancelled (timeout or cancel)
		fmt.Println(ctx.Err())
		return
	case <-ch:
		// Download successfully completed
		fmt.Printf("%s\n", url)
		return
	}
}

func main() {
	// Create context with timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Must call cancel to free resources

	const goroutineLimit = 3 // Maximum 3 goroutines simultaneously

	// List of files to download
	files := []string{"https://1.ru", "https://2.ru", "https://3.ru", "https://4.ru", "https://5.ru", "https://6.ru", "https://7.ru"}

	wg := sync.WaitGroup{}
	// Semaphore: buffered channel for goroutineLimit elements
	semaphore := make(chan struct{}, goroutineLimit)

	wg.Add(len(files)) // Add number of files to counter

	// Start goroutine for each file
	for _, file := range files {
		semaphore <- struct{}{} // Occupy slot in semaphore (blocks if no slots)

		go func() {
			defer func() {
				defer wg.Done() // Decrement counter on completion
				<-semaphore     // Release slot in semaphore
			}()

			downloadFile(ctx, file) // Execute download
		}()
	}

	wg.Wait() // Wait for all goroutines to complete
}
