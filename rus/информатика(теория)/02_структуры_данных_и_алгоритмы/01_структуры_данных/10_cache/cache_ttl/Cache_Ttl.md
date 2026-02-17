# ⏲️ Кэш с TTL

**Описание**: 
Кэш с TTL (Time To Live) — это структура данных, которая умеет "забывать" устаревшую информацию. Каждому элементу при записи выдается срок годности, после которого он считается недействительным.

- **Как это устроено внутри**: При сохранении данных алгоритм записывает не только значение, но и метку времени смерти (`expiry_time = now + TTL`). 
  - *Ленивое удаление*: Элемент проверяется на "свежесть" только в момент обращения (Get). Если время вышло — он удаляется.
  - *Активное удаление*: Фоновый процесс (воркер) периодически сканирует память и вычищает "трупы" данных, чтобы они не занимали место.
- **Аналогия**: Представьте холодильник. Вы кладете туда молоко с чеком, на котором написана дата истечения срока. Если вы достали молоко и увидели, что дата прошла — вы его выкидываете. А раз в неделю вы проводите полную ревизию и выкидываете всё просроченное, даже если не собирались это есть.


### Преимущества и недостатки
✅ **Плюсы**:
1. **Свежесть данных**: Вы можете быть уверены, что пользователь не увидит информацию столетней давности.
2. **Авто-очистка**: Память не забивается мусором бесконечно, данные сами "уходят" со временем.

❌ **Минусы**:
1. **Дополнительные расходы**: Хранение меток времени и работа фонового чистильщика потребляют ресурсы CPU и RAM.
2. **Риск "протухания"**: Если TTL слишком мал, вы будете слишком часто ходить в основную базу данных, убивая смысл кэша.

---

elem - элемент кэша с данными и временем жизни
Cache - структура кэша с TTL
New - создает новый кэш с указанным TTL
Set - добавляет пару ключ-значение в кэш с TTL
Stop - останавливает фоновую очистку кэша
Get - возвращает значение по ключу из кэша, проверяя TTL
Проверяем, не истекло ли время жизни
Del - удаляет элемент из кэша
clearByTTL - запускает фоновую очистку просроченных элементов
delete - удаляет элемент из кэша
clear - удаляет все просроченные элементы из кэша


## 💻 Реализация

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrNotFound = errors.New("key not found or expired")

// Item представляет элемент кэша с метаданными о времени истечения
type Item struct {
	Value      interface{}
	ExpiryTime time.Time
}

// TTLCache — реализация кэша с временем жизни элементов
type TTLCache struct {
	sync.RWMutex
	items      map[string]Item
	cleanupInt time.Duration
	stopCh     chan struct{}
}

// NewTTLCache создает новый кэш и запускает фоновую очистку
func NewTTLCache(cleanupInterval time.Duration) *TTLCache {
	cache := &TTLCache{
		items:      make(map[string]Item),
		cleanupInt: cleanupInterval,
		stopCh:     make(chan struct{}),
	}

	// Запускаем фоновый процесс очистки (активное удаление)
	go cache.startCleanup()

	return cache
}

// Set добавляет элемент в кэш с заданным TTL
func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		ExpiryTime: time.Now().Add(ttl),
	}
}

// Get возвращает значение, если оно существует и не просрочено
func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	// Ленивая проверка при обращении
	if time.Now().After(item.ExpiryTime) {
		return nil, false
	}

	return item.Value, true
}

// Delete принудительно удаляет ключ
func (c *TTLCache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.items, key)
}

// startCleanup периодически сканирует кэш на предмет просроченных данных
func (c *TTLCache) startCleanup() {
	ticker := time.NewTicker(c.cleanupInt)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanup()
		case <-c.stopCh:
			return
		}
	}
}

func (c *TTLCache) cleanup() {
	c.Lock()
	defer c.Unlock()

	now := time.Now()
	for key, item := range c.items {
		if now.After(item.ExpiryTime) {
			delete(c.items, key)
		}
	}
}

// Stop останавливает фоновую очистку
func (c *TTLCache) Stop() {
	close(c.stopCh)
}

func main() {
	// Очистка каждые 100мс
	cache := NewTTLCache(100 * time.Millisecond)
	defer cache.Stop()

	cache.Set("token", "secret_data", 500*time.Millisecond)

	val, ok := cache.Get("token")
	fmt.Printf("Get: %v, OK: %v\n", val, ok)

	fmt.Println("Ждем 1 секунду...")
	time.Sleep(1 * time.Second)

	val, ok = cache.Get("token")
	fmt.Printf("Get (after TTL): %v, OK: %v\n", val, ok)
}
```

```javascript
/**
 * TTLCache - реализация кэша с временем жизни (Time To Live).
 */
class TTLCache {
  constructor(defaultTTLMs = 60000) {
    this.storage = new Map();
    this.defaultTTL = defaultTTLMs;
    
    // Опционально: активная очистка каждые 30 секунд
    this.cleanupTimer = setInterval(() => this.cleanup(), 30000);
  }

  /**
   * Установить значение с TTL
   */
  set(key, value, ttlMs = this.defaultTTL) {
    const expiry = Date.now() + ttlMs;
    this.storage.set(key, { value, expiry });
  }

  /**
   * Получить значение (с ленивой проверкой)
   */
  get(key) {
    const item = this.storage.get(key);
    
    if (!item) return null;

    // Если время вышло - удаляем и возвращаем null
    if (Date.now() > item.expiry) {
      this.storage.delete(key);
      return null;
    }

    return item.value;
  }

  /**
   * Удалить ключ
   */
  delete(key) {
    return this.storage.delete(key);
  }

  /**
   * Активная очистка просроченных элементов
   */
  cleanup() {
    const now = Date.now();
    for (const [key, item] of this.storage.entries()) {
      if (now > item.expiry) {
        this.storage.delete(key);
      }
    }
  }

  /**
   * Остановить фоновую очистку
   */
  stop() {
    clearInterval(this.cleanupTimer);
  }

  /**
   * Получить размер кэша
   */
  size() {
    return this.storage.size;
  }
}

// Пример использования
const cache = new TTLCache(1000); // 1 сек дефолтный TTL

cache.set("user_session", { id: 1 }, 500); // на 500 мс

setTimeout(() => {
  console.log("Get after 200ms:", cache.get("user_session")); // { id: 1 }
}, 200);

setTimeout(() => {
  console.log("Get after 700ms:", cache.get("user_session")); // null
}, 700);

```


## 🚀 Практические задачи
```go
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
	cache.Set(ctx, "user:1", "Борис Иванов")
	fmt.Println("Добавили 'user:1' -> 'Борис Иванов'")

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

<!-- QUIZ_START 
[
    {
        "question": "Что означает аббревиатура TTL в контексте кэширования?",
        "options": ["Total Time Lost", "Time To Live (Срок годности данных)", "Table To List", "Transistor-Transistor Logic"],
        "correctIndex": 1
    },
    {
        "question": "В чем разница между 'ленивым' и 'активным' удалением просроченных данных?",
        "options": ["Ленивое удаляет только большие файлы, активное - все", "Ленивое проверяет TTL при запросе (Get), активное использует фоновый процесс-чистильщик", "Ленивое работает только ночью", "Разницы нет"],
        "correctIndex": 1
    },
    {
        "question": "Какой риск возникает при слишком малом значении TTL?",
        "options": ["Процессор будет работать слишком тихо", "Слишком частые обращения к основной базе данных (Cache Miss), что сводит пользу кэша к нулю", "Данные станут слишком свежими", "Память переполнится"],
        "correctIndex": 1
    }
]
QUIZ_END -->

```
