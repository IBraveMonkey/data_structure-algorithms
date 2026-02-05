/*
 * The Limiter pattern is used to limit the number of concurrently executing operations.
 * This specific example limits the number of active connections (or goroutines) using a worker pool.
 * Usage example: limiting the number of concurrent HTTP requests to an API to avoid exceeding limits.
 */
package limiter

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Request1 represents a request to send
type Request1 struct {
	Payload string // Request data
}

// Client1 interface for working with requests
type Client1 interface {
	SendRequest(ctx context.Context, request Request1) error
	WithLimiter(ctx context.Context, requests []Request1)
}

type client1 struct{}

// SendRequest sends a single request (simulation)
func (c client1) SendRequest(ctx context.Context, request Request1) error {
	time.Sleep(100 * time.Millisecond) // Simulate network request
	fmt.Println("sending request", request.Payload)
	return nil
}

// Limiting the number of concurrently working connects (workers)
var maxConnects = 10

// WithLimiterWorkerPool limits the number of connections via a worker pool
func (c client1) WithLimiterWorkerPool(ctx context.Context, ch chan Request1) {
	wg := sync.WaitGroup{}

	// Start a fixed number of workers
	wg.Add(maxConnects) // Add number of workers to counter

	for range maxConnects {
		go func() {
			defer wg.Done() // Decrement counter on worker completion

			// Worker processes all requests from channel
			for req := range ch {
				c.SendRequest(ctx, req) // Process request
			}
		}()
	}

	wg.Wait() // Wait for all workers to complete
}

// maxConnectsFn demonstrates usage of the pattern
func maxConnectsFn() {
	ctx := context.Background() // Create context
	c := client1{}

	// Create 1000 requests
	requests := make([]Request1, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request1{Payload: strconv.Itoa(i)}
	}

	// Process requests with a limit of 10 workers
	c.WithLimiterWorkerPool(ctx, generate(requests))
}

// generate creates a channel and sends all requests to it
func generate(reqs []Request1) chan Request1 {
	// Create unbuffered channel
	ch := make(chan Request1)

	// Goroutine to send requests to channel
	go func() {
		for _, v := range reqs {
			ch <- v // Send each request
		}
		close(ch) // Close channel after sending all requests
	}()

	return ch
}
