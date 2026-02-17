# ⛓️ Паттерн Pipeline (Конвейер)

**Pipeline** — это паттерн, который разбивает сложную задачу на последовательность отдельных этапов (стадий). Каждый этап представлен горутиной, которая читает данные из входного канала, обрабатывает их и отправляет результат в выходной канал для следующего этапа.

---


### 🧠 Концепция

Это как сборочная линия на заводе: одна машина делает каркас, вторая ставит колеса, третья красит кузов. Данные "текут" через серию фильтров и преобразователей.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph LR
    Src[Source] --> S1[Stage 1: Gen]
    S1 --> S2[Stage 2: Mul]
    S2 --> S3[Stage 3: Filter]
    S3 --> Sink[Consumer]
    linkStyle default stroke:#009688,stroke-width:2px;







```

---


### 💻 Реализация

В этом примере мы создаем двухстадийный конвейер: первый этап генерирует числа, а второй возводит их в квадрат.

```go
package main

import "fmt"

// gen — первая стадия: генерирует числа из слайса
func gen(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, number := range numbers {
			out <- number
		}
	}()
	return out
}

// mul — вторая стадия: возводит каждое число в квадрат
func mul(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for number := range in {
			out <- number * number
		}
	}()
	return out
}

func runPipelineDemo() {
	fmt.Println("Конвейер запущен...")

	// Создаем цепочку: gen -> mul
	source := gen(1, 2, 3, 4, 5) // Числа от 1 до 5
	pipeline := mul(source)      // Возведение в квадрат

	// Читаем финальный результат
	for value := range pipeline {
		fmt.Printf("Результат: %d\n", value)
	}
	fmt.Println("Обработка завершена.")
}

func main() {
	runPipelineDemo()
}
```

---


### 💡 Особенности

1. **Композиция**: Вы можете легко добавлять новые этапы в середину конвейера, не меняя существующий код.
2. **Параллелизм**: Каждый этап работает в своей горутине, что позволяет стадиям выполняться одновременно для разных порций данных.
3. **Экономия памяти**: Данные обрабатываются потоково (stream), что позволяет работать с объемами данных, превышающими оперативную память.

> [!TIP]
> Используйте Pipeline, когда вам нужно применить серию независимых преобразований к потоку данных.
