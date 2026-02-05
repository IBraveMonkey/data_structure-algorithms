/*
 * The Limiter pattern is used to limit the number of concurrently executing operations.
 * This specific example limits the number of concurrently running goroutines using a semaphore.
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

type Request2 struct {
	Payload string
}

type Client2 interface {
	SendRequest(ctx context.Context, request Request2) error
	WithLimiter(ctx context.Context, requests []Request2)
}

type client2 struct{}

func (c client2) SendRequest2(ctx context.Context, request Request2) error {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("sending request", request.Payload)
	return nil
}

var maxGoroutines = 100

// limit the number of concurrently working goroutines (100)
func (c client2) WithLimiterSemaphore(ctx context.Context, reqs []Request2) {
	tokens := make(chan struct{}, maxGoroutines)
	wg := sync.WaitGroup{}

	wg.Add(len(reqs))
	for _, req := range reqs {
		tokens <- struct{}{}
		go func() {
			defer func() {
				wg.Done()
				<-tokens
			}()
			c.SendRequest2(ctx, req)
		}()
	}

	wg.Wait()
}

func maxGoroutineFn() {
	ctx := context.Background()        // Create background context for operations.
	c := client2{}                     // Initialize client for sending requests.
	requests := make([]Request2, 1000) // Create slice of 1000 requests.
	for i := 0; i < 1000; i++ {
		requests[i] = Request2{Payload: strconv.Itoa(i)}
	}
	c.WithLimiterSemaphore(ctx, requests)
}
