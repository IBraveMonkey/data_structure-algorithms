### 🗺️ Map: Deep Dive

Maps in Go are hash tables. The main "gotchas" are related to uninitialized (nil) maps and concurrent access.

---

### 🚦 Operations with Nil Map (Nil Map Operations)

| Operation | Result |
| :--- | :--- |
| **Read** (`v := m[k]`) | ✅ Safe (returns zero-value) |
| **Delete** (`delete(m, k)`) | ✅ Safe (no-op) |
| **Iteration** (`range m`) | ✅ Safe (0 iterations) |
| **Write** (`m[k] = v`) | 🔥 **PANIC** |

---

### 💻 Code Examples (Original Examples)

```go
package main

import (
	"fmt"
)

// Reading from nil map -> OK (returns zero-value)
// Читает из nil-мапы -> нормально (возвращает zero-value)
func readFromNilMap() {
	var data map[int]int
	_ = data[100]
}

// Deleting from nil map -> Safe (no-op)
// Удаляет из nil-мапы -> безопасно (ничего не происходит)
func deleteFromNilMap() {
	var data map[int]int
	delete(data, 100)
}

// Writing to nil map -> PANIC
// Пишет в nil-мапу -> PANIC
func writeToNilMap() {
	var data map[int]int
	data[100] = 100
}

// Iterating over nil map -> Safe (0 iterations)
// Итерирует по nil-мапе -> безопасно (0 итераций)
func rangeByNilMap() {
	var data map[int]int
	for range data {
	}
}

// Rewriting existing key -> OK
// Перезаписывает существующий ключ -> OK
func rewriteExistingKey() {
	data := make(map[int]int)
	data[100] = 500
	data[100] = 1000
}

// Deleting non-existing key -> Safe
// Удаляет несуществующий ключ -> безопасно
func deleteNonExistingKey() {
	data := make(map[int]int)
	delete(data, 100)
}
```

---

> [!IMPORTANT]
> Maps are **not thread-safe**. For concurrent access, use `sync.Mutex` or `sync.Map`. Attempting simultaneous writes will cause a `fatal error: concurrent map writes`.
