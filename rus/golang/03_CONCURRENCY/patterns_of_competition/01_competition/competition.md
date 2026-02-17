# 🏁 Паттерн Competition (Racing)

**Competition** (или _Racing_) — это паттерн конкурентности, при котором несколько горутин выполняют одну и ту же задачу, и мы используем результат той, которая завершилась **первой**. Все остальные результаты игнорируются.

Этот подход идеально подходит для систем с высокой доступностью, где мы можем отправить запрос нескольким серверам-репликам и взять ответ от самого быстрого.

---


### 🧠 Концепция

Представьте, что вы вызываете такси в трех разных приложениях. Машина, которая приедет первой, забирает вас, а остальные заказы вы просто отменяете. В Go это реализуется через каналы и конструкцию `select`.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
sequenceDiagram
    participant Main as Main Goroutine
    participant W1 as Worker 1
    participant W2 as Worker 2
    participant W3 as Worker 3

    Main->>W1: Start Task
    Main->>W2: Start Task
    Main->>W3: Start Task
    Note over W1,W3: Processing...
    W2-->>Main: First Result!
    Note right of Main: Use Result from W2
    Note over W1,W3: Other results discarded







```

---


### 💻 Реализация

Ниже представлен пример реализации паттерна, где мы имитируем запросы к нескольким источникам данных.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// simulateSearch имитирует поиск данных с задержкой
func simulateSearch(id int) <-chan string {
    res := make(chan string)
    go func() {
        // Рандомная задержка до 3 секунд
        delay := time.Duration(rand.Intn(3000)) * time.Millisecond
        time.Sleep(delay)
        res <- fmt.Sprintf("Результат от сервера %d (заняло %v)", id, delay)
    }()
    return res
}

func runCompetitionDemo() {
    // Запускаем 3 конкурентные задачи
    c1 := simulateSearch(1)
    c2 := simulateSearch(2)
    c3 := simulateSearch(3)

    fmt.Println("Поиск запущен на 3-х серверах...")

    // Паттерн Competition: берем то, что пришло первым
    select {
    case res := <-c1:
        fmt.Println("ПОБЕДИТЕЛЬ:", res)
    case res := <-c2:
        fmt.Println("ПОБЕДИТЕЛЬ:", res)
    case res := <-c3:
        fmt.Println("ПОБЕДИТЕЛЬ:", res)
    case <-time.After(2 * time.Second):
        // Таймаут, если никто не успел за 2 секунды
        fmt.Println("ОШИБКА: Превышено время ожидания")
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    runCompetitionDemo()
}
```

---


### 💡 Особенности

1. **Скорость**: Общее время выполнения равно времени выполнения самого быстрого узла.
2. **Нагрузка**: Обратной стороной является избыточное потребление ресурсов, так как мы запускаем $N$ задач для получения одного ответа.
3. **Отмена**: В реальных проектах рекомендуется использовать `context.Context` для отмены "проигравших" горутин, чтобы они не продолжали тратить ресурсы после того, как победитель определен.

> [!TIP]
> Используйте этот паттерн, когда задержка (latency) критичнее, чем стоимость вычислительных ресурсов.
