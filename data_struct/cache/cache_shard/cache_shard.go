/*
Cache Sharding (Шардирование кэша)

Что это такое?
Шардирование кэша - это метод распределения данных кэша по нескольким независимым сегментам (шардам).
Каждый шард - это отдельный блок памяти с собственной синхронизацией, что позволяет уменьшить
блокировки и повысить производительность при многопоточном доступе.

Зачем это нужно?
- Повысить конкурентность при работе с кэшем
- Уменьшить contention (ожидание блокировок) при параллельном доступе
- Распределить нагрузку между несколькими сегментами

В чём смысл?
- Разделить кэш на несколько независимых частей
- Использовать хэш ключа для определения нужного шарда
- Каждый шард управляет доступом к своим данным независимо

Когда использовать?
- Когда кэш используется в многопоточной среде
- При высокой нагрузке с частыми операциями чтения/записи
- Когда lock contention становится узким местом
*/

package main

import (
	"fmt"

	"hash/fnv"
	"sync"
)

// ICache - интерфейс для кэша
type ICache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

// Shard - сегмент кэша с собственной синхронизацией
type Shard struct {
	data map[string]string
	mu   *sync.RWMutex
}

// Cache - структура кэша, состоящая из нескольких шардов
type Cache struct {
	shards []*Shard
}

// New - создает новый шардированный кэш с указанным количеством шардов
func New(shardCount int64) *Cache {

	shards := make([]*Shard, shardCount) // [shard, shard, shard{mu }]

	for i := range shards {
		shards[i] = &Shard{
			data: make(map[string]string),
			mu:   &sync.RWMutex{}, // Добавляем mutex для каждого шарда
		}
	}

	return &Cache{
		shards: shards,
	}
}

// Set - добавляет пару ключ-значение в кэш
func (c *Cache) Set(k string, v string) {
	shard := c.getShard(k)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[k] = v
}

// Get - возвращает значение по ключу из кэша
func (c *Cache) Get(k string) (string, bool) {
	shard := c.getShard(k)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	if val, ok := shard.data[k]; ok {
		return val, true
	}

	return "", false
}

// getShard - возвращает шард для указанного ключа, используя хэширование
func (c *Cache) getShard(key string) *Shard {
	hasher := fnv.New32()
	_, _ = hasher.Write([]byte(key))
	hash := hasher.Sum32()

	return c.shards[hash%uint32(len(c.shards))]
}

func main() {
	cache := New(10)
	cache.Set("salary", "500000")

	value, ok := cache.Get("salary")
	if !ok {
		fmt.Println("Value not found")
		return
	}

	fmt.Println("Value:", value)
}
