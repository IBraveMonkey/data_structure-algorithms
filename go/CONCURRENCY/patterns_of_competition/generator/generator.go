package genarator

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

// Как сделать чтобы функция работало только какое-то время - 3s
func PredictableTimeWork() {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("Done")
	case <-time.After(3 * time.Second):
		fmt.Println("Cancel with timeout")
	}
}

// Микропаттерн - Генератор, который создает канал и возращает его
func writer() <-chan int {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// Должна быть не блокируемой, надо делать в отдельной горутине
	go func() {
		defer wg.Done()
		for i := range 5 {
			ch <- i + 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 5 {
			ch <- i + 11
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func Generator() {
	ch := writer()

	for v := range ch {
		fmt.Println("v =", v)
	}
}
