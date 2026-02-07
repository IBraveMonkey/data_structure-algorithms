### ⚡ Channels: Deep Dive

In Go, working with channels requires understanding their states. Below are all main scenarios: from blocks and panics to correct operation.

---

### 🚦 State Matrix

| Operation | Nil Channel | Open Channel | Closed Channel |
| :--- | :--- | :--- | :--- |
| **Read** (`<-ch`) | 🛑 Block forever | ✅ OK / Block | ✅ Zero-value + `ok=false` |
| **Write** (`ch <- v`) | 🛑 Block forever | ✅ OK / Block | 🔥 **PANIC** |
| **Close** (`close(ch)`) | 🔥 **PANIC** | ✅ OK | 🔥 **PANIC** |

---

### 💻 Code Examples (Original Examples)

```go
package main

import (
	"fmt"
)

// Error: Block forever on writing to nil channel
// Ошибка: Блокировка навсегда при записи в nil канал
func writeToNilChannel() {
	var ch chan int
	ch <- 1
}

// Error: Block forever on reading from nil channel
// Ошибка: Блокировка навсегда при чтении из nil канала
func redToNilChannel() {
	var ch chan int
	<-ch
}

// Error: panic on writing to closed channel
// Ошибка: panic при записи в закрытый канал
func writeToClosedChannel() {
	ch := make(chan int, 2)
	close(ch)
	ch <- 20
}

// Error: block on range iteration over nil channel
// Ошибка: блокировка при итерации по nil каналу
func rangeNilChannel() {
	var ch chan int
	for range ch {
	}
}

// Error: panic on closing nil channel
// Ошибка: panic при закрытии nil канала
func closeNilChannel() {
	var ch chan int
	close(ch)
}

// Error: panic on closing channel multiple times
// Ошибка: panic при повторном закрытии канала
func closeChannelAnyTimes() {
	ch := make(chan int)
	close(ch)
	close(ch)
}

// Works: compares channels
// Работает: сравнивает каналы
func compareChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	equal1 := ch1 == ch2
	equal2 := ch1 == ch1

	fmt.Println(equal1)
	fmt.Println(equal2)
}

// Works: read from channel, close it, and read from closed channel
// Работает: читаем из канала, закрывает его и читаем из закрытого канала
func readFromChannel() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20

	val, ok := <-ch
	fmt.Println(val, ok)

	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)
}

// Works: read from one of channels via select
// Работает: читаем из одного из каналов через select
func readAnyChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 100
	}()

	go func() {
		ch2 <- 200
	}()

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}
}
```

---

> [!CAUTION]
> Writing to a closed channel or closing it multiple times are the most common causes of panics. Always think through the channel ownership logic.
