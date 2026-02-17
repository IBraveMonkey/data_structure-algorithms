# 🚦 Паттерн Semaphore (Семафор)

**Semaphore** — это паттерн, используемый для ограничения количества одновременно выполняемых операций или доступа к ограниченному ресурсу. В Go семафор чаще всего реализуется с помощью буферизованного канала, где размер буфера определяет максимальное количество "слотов" (разрешений).

---


### 🧠 Концепция

Представьте парковку на 3 места. Когда машина заезжает, она занимает одно место (отправляет данные в канал). Если мест нет, следующая машина ждет у шлагбаума. Когда машина уезжает, место освобождается (читаем из канала), и следующая машина может заехать.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph TD
    S[Resource Source] -->Gate{Semaphore: Limit 3}
    Gate -->|Slot 1| G1[Goroutine]
    Gate -->|Slot 2| G2[Goroutine]
    Gate -->|Slot 3| G3[Goroutine]
    G4[Goroutine 4] -.->|Wait| Gate
    linkStyle default stroke:#009688,stroke-width:2px;







```

---


### 💻 Реализация

Классическая реализация семафора в Go через пустой `struct{}` в буферизованном канале.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func runSemaphoreDemo() {
	const goroutineLimit = 3
	tasks := []int{1, 2, 3, 4, 5, 6, 7}

	wg := sync.WaitGroup{}
	// Семафор: буферизованный канал на 3 элемента
	sem := make(chan struct{}, goroutineLimit)

	for _, task := range tasks {
		wg.Add(1)

		// Занимаем слот: если канал полон, блокируемся
		sem <- struct{}{}

		go func(id int) {
			defer wg.Done()
			// Освобождаем слот при завершении
			defer func() { <-sem }()

			fmt.Printf("Воркер %d начал работу...\n", id)
			time.Sleep(1 * time.Second) // Имитация работы
			fmt.Printf("Воркер %d закончил.\n", id)
		}(task)
	}

	wg.Wait()
	fmt.Println("Все задачи выполнены.")
}

func main() {
	runSemaphoreDemo()
}
```

---


### 💡 Особенности

1. **Гибкость**: Вы можете динамически менять лимит, если используете переменную для размера буфера при создании.
2. **Простота**: Не требует сложных внешних библиотек, достаточно встроенных возможностей языка.
3. **Веса (Weighted Semaphores)**: Если задачи требуют разного "веса" (например, одна задача занимает 2 слота), лучше использовать пакет `golang.org/x/sync/semaphore`.

> [!TIP]
> Используйте Семафор, когда вам нужно ограничить нагрузку на внешнюю систему (например, не более 5 одновременных запросов к базе данных).
