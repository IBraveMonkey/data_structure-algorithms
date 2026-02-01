package main

import "sync"

// Ошибка: panic при двойной блокировке sync.Mutex
func lockAnyTimes() {
	mutex := sync.Mutex{}
	mutex.Lock()
	mutex.Lock()
}

// Ошибка: panic при разблокировке без предварительной блокировки
func unlockWithoutLock() {
	mutex := sync.Mutex{}
	mutex.Unlock()
}

// Нормально: mutex не запоминает свой :контекст:, может любой разблокировать
func unlockFromAnotherGoroutine() {
	mutex := sync.Mutex{}
	mutex.Lock()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		mutex.Unlock()
	}()

	wg.Wait()

	mutex.Lock()
	mutex.Unlock()
}

// Ошибка: panic при RUnlock на заблокированном мьютексе Lock()\Unlock и RLock()\RUnlock - так должно быть
func RUnlockLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	m.RUnlock()
}

// Ошибка: panic при Unlock на RLock
func UnlockRLockedMutex() {
	m := sync.RWMutex{}
	m.RLock()
	m.Unlock()
}

// Ошибка: блокировка при RLock на заблокированном мьютексе
func LockRLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	m.RLock()
}

func main() {
}
