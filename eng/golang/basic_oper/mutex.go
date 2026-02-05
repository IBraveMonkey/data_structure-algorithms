package main

import "sync"

// Error: panic on double locking sync.Mutex
func lockAnyTimes() {
	mutex := sync.Mutex{}
	mutex.Lock()
	mutex.Lock()
}

// Error: panic on unlock without prior lock
func unlockWithoutLock() {
	mutex := sync.Mutex{}
	mutex.Unlock()
}

// Normal: mutex does not remember its context, anyone can unlock
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

// Error: panic on RUnlock on locked mutex Lock()\Unlock and RLock()\RUnlock - this is how it should be
func RUnlockLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	m.RUnlock()
}

// Error: panic on Unlock on RLock
func UnlockRLockedMutex() {
	m := sync.RWMutex{}
	m.RLock()
	m.Unlock()
}

// Error: deadlock on RLock on locked mutex
func LockRLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	m.RLock()
}

func main() {
}
