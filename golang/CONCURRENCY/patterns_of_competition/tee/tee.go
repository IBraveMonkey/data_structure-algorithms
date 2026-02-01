package main

import (
	"fmt"
)

/*
 * Паттерн Tee-channel используется для разделения одного входящего канала
 * на два (или более) выходящих канала. Это похоже на тройник в сантехнике:
 * данные из одного источника дублируются в несколько независимых потоков.
 * Важно: чтение из выходящих каналов должно происходить параллельно,
 * иначе блокировка одного канала заблокирует весь поток.
 */

func tee(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range in {
			// Используем локальные копии каналов для безопасной итерации
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil // Обнуляем, чтобы больше не слать в него в этой итерации
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}

func main() {
	done := make(chan interface{})
	defer close(done)

	in := make(chan interface{})
	go func() {
		defer close(in)
		for i := 1; i <= 5; i++ {
			in <- i
		}
	}()

	out1, out2 := tee(done, in)

	// Читаем из двух каналов одновременно
	for val1 := range out1 {
		fmt.Printf("Поток 1 получил: %v\n", val1)
		fmt.Printf("Поток 2 получил: %v\n", <-out2)
	}
}
