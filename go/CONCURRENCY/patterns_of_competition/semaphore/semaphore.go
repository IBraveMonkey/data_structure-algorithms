package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func downloadFile(ctx context.Context, url string) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	case <-ch:
		fmt.Printf("%s\n", url)
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	const goroutineLimit = 3

	files := []string{"https://1.ru", "https://2.ru", "https://3.ru", "https://4.ru", "https://5.ru", "https://6.ru", "https://7.ru"}

	wg := sync.WaitGroup{}
	semaphore := make(chan struct{}, goroutineLimit)

	wg.Add(len(files))
	for _, file := range files {
		semaphore <- struct{}{}
		go func() {
			defer func() {
				defer wg.Done()
				<-semaphore
			}()

			downloadFile(ctx, file)
		}()
	}

	wg.Wait()
}
