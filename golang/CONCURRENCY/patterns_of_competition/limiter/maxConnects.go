/*
 * Паттерн Limiter (Ограничитель) используется для ограничения количества одновременно выполняемых операций.
 * Этот конкретный пример ограничивает количество активных соединений (или горутин), используя пул воркеров.
 * Пример использования: ограничение количества одновременных HTTP-запросов к API, чтобы не превысить лимиты.
 */
package limiter

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Request1 представляет запрос для отправки
type Request1 struct {
	Payload string // Данные запроса
}

// Client1 интерфейс для работы с запросами
type Client1 interface {
	SendRequest(ctx context.Context, request Request1) error
	WithLimiter(ctx context.Context, requests []Request1)
}

type client1 struct{}

// SendRequest отправляет один запрос (имитация)
func (c client1) SendRequest(ctx context.Context, request Request1) error {
	time.Sleep(100 * time.Millisecond) // Имитация сетевого запроса
	fmt.Println("sending request", request.Payload)
	return nil
}

// Ограничение количества одновременно работающих коннектов (воркеров)
var maxConnects = 10

// WithLimiterWorkerPool ограничивает количество соединений через пул воркеров
func (c client1) WithLimiterWorkerPool(ctx context.Context, ch chan Request1) {
	wg := sync.WaitGroup{}

	// Запускаем фиксированное количество воркеров
	wg.Add(maxConnects) // Добавляем в счетчик количество воркеров

	for range maxConnects {
		go func() {
			defer wg.Done() // Уменьшаем счетчик при завершении воркера

			// Воркер обрабатывает все запросы из канала
			for req := range ch {
				c.SendRequest(ctx, req) // Обрабатываем запрос
			}
		}()
	}

	wg.Wait() // Ждем завершения всех воркеров
}

// maxConnectsFn демонстрирует использование паттерна
func maxConnectsFn() {
	ctx := context.Background() // Создаем контекст
	c := client1{}

	// Создаем 1000 запросов
	requests := make([]Request1, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request1{Payload: strconv.Itoa(i)}
	}

	// Обрабатываем запросы с ограничением в 10 воркеров
	c.WithLimiterWorkerPool(ctx, generate(requests))
}

// generate создает канал и отправляет в него все запросы
func generate(reqs []Request1) chan Request1 {
	// Создаем небуферизованный канал
	ch := make(chan Request1)

	// Горутина для отправки запросов в канал
	go func() {
		for _, v := range reqs {
			ch <- v // Отправляем каждый запрос
		}
		close(ch) // Закрываем канал после отправки всех запросов
	}()

	return ch
}
