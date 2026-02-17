# 🗺️ Map: Deep Dive

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

<!-- QUIZ_START 

[
    {
        "question": "Which operation on a nil map will cause a panic?",
        "options": [
            "Reading a value by key",
            "Deleting a key using delete()",
            "Writing a value by key",
            "Iterating via range"
        ],
        "correctIndex": 2
    },
    {
        "question": "What does a map return when trying to read a value for a non-existent key?",
        "options": [
            "Panic",
            "nil",
            "The zero-value of the map's value type",
            "A runtime error"
        ],
        "correctIndex": 2
    },
    {
        "question": "Is it safe to delete a key from a map if that key doesn't exist?",
        "options": [
            "Yes, it is a safe 'no-op' operation",
            "No, it causes a panic",
            "Only if the map is not nil",
            "Yes, but only for string keys"
        ],
        "correctIndex": 0
    }
]

QUIZ_END -->

