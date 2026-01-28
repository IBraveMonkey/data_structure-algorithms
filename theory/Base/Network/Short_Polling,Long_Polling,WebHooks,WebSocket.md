**Short Polling**, **Long Polling**, **Webhooks**, и **WebSockets** — это технологии для обмена данными между клиентом (например, браузером) и сервером, часто используемые для обновления данных в реальном времени или около реального времени. Они решают задачу доставки событий или изменений от сервера к клиенту, но различаются по механизму, производительности и сценариям применения.

---

## Short Polling

### О чем

**Short Polling** — клиент периодически отправляет HTTP-запросы к серверу (например, каждые 5 секунд), чтобы проверить наличие новых данных. Сервер отвечает немедленно, даже если данных нет.

**Механизм**:

1. Клиент отправляет GET-запрос.
2. Сервер возвращает данные (или пустой ответ).
3. Клиент ждет заданный интервал и повторяет запрос.

### Пример

Веб-приложение для отображения новых сообщений в чате:

- **Сервер** (Go): Возвращает список новых сообщений.
- **Клиент** (JavaScript): Запрашивает сообщения каждые 5 секунд.

**Код (Сервер, Go)**:

```go
package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var (
	messages []Message
	mutex    sync.Mutex
)

func getMessages(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	json.NewDecoder(r.Body).Decode(&msg)
	mutex.Lock()
	messages = append(messages, msg)
	mutex.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/messages", getMessages)
	http.HandleFunc("/add", addMessage)
	http.ListenAndServe(":8080", nil)
}
```

**Код (Клиент, JavaScript)**:

```javascript
function pollMessages() {
  fetch("http://localhost:8080/messages")
    .then((response) => response.json())
    .then((messages) => {
      console.log("Новые сообщения:", messages);
      // Обновить UI
    })
    .catch((err) => console.error(err))
    .finally(() => {
      setTimeout(pollMessages, 5000); // Повтор каждые 5 секунд
    });
}
pollMessages();
```

### Преимущества

- Простота реализации.
- Совместимость с любым HTTP-сервером.
- Не требует постоянного соединения.

### Недостатки

- Высокая нагрузка на сервер из-за частых запросов.
- Задержка обновлений (зависит от интервала).
- Неэффективно при редких обновлениях.

### Когда использовать

- Простые приложения с редкими обновлениями (например, обновление статуса заказа).
- Когда WebSockets или Long Polling недоступны.
- Прототипы или MVP.

---

## Long Polling

### О чем

**Long Polling** — клиент отправляет HTTP-запрос, но сервер удерживает соединение открытым, пока не появятся новые данные или не истечет таймаут. После ответа клиент сразу отправляет новый запрос.

**Механизм**:

1. Клиент отправляет GET-запрос.
2. Сервер ждет данных (удерживает соединение).
3. При появлении данных сервер отвечает, клиент запрашивает снова.

### Пример

Чат-приложение с Long Polling:

- **Сервер** (Go): Удерживает запрос, пока не появится новое сообщение.
- **Клиент** (JavaScript): Повторяет запрос после ответа.

**Код (Сервер, Go)**:

```go
package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var (
	messages    []Message
	mutex       sync.Mutex
	listeners   = make(chan chan []Message)
	broadcast   = make(chan Message)
)

func longPoll(w http.ResponseWriter, r *http.Request) {
	ch := make(chan []Message)
	listeners <- ch
	select {
	case messages := <-ch:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	case <-time.After(30 * time.Second):
		w.WriteHeader(http.StatusNoContent)
	}
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	json.NewDecoder(r.Body).Decode(&msg)
	mutex.Lock()
	messages = append(messages, msg)
	mutex.Unlock()
	broadcast <- msg
	w.WriteHeader(http.StatusCreated)
}

func handleMessages() {
	for {
		select {
		case ch := <-listeners:
			select {
			case msg := <-broadcast:
				mutex.Lock()
				ch <- messages
				mutex.Unlock()
			case <-time.After(30 * time.Second):
				ch <- nil
			}
		}
	}
}

func main() {
	go handleMessages()
	http.HandleFunc("/poll", longPoll)
	http.HandleFunc("/add", addMessage)
	http.ListenAndServe(":8080", nil)
}
```

**Код (Клиент, JavaScript)**:

```javascript
function longPoll() {
  fetch("http://localhost:8080/poll")
    .then((response) => response.json())
    .then((messages) => {
      if (messages) {
        console.log("Новые сообщения:", messages);
        // Обновить UI
      }
      longPoll(); // Повторный запрос
    })
    .catch((err) => {
      console.error(err);
      setTimeout(longPoll, 1000); // Повтор при ошибке
    });
}
longPoll();
```

### Преимущества

- Меньшая задержка, чем у Short Polling.
- Меньше запросов при редких обновлениях.
- Совместимость с HTTP.

### Недостатки

- Удержание соединений нагружает сервер.
- Сложнее масштабировать.
- Таймауты требуют повторных запросов.

### Когда использовать

- Приложения с нечастыми, но важными обновлениями (например, уведомления).
- Когда WebSockets недоступны, но нужна меньшая задержка.
- Средние по сложности проекты.

---

## Webhooks

### О чем

**Webhooks** — сервер отправляет HTTP POST-запрос на URL клиента, когда происходит событие. Это "обратный" подход: сервер инициирует связь, а не клиент.

**Механизм**:

1. Клиент регистрирует URL для получения событий.
2. Сервер отправляет POST-запрос на этот URL при событии.
3. Клиент обрабатывает запрос.

### Пример

Система уведомлений о новых заказах:

- **Сервер** (Go): Отправляет webhook при новом заказе.
- **Клиент** (Go): Получает и обрабатывает webhook.

**Код (Сервер, Go)**:

```go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Order struct {
	ID     int    `json:"id"`
	Amount float64 `json:"amount"`
}

func notifyWebhook(order Order) {
	payload, _ := json.Marshal(order)
	http.Post("http://localhost:8081/webhook", "application/json", bytes.NewBuffer(payload))
}

func addOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	go notifyWebhook(order)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/orders", addOrder)
	http.ListenAndServe(":8080", nil)
}
```

**Код (Клиент, Go)**:

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	ID     int    `json:"id"`
	Amount float64 `json:"amount"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	log.Printf("Получен заказ: ID=%d, Amount=%.2f", order.ID, order.Amount)
	// Обновить UI или базу данных
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	http.ListenAndServe(":8081", nil)
}
```

### Преимущества

- Минимальная нагрузка на клиент (нет опросов).
- Мгновенные уведомления.
- Простота интеграции с API.

### Недостатки

- Требует доступного URL клиента.
- Сложности с надежностью (нужны повторы при сбоях).
- Не подходит для частых обновлений.

### Когда использовать

- Асинхронные события (например, платежи, обновления в CRM).
- Интеграция между системами (GitHub, Stripe).
- Когда клиент не может поддерживать постоянное соединение.

---

## WebSockets

### О чем

**WebSockets** — двунаправленное постоянное соединение между клиентом и сервером, позволяющее обмениваться данными в реальном времени без HTTP-запросов.

**Механизм**:

1. Клиент устанавливает WebSocket-соединение через HTTP handshake.
2. Сервер и клиент обмениваются сообщениями по одному соединению.
3. Соединение остается открытым до разрыва.

### Пример

Чат-приложение с WebSockets:

- **Сервер** (Go): Управляет WebSocket-соединениями.
- **Клиент** (JavaScript): Отправляет и получает сообщения.

**Код (Сервер, Go)**:

```go
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true
	for {
		var msg string
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	go handleMessages()
	http.HandleFunc("/ws", handleConnections)
	http.ListenAndServe(":8080", nil)
}
```

**Код (Клиент, JavaScript)**:

```javascript
const ws = new WebSocket("ws://localhost:8080/ws");

ws.onopen = () => {
  console.log("Подключено к WebSocket");
};

ws.onmessage = (event) => {
  console.log("Получено сообщение:", event.data);
  // Обновить UI
};

ws.onclose = () => {
  console.log("WebSocket закрыт");
};

function sendMessage() {
  const msg = "Привет!";
  ws.send(JSON.stringify(msg));
}
```

### Преимущества

- Минимальная задержка, двунаправленность.
- Эффективно для частых обновлений.
- Поддержка реального времени.

### Недостатки

- Сложность масштабирования (постоянные соединения).
- Требует поддержки WebSocket на сервере и клиенте.
- Высокое потребление ресурсов.

### Когда использовать

- Приложения реального времени (чаты, игры, стримы).
- Частые двунаправленные обновления.
- Когда важна минимальная задержка.

---

## Сравнение технологий

|                   |                            |              |                    |           |                                 |
| ----------------- | -------------------------- | ------------ | ------------------ | --------- | ------------------------------- |
| Технология        | Механизм                   | Задержка     | Нагрузка на сервер | Сложность | Применение                      |
| **Short Polling** | Периодические HTTP-запросы | Высокая      | Высокая            | Низкая    | Редкие обновления, прототипы    |
| **Long Polling**  | Удержание HTTP-соединения  | Средняя      | Средняя            | Средняя   | Уведомления, нечастые события   |
| **Webhooks**      | Сервер инициирует POST     | Низкая       | Низкая             | Средняя   | Асинхронные события, интеграции |
| **WebSockets**    | Постоянное двунаправленное | Очень низкая | Высокая            | Высокая   | Реальное время, чаты, игры      |

### Из одной серии?

Да, все эти технологии решают задачу доставки данных от сервера к клиенту или между системами, но отличаются подходом:

- **Short/Long Polling**: Клиент инициирует запросы (pull).
- **Webhooks**: Сервер инициирует уведомления (push).
- **WebSockets**: Постоянное соединение для push/pull.

### Рекомендации

- **Short Polling**: Используйте для простых приложений или когда другие методы недоступны. Например, проверка статуса раз в минуту.
- **Long Polling**: Подходит для уведомлений с умеренной частотой, если WebSockets не поддерживаются. Например, обновления в дашборде.
- **Webhooks**: Идеальны для интеграций и асинхронных событий. Например, уведомления о вебхуках GitHub.
- **WebSockets**: Лучший выбор для реального времени. Например, чаты, игровые платформы.

### Примечания

- **Масштабируемость**: WebSockets и Long Polling сложнее масштабировать из-за удержания соединений. Используйте брокеры сообщений (Redis, Kafka) для крупных систем.
- **Совместимость**: Short Polling работает везде, WebSockets требуют поддержки протокола.
- **Ресурсы**: Webhooks и Short Polling менее требовательны к серверу, чем WebSockets.
