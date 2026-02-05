package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://yandex.ru",
		"https://amazon.com",
		"https://youtube.com",
	}

	process(context.Background(), urls)
}

var client http.Client
var maxConnects = 10

func process(ctx context.Context, urls []string) map[int]int {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	hMap := make(map[int]int, len(urls))
	ch := make(chan string)

	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()

	wg.Add(maxConnects)
	for range maxConnects {
		go func() {
			defer wg.Done()

			for url := range ch {
				req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

				res, err := client.Do(req)
				if err != nil {
					fmt.Println("error occurred", err.Error())
				}

				mu.Lock()
				defer mu.Unlock()
				hMap[res.StatusCode]++
			}
		}()
	}

	wg.Wait()
	return hMap
}
