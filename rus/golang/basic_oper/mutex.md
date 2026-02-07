### 🔒 Мьютексы: Полный разбор (Mutex Deep Dive)

Мьютексы (`sync.Mutex` и `sync.RWMutex`) — это основные инструменты синхронизации. Ошибки в их использовании ведут к панике или вечной блокировке горутин (deadlock).

---

### 🚦 Причины паники и блокировок (Panic & Deadlock Causes)

| Ситуация | Тип | Результат |
| :--- | :--- | :--- |
| **Повторный Lock** | `sync.Mutex` | 🔥 **Deadlock** (навсегда) |
| **Unlock без Lock** | `sync.Mutex` | 🔥 **PANIC** |
| **Unlock на RLock** | `sync.RWMutex` | 🔥 **PANIC** |
| **RUnlock на Lock** | `sync.RWMutex` | 🔥 **PANIC** |
| **RLock при Lock** | `sync.RWMutex` | 🛑 **Block** (ожидание Unlock) |

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
	// mutex.Lock() // � Deadlock
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
		mutex.Unlock() // ✅ OK: can unlock from different goroutine
	}()

	wg.Wait()

	mutex.Lock() // ✅ OK: can lock again
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
	// m.RLock() // 🛑 Block forever
}
```

---

> [!TIP]
> Всегда проверяйте соответствие: `Lock()` -> `Unlock()`, `RLock()` -> `RUnlock()`. Если их перепутать, программа немедленно упадет с паникой.
