# 🔄 Шардированный кэш

**Описание**: 
Шардирование кэша (Cache Sharding) — это способ масштабирования кэша путем его разделения на несколько независимых фрагментов ("шардов"). В многопоточных программах это помогает избавиться от "пробок" при доступе к данным.

- **Как это устроено внутри**: Если у нас есть миллион записей в одной хеш-таблице под одним мьютексом (замком), то тысячи потоков будут выстраиваться в очередь, чтобы прочитать хоть что-то. Шардирование делит этот миллион на 10-20 маленьких таблиц, у каждой из которых свой замок. Хеш-функция определяет, в какой шард положить данные. Теперь потоки могут работать параллельно, если они обращаются к разным шардам.
- **Аналогия**: Представьте очередь в супермаркете с одной-единственной кассой. Это кэш без шардирования. Шардирование — это открытие 10 касс. Покупатели распределяются по разным очередям, и общая скорость обслуживания возрастает в разы.


### Преимущества и недостатки
✅ **Плюсы**:
1. **Высокая параллельность**: Позволяет многим потокам одновременно читать и писать данные без ожидания друг друга.
2. **Снижение нагрузки на CPU**: Процессор тратит меньше времени на управление блокировками и переключение задач.

❌ **Минусы**:
1. **Сложность**: Нужно гарантировать, что один и тот же ключ всегда попадает в один и тот же шард (эффективная хеш-функция).
2. **Перекос данных**: Если хеш-функция плохая, один шард может быть перегружен, а остальные будут простаивать.

---

ICache - интерфейс для кэша
Shard - сегмент кэша с собственной синхронизацией
Cache - структура кэша, состоящая из нескольких шардов
New - создает новый шардированный кэш с указанным количеством шардов
Set - добавляет пару ключ-значение в кэш
Get - возвращает значение по ключу из кэша
getShard - возвращает шард для указанного ключа, используя хэширование


## 💻 Реализация

```go
package main

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

// Shard представляет собой отдельный сегмент кэша с собственным мьютексом
type Shard struct {
	sync.RWMutex
	data map[string]interface{}
}

// Cache представляет собой шардированный кэш
type Cache struct {
	shardCount int
	shards     []*Shard
}

// NewCache создает новый шардированный кэш
func NewCache(shardCount int) *Cache {
	cache := &Cache{
		shardCount: shardCount,
		shards:     make([]*Shard, shardCount),
	}

	for i := 0; i < shardCount; i++ {
		cache.shards[i] = &Shard{
			data: make(map[string]interface{}),
		}
	}

	return cache
}

// getShardIndex вычисляет индекс шарда для данного ключа
func (c *Cache) getShardIndex(key string) int {
	hash := sha1.Sum([]byte(key))
	// Берем первый байт хэша для определения индекса
	return int(hash[0]) % c.shardCount
}

// Set добавляет или обновляет значение в кэше
func (c *Cache) Set(key string, value interface{}) {
	shardIndex := c.getShardIndex(key)
	shard := c.shards[shardIndex]

	shard.Lock()
	defer shard.Unlock()

	shard.data[key] = value
}

// Get возвращает значение из кэша по ключу
func (c *Cache) Get(key string) (interface{}, bool) {
	shardIndex := c.getShardIndex(key)
	shard := c.shards[shardIndex]

	shard.RLock()
	defer shard.RUnlock()

	val, ok := shard.data[key]
	return val, ok
}

// Delete удаляет ключ из кэша
func (c *Cache) Delete(key string) {
	shardIndex := c.getShardIndex(key)
	shard := c.shards[shardIndex]

	shard.Lock()
	defer shard.Unlock()

	delete(shard.data, key)
}

// Stats возвращает статистику распределения данных по шардам
func (c *Cache) Stats() map[int]int {
	stats := make(map[int]int)
	for i, shard := range c.shards {
		shard.RLock()
		stats[i] = len(shard.data)
		shard.RUnlock()
	}
	return stats
}

func main() {
	cache := NewCache(4) // Создаем кэш с 4 шардами

	cache.Set("user:1", "Борис")
	cache.Set("user:2", "Ильяс")

	if val, ok := cache.Get("user:1"); ok {
		fmt.Printf("Найдено: %v\n", val)
	}

	fmt.Println("Статистика шардов:", cache.Stats())
}
```

```javascript
/**
 * ShardedCache - реализация шардированного кэша на JavaScript.
 * В JS (Node.js) один поток (Event Loop), поэтому мьютексы не нужны,
 * но шардирование полезно для уменьшения размера отдельных Map
 * и логического разделения данных.
 */
class ShardedCache {
  constructor(shardCount = 8) {
    this.shardCount = shardCount;
    // Создаем массив шардов, каждый из которых - отдельный Map
    this.shards = Array.from({ length: shardCount }, () => new Map());
  }

  /**
   * Простая хеш-функция для строк (DJB2)
   */
  _hash(key) {
    let hash = 5381;
    for (let i = 0; i < key.length; i++) {
      hash = (hash * 33) ^ key.charCodeAt(i);
    }
    // Приводим к беззнаковому 32-битному целому
    return hash >>> 0;
  }

  /**
   * Получить индекс шарда для ключа
   */
  _getShardIndex(key) {
    const hash = this._hash(key.toString());
    return hash % this.shardCount;
  }

  /**
   * Установить значение
   */
  set(key, value) {
    const index = this._getShardIndex(key);
    this.shards[index].set(key, value);
  }

  /**
   * Получить значение
   */
  get(key) {
    const index = this._getShardIndex(key);
    return this.shards[index].get(key);
  }

  /**
   * Проверить существование ключа
   */
  has(key) {
    const index = this._getShardIndex(key);
    return this.shards[index].has(key);
  }

  /**
   * Удалить ключ
   */
  delete(key) {
    const index = this._getShardIndex(key);
    return this.shards[index].delete(key);
  }

  /**
   * Очистить весь кэш
   */
  clear() {
    this.shards.forEach(shard => shard.clear());
  }

  /**
   * Получить статистику заполнения шардов
   */
  getStats() {
    return this.shards.map((shard, index) => ({
      shard: index,
      size: shard.size
    }));
  }
}

// Пример использования
const cache = new ShardedCache(4);

cache.set("session_123", { user: "admin" });
cache.set("session_456", { user: "guest" });

console.log(cache.get("session_123")); // { user: "admin" }
console.log("Stats:", cache.getStats());

```

<!-- QUIZ_START 
[
    {
        "question": "Какую основную проблему решает шардирование кэша?",
        "options": ["Увеличение объема используемой памяти", "Блокировки (contention) при доступе к единому мьютексу в многопоточной среде", "Защита от вирусов", "Автоматический перевод названий ключей"],
        "correctIndex": 1
    },
    {
        "question": "Как определяется, в какой именно шард попадут данные при записи?",
        "options": ["Пользователь выбирает вручную", "Случайным образом", "С помощью хеш-функции от ключа", "По размеру значения"],
        "correctIndex": 2
    },
    {
        "question": "В чем заключается главный минус шардирования при плохой хеш-функции?",
        "options": ["Цвет данных меняется на красный", "Программа начинает работать медленнее из-за длинных имен переменных", "Возможен перекос (skew), когда один шард перегружен, а другие пусты", "Шарды нельзя объединить обратно"],
        "correctIndex": 2
    }
]
QUIZ_END -->

```

