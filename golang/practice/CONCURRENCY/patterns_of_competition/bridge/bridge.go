package main

import (
	"fmt"
)

/*
 * Паттерн Bridge-channel используется для "выравнивания" (flattening) потока каналов.
 * Если у вас есть канал, из которого приходят другие каналы (chan <-chan interface{}),
 * паттерн Bridge позволяет читать значения напрямую из одного результирующего канала.
 */

func bridge(done <-chan interface{}, chanStream <-chan (<-chan interface{})) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)
		for {
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}

			for val := range orDone(done, stream) {
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()

	return valStream
}

// Вспомогательная функция для корректного закрытия
func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	// Создаем поток, генерирующий каналы
	genChanStream := func() <-chan (<-chan interface{}) {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 5; i++ {
				stream := make(chan interface{})
				go func(v int) {
					defer close(stream)
					stream <- v
				}(i)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	// Читаем из всех каналов как из одного
	for val := range bridge(done, genChanStream()) {
		fmt.Printf("Получено значение: %v\n", val)
	}
}
