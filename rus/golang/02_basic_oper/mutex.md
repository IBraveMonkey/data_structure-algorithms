# 🔒 Мьютексы: Полный разбор (Mutex Deep Dive)

Мьютексы (`sync.Mutex` и `sync.RWMutex`) — это основные инструменты синхронизации. Ошибки в их использовании ведут к панике или вечной блокировке горутин (deadlock).

---


### 🚦 Причины паники и блокировок (Panic & Deadlock Causes)

| Ситуация            | Тип            | Результат                      |
| :------------------ | :------------- | :----------------------------- |
| **Повторный Lock**  | `sync.Mutex`   | 🔥 **Deadlock** (навсегда)     |
| **Unlock без Lock** | `sync.Mutex`   | 🔥 **PANIC**                   |
| **Unlock на RLock** | `sync.RWMutex` | 🔥 **PANIC**                   |
| **RUnlock на Lock** | `sync.RWMutex` | 🔥 **PANIC**                   |
| **RLock при Lock**  | `sync.RWMutex` | 🛑 **Block** (ожидание Unlock) |

---


### 💻 Примеры кода (Original Examples)

```go
package main

import (
	"fmt"
	"sync"
)

// Ошибка: panic при двойной блокировке sync.Mutex
func lockAnyTimes() {
	mutex := sync.Mutex{}
	mutex.Lock()
	// mutex.Lock() // 🛑 Deadlock
}

// Ошибка: panic при разблокировке без предварительной блокировки
func unlockWithoutLock() {
	mutex := sync.Mutex{}
	// mutex.Unlock() // 🔥 Panic!
}

// Нормально: mutex не запоминает свой контекст, любой может разблокировать
func unlockFromAnotherGoroutine() {
	mutex := sync.Mutex{}
	mutex.Lock()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		mutex.Unlock() // ✅ можешь разблокировать из другой горутины
	}()

	wg.Wait()

	mutex.Lock() // ✅ OK: можешь о5 заблокировать
	mutex.Unlock()
}

// Ошибка: panic при RUnlock на заблокированном мьютексе Lock()\Unlock
func RUnlockLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	// m.RUnlock() // 🔥 Panic!
}

// Ошибка: panic при Unlock на RLock
func UnlockRLockedMutex() {
	m := sync.RWMutex{}
	m.RLock()
	// m.Unlock() // 🔥 Panic!
}

// Ошибка: блокировка при RLock на заблокированном мьютексе
func LockRLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	// m.RLock() // 🛑 Block навсегда
}
```

---

> [!TIP]
> Всегда проверяйте соответствие: `Lock()` -> `Unlock()`, `RLock()` -> `RUnlock()`. Если их перепутать, программа немедленно упадет с паникой.

<!-- QUIZ_START 

[
    {
        "question": "Что произойдет при попытке вызвать Unlock() у мьютекса, который не был заблокирован через Lock()?",
        "options": [
            "Программа просто продолжит выполнение",
            "Произойдет panic",
            "Горутина заблокируется",
            "Мьютекс будет уничтожен"
        ],
        "correctIndex": 1
    },
    {
        "question": "Какое поведение ожидает горутину при повторном вызове Lock() на уже заблокированном тем же потоком sync.Mutex?",
        "options": [
            "Мьютекс поймет, что это та же горутина, и пропустит ее",
            "Произойдет panic",
            "Возникнет deadlock (вечная блокировка)",
            "Мьютекс разблокируется"
        ],
        "correctIndex": 2
    },
    {
        "question": "Можно ли разблокировать мьютекс из другой горутины, отличной от той, которая его заблокировала?",
        "options": [
            "Нет, это вызовет panic",
            "Да, в Go мьютекс не привязан к конкретной горутине",
            "Да, но только если это RWMutex",
            "Только с использованием пакета unsafe"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

