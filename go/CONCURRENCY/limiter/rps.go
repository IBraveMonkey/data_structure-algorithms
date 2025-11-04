package limiter

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Request3 struct {
	Payload string
}

type Client3 interface {
	SendRequest(ctx context.Context, request Request3) error
	WithLimiter(ctx context.Context, requests []Request3)
}

type client3 struct{}

func (c client3) SendRequest3(ctx context.Context, request Request3) error {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("sending request", request.Payload)
	return nil
}

var rps = 100

// ограничение кол-ва rps одновременно работающих (10000)
func (c client3) WithLimiterRps(ctx context.Context, reqs []Request3) {
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	wg := sync.WaitGroup{}

	wg.Add(len(reqs))
	for _, req := range reqs {
		<-ticker.C

		go func() {
			defer wg.Done()
			c.SendRequest3(ctx, req)
		}()
	}

	wg.Wait()
}

func rpsLimiter() {
	ctx := context.Background()
	c := client3{}
	requests := make([]Request3, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request3{Payload: strconv.Itoa(i)}
	}
	c.WithLimiterRps(ctx, requests)
}
