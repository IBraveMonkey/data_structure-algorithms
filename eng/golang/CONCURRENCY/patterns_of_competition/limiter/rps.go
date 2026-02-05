/*
 * The RPS Limiter (Rate Per Second) pattern limits the rate of operations.
 * Unlike maxGoroutine, which limits the number of concurrent goroutines,
 * RPS limiter limits the number of operations PER SECOND regardless of their duration.
 * Usage example: adhering to API limits (e.g., no more than 100 requests per second).
 */
package limiter

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Request3 struct {
	Payload string // Request data
}

type Client3 interface {
	SendRequest(ctx context.Context, request Request3) error
	WithLimiter(ctx context.Context, requests []Request3)
}

type client3 struct{}

// SendRequest executes HTTP request (simulation)
func (c client3) SendRequest(ctx context.Context, request Request3) error {
	time.Sleep(10 * time.Millisecond) // Simulate fast HTTP request
	fmt.Println("sending request", request.Payload)
	return nil
}

// WithLimiterRPS limits the request execution rate
func (c client3) WithLimiterRPS(ctx context.Context, requests []Request3) {
	const rps = 10 // Maximum 10 requests per second

	// Create request that sends signal 10 times per second
	ticker := time.NewTicker(time.Second / rps)
	defer ticker.Stop() // Stop ticker on exit

	// Process each request
	for _, req := range requests {
		<-ticker.C              // Wait for permission from ticker (1/10 second)
		c.SendRequest(ctx, req) // Execute request
	}
}

func rpsFn() {
	ctx := context.Background()
	c := client3{}

	// Create 100 requests
	requests := make([]Request3, 100)
	for i := 0; i < 100; i++ {
		requests[i] = Request3{Payload: strconv.Itoa(i)}
	}

	start := time.Now()

	// Send requests with 10 RPS limit
	// Should take ~10 seconds (100 requests / 10 RPS)
	c.WithLimiterRPS(ctx, requests)

	fmt.Printf("Processing took: %v\n", time.Since(start))
}
