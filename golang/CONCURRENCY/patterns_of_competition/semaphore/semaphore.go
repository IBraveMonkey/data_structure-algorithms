package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
 * Паттерн Semaphore (Семафор) используется для ограничения количества одновременно выполняемых горутин.
 * Он позволяет контролировать доступ к ресурсам, которые имеют ограниченную пропускную способность.
 * Пример использования: ограничение количества одновременных HTTP-запросов к API, чтобы не превысить лимиты.
 */

// downloadFile имитирует загрузку файла с контекстом для отмены
func downloadFile(ctx context.Context, url string) {
	// Создаем канал для сигнала завершения загрузки
	ch := make(chan struct{})

	// Запускаем горутину для имитации загрузки
	go func() {
		time.Sleep(1 * time.Second) // Имитация загрузки (1 секунда)
		close(ch)                   // Сигнализируем о завершении
	}()

	// Ждем либо завершения загрузки, либо отмены контекста
	select {
	case <-ctx.Done():
		// Контекст был отменен (таймаут или cancel)
		fmt.Println(ctx.Err())
		return
	case <-ch:
		// Загрузка успешно завершена
		fmt.Printf("%s\n", url)
		return
	}
}

func main() {
	// Создаем контекст с таймаутом 3 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Обязательно вызываем cancel для освобождения ресурсов

	const goroutineLimit = 3 // Максимум 3 горутины одновременно

	// Список файлов для скачивания
	files := []string{"https://1.ru", "https://2.ru", "https://3.ru", "https://4.ru", "https://5.ru", "https://6.ru", "https://7.ru"}

	wg := sync.WaitGroup{}
	// Семафор: буферизованный канал на goroutineLimit элементов
	semaphore := make(chan struct{}, goroutineLimit)

	wg.Add(len(files)) // Добавляем количество файлов в счетчик

	// Запускаем горутину для каждого файла
	for _, file := range files {
		semaphore <- struct{}{} // Занимаем слот в семафоре (блокируется если слотов нет)

		go func() {
			defer func() {
				defer wg.Done() // Уменьшаем счетчик при завершении
				<-semaphore     // Освобождаем слот в семафоре
			}()

			downloadFile(ctx, file) // Выполняем загрузку
		}()
	}

	wg.Wait() // Ждем завершения всех горутин
}
