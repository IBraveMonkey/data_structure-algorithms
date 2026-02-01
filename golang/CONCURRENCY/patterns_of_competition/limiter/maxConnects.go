/*
 * Паттерн Limiter (Ограничитель) используется для ограничения количества одновременно выполняемых операций.
 * Этот конкретный пример ограничивает количество активных соединений (или горутин), используя пул воркеров.
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

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request)
}

type client struct{}

func (c client) SendRequest(ctx context.Context, request Request) error {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("sending request", request.Payload)
	return nil
}

// ограничение кол-ва коннектов(работающих горутин)
var maxConnects = 10

func (c client) WithLimiterWorkerPool(ctx context.Context, ch chan Request) {
	wg := sync.WaitGroup{}

	wg.Add(maxConnects)
	for range maxConnects {
		go func() {
			defer wg.Done()
			for req := range ch {
				c.SendRequest(ctx, req)
			}
		}()
	}

	wg.Wait()
}

func maxConnectsFn() {
	ctx := context.Background()
	c := client{}
	requests := make([]Request, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request{Payload: strconv.Itoa(i)}
	}
	c.WithLimiterWorkerPool(ctx, generate(requests))
}

func generate(reqs []Request) chan Request {
	ch := make(chan Request)

	go func() {
		for _, v := range reqs {
			ch <- v
		}
		close(ch)
	}()

	return ch
}
