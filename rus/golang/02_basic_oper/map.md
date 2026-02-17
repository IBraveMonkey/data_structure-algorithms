# 🗺️ Мапы: Полный разбор (Map Deep Dive)

Мапы в Go — это хеш-таблицы. Основные "подводные камни" связаны с неинициализированными (nil) мапами и конкурентным доступом.

---


### 🚦 Операции с Nil мапой (Nil Map Operations)

| Операция | Результат |
| :--- | :--- |
| **Чтение** (`v := m[k]`) | ✅ Безопасно (возвращает zero-value) |
| **Удаление** (`delete(m, k)`) | ✅ Безопасно (ничего не происходит) |
| **Итерация** (`range m`) | ✅ Безопасно (0 итераций) |
| **Запись** (`m[k] = v`) | 🔥 **PANIC** |

---


### 💻 Примеры кода (Original Examples)

```go
package main

import (
	"fmt"
)

// Читает из nil-мапы -> нормально (возвращает zero-value)
func readFromNilMap() {
	var data map[int]int
	_ = data[100]
}

// Удаляет из nil-мапы -> безопасно (ничего не происходит)
func deleteFromNilMap() {
	var data map[int]int
	delete(data, 100)
}

// Пишет в nil-мапу -> PANIC
func writeToNilMap() {
	var data map[int]int
	data[100] = 100
}

// Итерирует по nil-мапе -> безопасно (0 итераций)
func rangeByNilMap() {
	var data map[int]int
	for range data {
	}
}

// Перезаписывает существующий ключ -> OK
func rewriteExistingKey() {
	data := make(map[int]int)
	data[100] = 500
	data[100] = 1000
}

// Удаляет несуществующий ключ -> безопасно
func deleteNonExistingKey() {
	data := make(map[int]int)
	delete(data, 100)
}
```

---

> [!IMPORTANT]
> Мапа **не потокобезопасна**. Для конкурентного доступа используйте `sync.Mutex` или `sync.Map`. Попытка одновременной записи вызовет `fatal error: concurrent map writes`.

<!-- QUIZ_START 

[
    {
        "question": "Какая операция с nil-мапой приведет к панике?",
        "options": [
            "Чтение значения по ключу",
            "Удаление ключа через delete()",
            "Запись значения по ключу",
            "Итерация через range"
        ],
        "correctIndex": 2
    },
    {
        "question": "Что вернет мапа при попытке прочитать значение по несуществующему ключу?",
        "options": [
            "Panic",
            "nil",
            "Нулевое значение (zero-value) типа данных значения",
            "Ошибку выполнения"
        ],
        "correctIndex": 2
    },
    {
        "question": "Безопасно ли удалять ключ из мапы, если этого ключа в ней нет?",
        "options": [
            "Да, это безопасная операция 'no-op'",
            "Нет, это вызовет panic",
            "Только если мапа не nil",
            "Да, но только для строковых ключей"
        ],
        "correctIndex": 0
    }
]

QUIZ_END -->

