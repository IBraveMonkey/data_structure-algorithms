# 🛡️ Limiter: Max Connections

**Max Connections Limiter** — это вариация паттерна ограничения ресурсов, которая фокусируется на контроле количества активных сетевых соединений или воркеров. Это гарантирует, что ваша система не откроет больше сокетов или не задействует больше вычислительных ресурсов, чем может выдержать инфраструктура.

---


### 🧠 Концепция

Это похоже на работу колл-центра, где есть только 10 операторов. Даже если звонят 100 человек одновременно, только 10 смогут разговаривать, остальные будут ждать в очереди, пока не освободится линия.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph TD
    U[Users/Requests] --> Q[Queue]
    Q --> Pool{Worker Pool: 10}
    Pool --> W1[Worker]
    Pool --> W2[Worker]
    Pool --> WN[Worker 10]
    linkStyle default stroke:#009688,stroke-width:2px;







```

---


### 💻 Реализация

В этой реализации мы ограничиваем количество одновременно работающих горутин-обработчиков (воркеров), которые читают из общего канала задач.

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func runMaxConnectsDemo() {
	const maxConnections = 3 // Лимит соединений
	const totalRequests = 10

	taskCh := make(chan int, totalRequests)
	var wg sync.WaitGroup

	fmt.Printf("Запуск системы с лимитом %d соединений...\n", maxConnections)

	// Создаем пул воркеров, который и является нашим лимитером
	for i := 1; i <= maxConnections; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for taskID := range taskCh {
				fmt.Printf("Воркер %d обрабатывает запрос %d\n", workerID, taskID)
				time.Sleep(500 * time.Millisecond) // Имитация сетевой работы
			}
		}(i)
	}

	// Отправляем запросы
	for i := 1; i <= totalRequests; i++ {
		taskCh <- i
	}
	close(taskCh)

	wg.Wait()
	fmt.Println("Все запросы обработаны.")
}

func main() {
	runMaxConnectsDemo()
}
```

---


### 💡 Особенности

1. **Стабильность**: Предотвращает падение сервиса из-за нехватки файловых дескрипторов или памяти.
2. **Предсказуемость**: Вы имеете четкий контроль над пропускной способностью системы.
3. **Очередность**: Входящие запросы не отклоняются, а ставятся в очередь (буфер канала).

> [!NOTE]
> Этот подход часто называют "Пул воркеров", и он является наиболее надежным способом реализации жестких ограничений в Go.
