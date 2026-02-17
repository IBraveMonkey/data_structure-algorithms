# 👷 Паттерн Worker Pool (Пул воркеров)

**Worker Pool** — это паттерн, который ограничивает количество одновременно запущенных горутин для обработки очереди задач. Это позволяет эффективно использовать ресурсы процессора и памяти, не допуская их перегрузки при большом количестве входящих заданий.

---


### 🧠 Концепция

Представьте мастерскую, где работают три мастера (воркера). На столе лежит стопка заказов (канал задач). Мастера по очереди берут заказы, выполняют их и складывают готовые изделия в коробку (канал результатов).

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph TD
    Tasks[Task Queue] --> W1[Worker 1]
    Tasks --> W2[Worker 2]
    Tasks --> W3[Worker 3]
    W1 --> Res[Results Channel]
    W2 --> Res
    W3 --> Res
    linkStyle default stroke:#009688,stroke-width:2px;







```

---


### 💻 Реализация

В этом примере мы создаем пул из нескольких воркеров для имитации обработки файлов.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Task представляет задачу для обработки
type Task struct {
	ID       int
	Filename string
}

// Worker читает задачи из канала и обрабатывает их
func Worker(id int, taskCh <-chan Task, resCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskCh {
		// Имитация обработки (1 секунда)
		fmt.Printf("Воркер %d начал обработку %s\n", id, task.Filename)
		time.Sleep(time.Second)
		resCh <- fmt.Sprintf("Воркер %d завершил %s", id, task.Filename)
	}
}

func runWorkerPoolDemo() {
	const numWorkers = 3
	const numTasks = 5

	taskCh := make(chan Task, numTasks)
	resCh := make(chan string, numTasks)
	var wg sync.WaitGroup

	// Запускаем пул воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, taskCh, resCh, &wg)
	}

	// Отправляем задачи в пул
	go func() {
		for i := 1; i <= numTasks; i++ {
			taskCh <- Task{ID: i, Filename: fmt.Sprintf("image_%d.jpg", i)}
		}
		close(taskCh) // Закрываем канал задач, когда всё отправили
	}()

	// Ждем завершения воркеров и закрываем канал результатов
	go func() {
		wg.Wait()
		close(resCh)
	}()

	fmt.Println("Пул воркеров запущен...")

	// Читаем результаты
	for res := range resCh {
		fmt.Println(res)
	}
	fmt.Println("Все задачи завершены.")
}

func main() {
	runWorkerPoolDemo()
}
```

---


### 💡 Особенности

1. **Контроль ресурсов**: Вы точно знаете, сколько памяти и процессорного времени потребляет приложение, так как количество горутин фиксировано.
2. **Очередность**: Использование буферизованных каналов позволяет сглаживать пиковые нагрузки.
3. **Безопасность**: Задачи не теряются, а ждут своего воркера в очереди.

> [!IMPORTANT]
> Всегда закрывайте канал задач (`taskCh`), чтобы воркеры могли выйти из цикла `range` и завершить свою работу.
