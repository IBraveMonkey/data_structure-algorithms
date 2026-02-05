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

// User представляет пользователя для обработки
type User struct {
	Name string // Имя пользователя
}

// fetch имитирует получение данных пользователя (например, из API)
func fetch(ctx context.Context, user User) (string, error) {
	ch := make(chan struct{}) // Канал для сигнала завершения

	// Горутина для имитации асинхронной работы
	go func() {
		time.Sleep(10 * time.Millisecond) // Имитация задержки запроса
		close(ch)                         // Сигнализируем о завершении
	}()

	// Ждем завершения работы или отмены контекста
	select {
	case <-ch:
		return user.Name, nil // Успешное завершение
	case <-ctx.Done():
		return "", ctx.Err() // Контекст отменен или истек таймаут
	}
}

// process обрабатывает список пользователей параллельно с помощью errgroup
func process(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0) // Карта для подсчета пользователей
	mu := sync.Mutex{}                 // Мьютекс для защиты карты от гонки данных

	// Создаем errgroup с контекстом (автоматическая отмена при ошибке)
	egroup, ectx := errgroup.WithContext(ctx)
	egroup.SetLimit(100) // Ограничиваем количество одновременных горутин

	// Запускаем горутину для каждого пользователя
	for _, u := range users {
		egroup.Go(func() error {
			// Получаем данные пользователя с контекстом errgroup
			name, err := fetch(ectx, u)
			if err != nil {
				return err // Возвращаем ошибку (остановит другие горутины через контекст)
			}

			// Безопасно обновляем карту с блокировкой
			mu.Lock()
			defer mu.Unlock()
			names[name] = names[name] + 1 // Увеличиваем счетчик для имени
			return nil
		})
	}

	// Ждем завершения всех горутин или первой ошибки
	if err := egroup.Wait(); err != nil {
		return nil, err // Возвращаем ошибку если была
	}

	return names, nil // Возвращаем результат
}

func main() {
	// Тестовые данные
	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background() // Создаем базовый контекст

	start := time.Now()
	// Обрабатываем пользователей параллельно
	res, err := process(ctx, names)
	if err != nil {
		fmt.Println("an error occured:", err.Error())
	}

	fmt.Print("time: ", time.Since(start)) // Время выполнения
	fmt.Println(res)                       // Результат: map с количеством каждого имени
}

// ===================================================================================
// Продвинутая реализация с ограничением количества горутин
// ===================================================================================
// Эта реализация использует errgroup.SetLimit для ограничения количества одновременно работающих горутин.
// Это полезно, когда нужно избежать перегрузки системы или превышения лимитов на количество соединений.

// fetchAdvanced имитирует получение данных с явной обработкой таймаута
func fetchAdvanced(ctx context.Context, user User) (string, error) {
	ch := make(chan struct{}) // Канал для сигнала завершения

	// Горутина для имитации асинхронной работы
	go func() {
		time.Sleep(10 * time.Millisecond) // Имитация задержки запроса
		close(ch)                         // Сигнализируем о завершении
	}()

	// Ждем завершения работы или отмены контекста
	select {
	case <-ch:
		return user.Name, nil // Успешное завершение
	case <-ctx.Done():
		return "", errors.New("timed out") // Явная ошибка таймаута
	}
}

// processAdvanced показывает расширенное использование errgroup с SetLimit
func processAdvanced(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0) // Карта для подсчета пользователей
	mu := sync.Mutex{}                 // Мьютекс для защиты карты от гонки данных

	// Создаем errgroup с контекстом
	egroup, ectx := errgroup.WithContext(ctx)
	egroup.SetLimit(100) // Ограничиваем количество одновременно работающих горутин

	// ВАЖНО: SetLimit означает что одновременно будет работать максимум 100 горутин.
	// Если запущено больше задач, остальные будут ждать своей очереди.

	// Запускаем горутину для каждого пользователя
	for _, u := range users {
		egroup.Go(func() error {
			// Получаем данные пользователя с контекстом errgroup
			name, err := fetchAdvanced(ectx, u)
			if err != nil {
				return err // Возвращаем ошибку (остановит другие горутины)
			}

			// Безопасно обновляем карту с блокировкой
			mu.Lock()
			defer mu.Unlock()
			names[name] = names[name] + 1 // Увеличиваем счетчик для имени

			return nil
		})
	}

	// Ждем завершения всех горутин или первой ошибки
	if err := egroup.Wait(); err != nil {
		return nil, err // Возвращаем ошибку если была
	}

	return names, nil // Возвращаем результат
}

// mainAdvanced демонстрирует использование расширенной версии
func mainAdvanced() {
	// Тестовые данные
	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background() // Создаем базовый контекст

	start := time.Now()
	// Обрабатываем пользователей параллельно с ограничением 100 горутин
	res, err := processAdvanced(ctx, names)
	if err != nil {
		fmt.Println("an error occured: ", err.Error())
	}

	fmt.Println("time: ", time.Since(start)) // Время выполнения
	fmt.Println(res)                         // Результат: map с количеством каждого имени
}
