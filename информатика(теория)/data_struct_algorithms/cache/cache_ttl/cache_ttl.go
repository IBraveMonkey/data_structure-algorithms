/*
Cache with TTL (Кэш с временем жизни)

Что это такое?
Кэш с TTL (Time To Live) - это структура данных, в которой каждый элемент имеет ограниченное
время жизни. После истечения времени жизни элемент автоматически удаляется из кэша,
освобождая память и обеспечивая актуальность данных.

Зачем это нужно?
- Автоматически удалять устаревшие данные
- Ограничивать объем используемой памяти
- Обеспечивать свежесть данных в кэше

В чём смысл?
- Каждому элементу назначается время жизни при добавлении
- Элементы удаляются по истечении времени жизни
- Поддерживается актуальность данных без ручной очистки

Когда использовать?
- Когда данные в кэше могут устаревать
- Для ограничения использования памяти
- Когда важно иметь актуальные данные
*/

package cache_ttl

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("key not found")

// elem - элемент кэша с данными и временем жизни
type elem struct {
	value    string
	exp_date time.Time
}

// Cache - структура кэша с TTL
type Cache struct {
	storage map[string]elem
	mu      *sync.RWMutex
	TTL     time.Duration // Время жизни элементов
	done    chan struct{} // Канал для остановки горутины очистки
}

// New - создает новый кэш с указанным TTL
func New(ttl time.Duration) *Cache {
	cache := &Cache{
		storage: make(map[string]elem),
		mu:      &sync.RWMutex{},
		TTL:     ttl,
		done:    make(chan struct{}),
	}

	cache.clearByTTL()

	return cache
}

// Set - добавляет пару ключ-значение в кэш с TTL
func (c *Cache) Set(_ context.Context, key, value string) error {

	c.mu.Lock()
	el := elem{
		value:    value,
		exp_date: time.Now().Add(c.TTL), // Устанавливаем время жизни
	}

	c.storage[key] = el
	c.mu.Unlock()

	return nil
}

// Stop - останавливает фоновую очистку кэша
func (c *Cache) Stop() {
	close(c.done)
}

// Get - возвращает значение по ключу из кэша, проверяя TTL
func (c *Cache) Get(_ context.Context, key string) (string, error) {

	c.mu.RLock()
	el, ok := c.storage[key]
	c.mu.RUnlock()

	if !ok {
		return "", ErrNotFound
	}

	// Проверяем, не истекло ли время жизни
	if el.exp_date.Before(time.Now()) {
		c.delete(key) // Удаляем просроченный элемент
		return "", ErrNotFound
	}

	return el.value, nil
}

// Del - удаляет элемент из кэша
func (c *Cache) Del(_ context.Context, key string) error {
	c.delete(key)
	return nil
}

// clearByTTL - запускает фоновую очистку просроченных элементов
func (c *Cache) clearByTTL() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				c.clear() // Очищаем просроченные элементы
			case <-c.done:
				return
			}
		}
	}()
}

// delete - удаляет элемент из кэша
func (c *Cache) delete(key string) {
	c.mu.Lock()
	delete(c.storage, key)
	c.mu.Unlock()
}

// clear - удаляет все просроченные элементы из кэша
func (c *Cache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, el := range c.storage {
		if el.exp_date.Before(time.Now()) {
			delete(c.storage, key)
		}
	}
}
