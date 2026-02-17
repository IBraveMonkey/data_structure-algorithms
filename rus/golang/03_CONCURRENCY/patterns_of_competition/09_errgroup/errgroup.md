# ❌ Паттерн ErrGroup

**ErrGroup** — это расширенная версия `sync.WaitGroup`, которая не только ждет завершения группы горутин, но и перехватывает первую возникшую ошибку. Если одна из горутин возвращает ошибку, `ErrGroup` может автоматически отменить контекст для всех остальных горутин в группе.

---


### 🧠 Концепция

Это как командная работа над проектом. Если хотя бы один специалист (горутина) обнаруживает критическую проблему (ошибку), он сообщает об этом менеджеру (ErrGroup), который тут же дает команду всем остальным прекратить работу, так как проект больше не может быть завершен успешно.

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph TD
    EG[ErrGroup Manager] --> G1[Goroutine 1]
    EG --> G2[Goroutine 2]
    EG --> G3[Goroutine 3]
    G2 -- "Returns Error" --> EG
    EG -- "Cancel Context" ---> G1
    EG -- "Cancel Context" ---> G3
    linkStyle default stroke:#009688,stroke-width:2px;







```

---


### 💻 Реализация

Для использования `ErrGroup` необходимо импортировать пакет `golang.org/x/sync/errgroup`.

```go
package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func runErrGroupDemo() {
	// Создаем ErrGroup и контекст, который отменится при первой ошибке
	g, ctx := errgroup.WithContext(context.Background())

	urls := []string{"site1.com", "site2.com", "error-site.com", "site3.com"}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			// Имитация запроса
			select {
			case <-time.After(500 * time.Millisecond):
				if url == "error-site.com" {
					return fmt.Errorf("ошибка при запросе к %s", url)
				}
				fmt.Printf("Успешный запрос к %s\n", url)
				return nil
			case <-ctx.Done():
				// Прерываем работу, если другая горутина вернула ошибку
				return ctx.Err()
			}
		})
	}

	// Ждем завершения или первой ошибки
	if err := g.Wait(); err != nil {
		fmt.Printf("Группа завершилась с ошибкой: %v\n", err)
	} else {
		fmt.Println("Все запросы выполнены успешно!")
	}
}

func main() {
	runErrGroupDemo()
}
```

---


### 💡 Особенности

1. **Конфиденциальность**: Возвращается только самая первая возникшая ошибка. Остальные маскируются.
2. **Управление лимитами**: С помощью метода `SetLimit(n)` можно ограничить количество одновременно работающих горутин.
3. **Автоматизация**: Больше не нужно вручную обрабатывать каналы для сбора ошибок из разных горутин.

> [!IMPORTANT]
> `ErrGroup` идеально подходит для параллельного выполнения независимых задач, результат которых обязан быть успешным для всех участников.
