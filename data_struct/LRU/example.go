package lru

import "fmt"

// Example демонстрирует использование LRU кэша
func Example() {
	// Создаем LRU кэш с вместимостью 2
	cache := New(2)

	// Добавляем элементы
	cache.Put(1, 1) // кэш: [1=1]
	fmt.Println("Put(1, 1)")

	cache.Put(2, 2) // кэш: [2=2, 1=1]
	fmt.Println("Put(2, 2)")

	// Получаем элемент
	val := cache.Get(1)
	fmt.Printf("Get(1): %d (ожидается 1)\n", val) // кэш: [1=1, 2=2]

	// Добавляем еще один элемент (приведет к удалению наименее используемого - 2)
	cache.Put(3, 3) // кэш: [3=3, 1=1]
	fmt.Println("Put(3, 3) - вытесняет ключ 2")

	val = cache.Get(2)
	fmt.Printf("Get(2): %d (ожидается -1)\n", val) // 2 был удален

	cache.Put(4, 4) // кэш: [4=4, 3=3] - вытесняет 1
	fmt.Println("Put(4, 4) - вытесняет ключ 1")

	val = cache.Get(1)
	fmt.Printf("Get(1): %d (ожидается -1)\n", val)

	val = cache.Get(3)
	fmt.Printf("Get(3): %d (ожидается 3)\n", val)

	val = cache.Get(4)
	fmt.Printf("Get(4): %d (ожидается 4)\n", val)
}
