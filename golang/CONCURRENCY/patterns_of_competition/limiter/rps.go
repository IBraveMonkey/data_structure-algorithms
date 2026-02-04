/*
 * Паттерн RPS Limiter (Rate Per Second) ограничивает скорость выполнения операций.
 * В отличие от maxGoroutine, который ограничивает количество одновременных горутин,
 * RPS limiter ограничивает количество операций В СЕКУНДУ независимо от их длительности.
 * Пример использования: соблюдение лимитов API (например, не более 100 запросов в секунду).
 */
package limiter

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Request3 struct {
	Payload string // Данные запроса
}

type Client3 interface {
	SendRequest(ctx context.Context, request Request3) error
	WithLimiter(ctx context.Context, requests []Request3)
}

type client3 struct{}

// SendRequest выполняет HTTP-запрос (имитация)
func (c client3) SendRequest(ctx context.Context, request Request3) error {
	time.Sleep(10 * time.Millisecond) // Имитация быстрого HTTP-запроса
	fmt.Println("sending request", request.Payload)
	return nil
}

// WithLimiterRPS ограничивает скорость выполнения запросов
func (c client3) WithLimiterRPS(ctx context.Context, requests []Request3) {
	const rps = 10 // Максимум 10 запросов в секунду

	// Создаем тикер, который посылает сигнал 10 раз в секунду
	ticker := time.NewTicker(time.Second / rps)
	defer ticker.Stop() // Останавливаем тикер при выходе

	// Обрабатываем каждый запрос
	for _, req := range requests {
		<-ticker.C              // Ждем разрешения от тикера (1/10 секунды)
		c.SendRequest(ctx, req) // Выполняем запрос
	}
}

func rpsFn() {
	ctx := context.Background()
	c := client3{}

	// Создаем 100 запросов
	requests := make([]Request3, 100)
	for i := 0; i < 100; i++ {
		requests[i] = Request3{Payload: strconv.Itoa(i)}
	}

	start := time.Now()

	// Отправляем запросы с ограничением 10 RPS
	// Должno занять ~10 секунд (100 запросов / 10 RPS)
	c.WithLimiterRPS(ctx, requests)

	fmt.Printf("Обработка заняла: %v\n", time.Since(start))
}
