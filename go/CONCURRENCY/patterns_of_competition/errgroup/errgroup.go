package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
* Паттерн Errgroup используется для управления группой горутин, которые могут возвращать ошибки.
* Он позволяет запускать несколько горутин и ждать их завершения, при этом возвращая первую ошибку, если таковая произошла.
* Пример использования: выполнение нескольких HTTP-запросов параллельно и возврат результата или ошибки от самого медленного запроса.
 */
type User struct {
	Name string
}

func fetch(ctx context.Context, user User) (string, error) {
	ch := make(chan struct{})

	go func() {
		time.Sleep(10 * time.Millisecond)
		close(ch)
	}()

	select {
	case <-ch:
		return user.Name, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func process(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0)
	mu := sync.Mutex{}

	egroup, ectx := errgroup.WithContext(ctx)
	egroup.SetLimit(100)

	for _, u := range users {
		egroup.Go(func() error {
			name, err := fetch(ectx, u)
			if err != nil {
				return err
			}

			mu.Lock()
			defer mu.Unlock()
			names[name] = names[name] + 1
			return nil
		})
	}

	if err := egroup.Wait(); err != nil {
		return nil, err
	}

	return names, nil
}

func main() {
	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background()

	start := time.Now()
	res, err := process(ctx, names)
	if err != nil {
		fmt.Println("an error occured:", err.Error())
	}

	fmt.Print("time: ", time.Since(start))
	fmt.Println(res)
}

// Продвинутая реализация с ограничением количества горутин
// Эта реализация использует errgroup.SetLimit для ограничения количества одновременно работающих горутин.
// Это полезно, когда нужно избежать перегрузки системы или превышения лимитов на количество соединений.
func fetchAdvanced(ctx context.Context, user User) (string, error) {
	ch := make(chan struct{})

	go func() {
		time.Sleep(10 * time.Millisecond)
		close(ch)
	}()

	select {
	case <-ch:
		return user.Name, nil
	case <-ctx.Done():
		return "", errors.New("timed out")
	}
}

func processAdvanced(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0)
	mu := sync.Mutex{}

	egroup, ectx := errgroup.WithContext(ctx)
	egroup.SetLimit(100) // Ограничиваем количество одновременно работающих горутин

	for _, u := range users {
		egroup.Go(func() error {
			name, err := fetchAdvanced(ectx, u)
			if err != nil {
				return err
			}

			mu.Lock()
			defer mu.Unlock()
			names[name] = names[name] + 1

			return nil
		})
	}

	if err := egroup.Wait(); err != nil {
		return nil, err
	}

	return names, nil
}

func mainAdvanced() {
	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background()

	start := time.Now()
	res, err := processAdvanced(ctx, names)
	if err != nil {
		fmt.Println("an error occured: ", err.Error())
	}

	fmt.Println("time: ", time.Since(start))
	fmt.Println(res)
}
