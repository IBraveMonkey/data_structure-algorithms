**Forward Proxy**, **Reverse Proxy**, **API Gateway**, и **Load Balancer** — сетевые компоненты для управления трафиком между клиентами и серверами. Они решают задачи маршрутизации, балансировки, безопасности и упрощения взаимодействия, но имеют разные цели и сценарии применения. Конспект объясняет их суть, различия, технологии и примеры использования с кодом на Go.

---

## Forward Proxy (Proxy)

### О чем

**Forward Proxy** (или просто Proxy) — сервер, который действует от имени клиента, перенаправляя его запросы к внешним серверам (например, веб-сайтам). Скрывает клиента от внешнего мира, обеспечивая анонимность, фильтрацию или кэширование.

**Функции**:

- Скрытие IP-адреса клиента.
- Фильтрация контента (например, блокировка сайтов).
- Кэширование для ускорения доступа.
- Обход географических ограничений.

**Популярные технологии**:

- **Squid**: Классический прокси-сервер с кэшированием.
- **Privoxy**: Прокси для фильтрации и анонимности.
- **Shadowsocks**: Прокси для обхода блокировок.
- **Tor**: Сеть для анонимности через прокси.

### Пример

Forward Proxy (Go) перенаправляет запросы клиента к внешнему серверу (например, example.com), добавляя заголовок для идентификации.

**Код (Go)**:

```go
package main

import (
	"io"
	"log"
	"net/http"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
	// Создаем новый запрос к целевому серверу
	req, err := http.NewRequest(r.Method, "http://example.com"+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Копируем заголовки клиента
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	req.Header.Add("X-Forwarded-For", r.RemoteAddr)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Копируем ответ клиенту
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", handleProxy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Пример работы**:

- Клиент настраивает браузер на прокси localhost:8080.
- Запрос к http://example.com идет через прокси, который добавляет заголовок X-Forwarded-For и возвращает ответ.

### Преимущества

- Анонимность клиента.
- Кэширование для экономии трафика.
- Фильтрация нежелательного контента.

### Недостатки

- Дополнительная задержка.
- Требует настройки на стороне клиента.
- Ограниченные возможности для серверной логики.

### Когда использовать

- Корпоративные сети для фильтрации или мониторинга.
- Обход блокировок (VPN-подобные сценарии).
- Кэширование в локальных сетях.

---

## Reverse Proxy

### О чем

**Reverse Proxy** — сервер, принимающий запросы от клиентов и перенаправляющий их к бэкенд-серверам, скрывая их структуру. Выполняет маршрутизацию, кэширование и SSL-терминацию.

**Функции**:

- Скрытие бэкенд-серверов.
- Кэширование ответов.
- Базовая балансировка нагрузки.
- SSL-терминация.

**Популярные технологии**:

- **Nginx**: Высокопроизводительный прокси.
- **Apache HTTP Server**: Универсальный веб-сервер.
- **HAProxy**: Прокси для TCP/HTTP.
- **Caddy**: Прокси с автоматическим HTTPS.

### Пример

Reverse Proxy (Go) направляет запросы по путям к двум бэкенд-серверам.

**Код (Go)**:

```go
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	static, _ := url.Parse("http://localhost:8081")
	api, _ := url.Parse("http://localhost:8082")

	staticProxy := httputil.NewSingleHostReverseProxy(static)
	apiProxy := httputil.NewSingleHostReverseProxy(api)

	mux := http.NewServeMux()
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		staticProxy.ServeHTTP(w, r)
	})
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		apiProxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

**Пример работы**:

- Запрос http://localhost:8080/static/images → localhost:8081.
- Запрос http://localhost:8080/api/users → localhost:8082.

### Преимущества

- Простота маршрутизации.
- Скрытие серверов.
- Поддержка кэширования/SSL.

### Недостатки

- Ограниченная функциональность.
- Не подходит для сложных API.
- Меньше возможностей для мониторинга.

### Когда использовать

- Маршрутизация статического контента и API.
- Кэширование веб-страниц.
- SSL-терминация.

---

## API Gateway

### О чем

**API Gateway** — Reverse Proxy, специализированный для управления API. Обеспечивает единую точку входа, добавляя аутентификацию, лимитирование, мониторинг и трансформацию.

**Функции**:

- Аутентификация (JWT, OAuth).
- Лимитирование скорости.
- Логирование/мониторинг.
- Агрегация данных.

**Популярные технологии**:

- **Kong**: API Gateway с плагинами.
- **AWS API Gateway**: Облачный шлюз.
- **Tyk**: Легковесный шлюз.
- **Traefik**: Прокси с API-функциями.
- **Apigee**: Корпоративный шлюз.

### Пример

API Gateway (Go) с аутентификацией и маршрутизацией.

**Код (Go)**:

```go
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func main() {
	users, _ := url.Parse("http://localhost:8081")
	orders, _ := url.Parse("http://localhost:8082")

	usersProxy := httputil.NewSingleHostReverseProxy(users)
	ordersProxy := httputil.NewSingleHostReverseProxy(orders)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", authenticate(func(w http.ResponseWriter, r *http.Request) {
		usersProxy.ServeHTTP(w, r)
	}))
	mux.HandleFunc("/api/orders", authenticate(func(w http.ResponseWriter, r *http.Request) {
		ordersProxy.ServeHTTP(w, r)
	}))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

**Пример работы**:

- Запрос GET /api/users с токеном → localhost:8081.
- Без токена → 401 Unauthorized.

### Преимущества

- Управление API.
- Безопасность и мониторинг.
- Агрегация данных.

### Недостатки

- Сложность настройки.
- Возможное узкое место.
- Требует ресурсов.

### Когда использовать

- Микросервисы с API.
- Аутентификация/мониторинг.
- Агрегация данных.

---

## Load Balancer

### О чем

**Load Balancer** распределяет трафик между серверами для масштабируемости и отказоустойчивости. Работает на уровнях L4 (TCP/UDP) или L7 (HTTP).

**Функции**:

- Балансировка (round-robin, least connections).
- Health checks.
- Автомасштабирование.

**Популярные технологии**:

- **Nginx**: L7-балансировщик.
- **HAProxy**: L4/L7-балансировщик.
- **AWS ELB**: Облачный балансировщик.
- **F5 BIG-IP**: Корпоративный.
- **Traefik**: Для контейнеров.

### Пример

Load Balancer (Go) с round-robin.

**Код (Go)**:

```go
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Backend struct {
	URL   *url.URL
	Proxy *httputil.ReverseProxy
}

type LoadBalancer struct {
	backends []*Backend
	current  uint32
}

func (lb *LoadBalancer) NextBackend() *Backend {
	index := atomic.AddUint32(&lb.current, 1) % uint32(len(lb.backends))
	return lb.backends[index]
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.NextBackend()
	backend.Proxy.ServeHTTP(w, r)
}

func main() {
	backends := []*Backend{
		{URL: mustParseURL("http://localhost:8081"), Proxy: httputil.NewSingleHostReverseProxy(mustParseURL("http://localhost:8081"))},
		{URL: mustParseURL("http://localhost:8082"), Proxy: httputil.NewSingleHostReverseProxy(mustParseURL("http://localhost:8082"))},
	}

	lb := &LoadBalancer{backends: backends}
	log.Fatal(http.ListenAndServe(":8080", lb))
}

func mustParseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	return u
}
```

**Пример работы**:

- Запросы к localhost:8080 распределяются между 8081, 8082.

### Преимущества

- Масштабируемость.
- Отказоустойчивость.
- Простая интеграция.

### Недостатки

- Ограниченная логика.
- Требует health checks.
- Точка отказа.

### Когда использовать

- Высоконагруженные системы.
- Кластеры серверов.
- Kubernetes/Docker.

---

## Сравнение технологий

|                   |                       |                                           |           |                                      |
| ----------------- | --------------------- | ----------------------------------------- | --------- | ------------------------------------ |
| Технология        | Основная роль         | Функции                                   | Сложность | Применение                           |
| **Forward Proxy** | Скрытие клиента       | Анонимность, кэширование, фильтрация      | Низкая    | Корпоративные сети, обход блокировок |
| **Reverse Proxy** | Скрытие серверов      | Маршрутизация, кэширование, SSL           | Низкая    | Веб-приложения, статический контент  |
| **API Gateway**   | Управление API        | Аутентификация, лимитирование, мониторинг | Средняя   | Микросервисы, API для клиентов       |
| **Load Balancer** | Балансировка нагрузки | Распределение трафика, health checks      | Средняя   | Высоконагруженные системы, кластеры  |

### Из одной серии?

Да, все технологии управляют трафиком, но:

- **Forward Proxy** работает от имени клиента, скрывая его.
- **Reverse Proxy** работает от имени сервера, скрывая бэкенд.
- **API Gateway** расширяет Reverse Proxy для API.
- **Load Balancer** фокусируется на распределении нагрузки.

### Ключевые различия

- **Forward Proxy**:
  - Клиент → Proxy → Интернет.
  - Скрывает клиента (анонимность, фильтрация).
- **Reverse Proxy**:
  - Клиент → Proxy → Бэкенд.
  - Скрывает серверы, маршрутизирует.
- **API Gateway**:
  - Клиент → Gateway → Микросервисы.
  - Управляет API, добавляет безопасность.
- **Load Balancer**:
  - Клиент → Balancer → Серверы.
  - Распределяет нагрузку.

### Примеры сценариев

- **Forward Proxy**:
  - Squid для корпоративной фильтрации.
  - Shadowsocks для обхода блокировок.
- **Reverse Proxy**:
  - Nginx для маршрутизации /static и /api.
  - Caddy для HTTPS.
- **API Gateway**:
  - Kong для JWT-аутентификации.
  - AWS API Gateway для Lambda.
- **Load Balancer**:
  - HAProxy для кластера API.
  - AWS ELB для Kubernetes.

### Рекомендации

- **Forward Proxy**: Для анонимности, фильтрации или обхода ограничений. Например, Squid в офисе.
- **Reverse Proxy**: Для веб-приложений с простой маршрутизацией. Например, Nginx для статики.
- **API Gateway**: Для микросервисов с API. Например, Kong для аутентификации.
- **Load Balancer**: Для высоконагруженных систем. Например, HAProxy для кластера.

### Примечания

- **Комбинирование**: Forward Proxy может использоваться клиентами, Reverse Proxy — перед API Gateway, Load Balancer — перед серверами.
- **Код**: Примеры минималистичны. В продакшене используйте Nginx, Kong, HAProxy с конфигурациями.
- **Фронтенд**: Компоненты обслуживают фронтенд (например, React), но акцент на серверной части.
