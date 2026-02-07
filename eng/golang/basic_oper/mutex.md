### 🔒 Mutexes: Complete Guide (Deep Dive)

Mutexes (`sync.Mutex` and `sync.RWMutex`) are primary synchronization tools. Usage errors lead to a panic or a permanent block (deadlock).

---

### 🚦 Causes of Panic and Blocks (Panic & Deadlock Causes)

| Situation | Type | Result |
| :--- | :--- | :--- |
| **Double Lock** | `sync.Mutex` | 🔥 **Deadlock** (forever) |
| **Unlock without Lock** | `sync.Mutex` | 🔥 **PANIC** |
| **Unlock on RLock** | `sync.RWMutex` | 🔥 **PANIC** |
| **RUnlock on Lock** | `sync.RWMutex` | 🔥 **PANIC** |
| **RLock while Locked** | `sync.RWMutex` | 🛑 **Block** (waiting for Unlock) |

---

### 💻 Code Examples (Original Examples)

```go
package main

import (
	"fmt"
	"sync"
)

// Error: panic on double locking sync.Mutex (leads to deadlock)
// Ошибка: panic при двойной блокировке sync.Mutex
func lockAnyTimes() {
	mutex := sync.Mutex{}
	mutex.Lock()
	// mutex.Lock() // 🔥 Deadlock
}

// Error: panic on unlock without prior lock
// Ошибка: panic при разблокировке без предварительной блокировки
func unlockWithoutLock() {
	mutex := sync.Mutex{}
	// mutex.Unlock() // 🔥 Panic!
}

// Normal: mutex does not remember its context, anyone can unlock
// Нормально: mutex не запоминает свой контекст, любой может разблокировать
func unlockFromAnotherGoroutine() {
	mutex := sync.Mutex{}
	mutex.Lock()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		mutex.Unlock() // ✅ OK: anyone can unlock in Go
	}()

	wg.Wait()

	mutex.Lock() // ✅ OK: can lock again
	mutex.Unlock()
}

// Error: panic on RUnlock on locked mutex (Lock-ed, not RLock-ed)
// Ошибка: panic при RUnlock на заблокированном мьютексе Lock()\Unlock
func RUnlockLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	// m.RUnlock() // 🔥 Panic!
}

// Error: panic on Unlock on RLock
// Ошибка: panic при Unlock на RLock
func UnlockRLockedMutex() {
	m := sync.RWMutex{}
	m.RLock()
	// m.Unlock() // 🔥 Panic!
}

// Error: deadlock/block on RLock on locked mutex
// Ошибка: блокировка при RLock на заблокированном мьютексе
func LockRLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	// m.RLock() // 🛑 Block forever
}
```

---

> [!TIP]
> Always check the match: `Lock()` -> `Unlock()`, `RLock()` -> `RUnlock()`. Mixing them up causes immediate panics.
