package cache_ttl

import (
	"context"
	"fmt"
	"time"
)

// Example демонстрирует использование кэша с TTL
func Example() {
	// Создаем кэш с временем жизни 100 мс
	cache := New(100 * time.Millisecond)
	ctx := context.Background()

	// Добавляем значение
	cache.Set(ctx, "user:1", "John Doe")
	fmt.Println("Добавили 'user:1' -> 'John Doe'")

	// Получаем значение
	val, err := cache.Get(ctx, "user:1")
	if err == nil {
		fmt.Printf("Получили: %s\n", val)
	}

	// Ждем истечения TTL
	fmt.Println("Ждем 200 мс...")
	time.Sleep(200 * time.Millisecond)

	// Пытаемся получить устаревшее значение
	val, err = cache.Get(ctx, "user:1")
	if err != nil {
		fmt.Printf("Ошибка получения (ожидалось): %v\n", err)
	} else {
		fmt.Printf("Неожиданно получили: %s\n", val)
	}

	// Останавливаем очистку
	cache.Stop()
}
