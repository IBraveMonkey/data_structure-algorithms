package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Паттерн Fan-In (Веер-в) используется для объединения нескольких каналов в один канал.
 * Он позволяет собирать результаты из нескольких источников в одном месте.
 * Пример использования: сбор результатов из нескольких HTTP-запросов в один канал для дальнейшей обработки.
 */

// MergeChannels объединяет несколько входящих каналов в один выходной канал
func MergeChannels(channels ...<-chan int) <-chan int {
	// Создаем результирующий канал для объединенных данных
	res := make(chan int)
	// WaitGroup для отслеживания завершения всех горутин-читателей
	wg := sync.WaitGroup{}

	// Добавляем в счетчик количество каналов (один читатель на канал)
	wg.Add(len(channels))

	// Запускаем отдельную горутину для каждого входного канала
	for _, channel := range channels {
		go func() {
			defer wg.Done()              // Уменьшаем счетчик после завершения чтения
			for value := range channel { // Читаем все значения из канала
				res <- value // Перенаправляем в общий результирующий канал
			}
		}()
	}

	// Отдельная горутина для закрытия результирующего канала
	go func() {
		wg.Wait()  // Ждем, пока все каналы будут прочитаны
		close(res) // Закрываем результирующий канал
	}()

	// Возвращаем канал только для чтения
	return res
}

// Закомментированный альтернативный вариант той же функции
// func MergeChannels(channels ...<-chan int) <-chan int {
// 	wg := sync.WaitGroup{}
// 	wg.Add(len(channels))
//
// 	result := make(chan int)
// 	for _, channel := range channels {
// 		go func() {
// 			defer wg.Done()
// 			for value := range channel {
// 				result <- value
// 			}
// 		}()
// 	}
//
// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()
//
// 	return result
// }

func main() {
	// Создаем три канала для демонстрации
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запускаем горутину-отправителя данных в три канала
	go func() {
		// defer гарантирует закрытие всех каналов при завершении
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		// Генерируем 100 чисел, распределяя их по трем каналам
		for i := 0; i < 100; i += 3 {
			ch1 <- i                           // Канал 1: 0, 3, 6, 9, ...
			ch2 <- i + 1                       // Канал 2: 1, 4, 7, 10, ...
			ch3 <- i + 2                       // Канал 2: 2, 5, 8, 11, ...
			time.Sleep(100 * time.Millisecond) // Задержка для наглядности
		}
	}()

	// Объединяем три канала в один и читаем из него
	for value := range MergeChannels(ch1, ch2, ch3) {
		fmt.Println(value) // Выводим значения по мере их поступления
	}
}
