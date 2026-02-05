package main

import (
	"fmt"
)

/*
 * Паттерн Bridge-channel используется для "выравнивания" (flattening) потока каналов.
 * Если у вас есть канал, из которого приходят другие каналы (chan <-chan interface{}),
 * паттерн Bridge позволяет читать значения напрямую из одного результирующего канала.
 */

// bridge объединяет поток каналов в один канал значений
func bridge(done <-chan interface{}, chanStream <-chan (<-chan interface{})) <-chan interface{} {
	// Создаем результирующий канал для значений
	valStream := make(chan interface{})

	// Запускаем горутину для обработки потока каналов
	go func() {
		defer close(valStream) // Закрываем выходной канал при завершении

		for {
			var stream <-chan interface{} // Текущий обрабатываемый канал

			// Ждем следующий канал из потока или сигнала завершения
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					// Поток каналов закрыт, завершаем работу
					return
				}
				stream = maybeStream // Получили новый канал для чтения
			case <-done:
				// Получен сигнал завершения
				return
			}

			// Читаем все значения из текущего канала
			for val := range orDone(done, stream) {
				select {
				case valStream <- val: // Отправляем значение в результирующий канал
				case <-done:
					// Прерываем если получен сигнал завершения
					return
				}
			}
		}
	}()

	// Возвращаем канал только для чтения
	return valStream
}

// orDone - вспомогательная функция для корректного закрытия каналов с учетом done
func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	// Создаем выходной канал
	valStream := make(chan interface{})

	// Запускаем горутину для чтения с учетом done
	go func() {
		defer close(valStream) // Закрываем канал при выходе

		for {
			select {
			case <-done:
				// Получен сигнал завершения
				return
			case v, ok := <-c:
				if !ok {
					// Входной канал закрыт
					return
				}
				// Пытаемся отправить значение в выходной канал
				select {
				case valStream <- v: // Успешно отправили
				case <-done:
					// Прервано сигналом завершения
					return
				}
			}
		}
	}()

	return valStream
}

func main() {
	// Создаем канал для сигнала завершения
	done := make(chan interface{})
	defer close(done) // Закрываем при завершении main

	// genChanStream - генератор, который создает поток каналов
	genChanStream := func() <-chan (<-chan interface{}) {
		// Канал каналов (поток каналов)
		chanStream := make(chan (<-chan interface{}))

		go func() {
			defer close(chanStream) // Закрываем поток каналов при завершении

			// Создаем 5 каналов и отправляем их в поток
			for i := 0; i < 5; i++ {
				stream := make(chan interface{}) // Создаем новый канал

				// Горутина для отправки значения в канал
				go func(v int) {
					defer close(stream) // Закрываем канал после отправки
					stream <- v         // Отправляем значение
				}(i)

				chanStream <- stream // Отправляем канал в поток каналов
			}
		}()

		return chanStream
	}

	// Читаем из всех каналов как из одного благодаря bridge
	for val := range bridge(done, genChanStream()) {
		fmt.Printf("Получено значение: %v\n", val) // Выведет числа 0, 1, 2, 3, 4
	}
}
