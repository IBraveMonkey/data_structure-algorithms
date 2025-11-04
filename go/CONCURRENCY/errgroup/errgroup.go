package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

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
		return "", errors.New("timed out")
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
		fmt.Println("an error occured: ", err.Error())
	}

	fmt.Println("time: ", time.Since(start))
	fmt.Println(res)
}
