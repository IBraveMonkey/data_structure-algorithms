package main

import (
	"fmt"
	"time"
)

/*
 * Паттерн Or-channel используется для объединения нескольких каналов в один,
 * который закрывается, как только закрывается ЛЮБОЙ из входящих каналов.
 * Это полезно, когда у вас есть несколько асинхронных задач, и вам достаточно
 * завершения любой из них (например, несколько таймаутов или источников данных).
 */

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Базовые случаи для рекурсии
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			// Рекурсивно объединяем каналы
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	// Ждем закрытия любого из этих каналов
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(10*time.Second),
	)

	fmt.Printf("Завершено через %v\n", time.Since(start))
}
