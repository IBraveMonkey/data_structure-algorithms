package fan_out

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Паттерн Fan-Out (Веер-из) используется для распределения данных из одного канала в несколько других каналов.
 * Он позволяет обрабатывать данные несколькими горутинами параллельно.
 * Пример использования: обработка задач несколькими воркерами, где одна очередь задач распределяется между несколькими обработчиками.
 */

// worker обрабатывает задачи из входного канала
func worker(
	id int, // ID воркера для идентификации в логах
	input <-chan int, // Канал для чтения задач (только чтение)
	wg *sync.WaitGroup, // WaitGroup для синхронизации завершения
) {
	defer wg.Done()          // Уменьшаем счетчик при завершении воркера
	for num := range input { // Читаем задачи пока канал не закрыт
		fmt.Printf("Worker %d processing %d\n", id, num)
		time.Sleep(500 * time.Millisecond) // Имитация обработки
		fmt.Printf("Worker %d finished %d, result: %d\n", id, num, num*2)
	}
}

// FanOut демонстрирует паттерн Fan-Out
func FanOut() {
	const numWorkers = 3                     // Количество воркеров
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8} // Задачи для обработки

	// Создаем отдельный канал для каждого воркера
	inputs := make([]chan int, numWorkers)
	var wg sync.WaitGroup

	// Запускаем все воркеры
	for i := 0; i < numWorkers; i++ {
		inputs[i] = make(chan int)     // Создаем канал для воркера
		wg.Add(1)                      // Добавляем в счетчик
		go worker(i+1, inputs[i], &wg) // Запускаем воркера (ID с 1)
	}

	// Горутина для распределения задач между воркерами
	go func() {
		for i, num := range numbers {
			// Распределяем задачи по кругу (round-robin)
			inputs[i%numWorkers] <- num // Отправляем задачу в канал воркера
		}

		// Закрываем все каналы после распределения всех задач
		for _, in := range inputs {
			close(in) // Сигнализируем воркеру о завершении
		}
	}()

	wg.Wait() // Ждем завершения всех воркеров
	fmt.Println("All workers done")
}

// JoinChannels объединяет несколько каналов в один (обратная операция к Fan-Out)
func JoinChannels(chs ...<-chan int) <-chan int {
	// Создаем выходной канал для объединенных данных
	mergedCh := make(chan int)

	// Запускаем горутину для слияния каналов
	go func() {
		wg := &sync.WaitGroup{}

		// Добавляем горутину для каждого входного канала
		wg.Add(len(chs))

		// Читаем из каждого канала в отдельной горутине
		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()      // Уменьшаем счетчик при завершении
				for id := range ch { // Читаем все значения из канала
					mergedCh <- id // Перенаправляем в общий канал
				}
			}(ch, wg)

			wg.Wait()       // Ждем завершения всех горутин
			close(mergedCh) // Закрываем выходной канал
		}
	}()

	// Возвращаем канал только для чтения
	return mergedCh
}
