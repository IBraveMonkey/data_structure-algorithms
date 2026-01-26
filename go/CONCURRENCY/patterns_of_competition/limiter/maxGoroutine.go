/*
 * Паттерн Limiter (Ограничитель) используется для ограничения количества одновременно выполняемых операций.
 * Этот конкретный пример ограничивает количество одновременно работающих горутин с помощью семафора.
 * Пример использования: ограничение количества одновременных HTTP-запросов к API, чтобы не превысить лимиты.
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

// ограничение кол-ва горутин одновременно работающих (100)
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
	ctx := context.Background()
	c := client2{}
	requests := make([]Request2, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request2{Payload: strconv.Itoa(i)}
	}
	c.WithLimiterSemaphore(ctx, requests)
}
