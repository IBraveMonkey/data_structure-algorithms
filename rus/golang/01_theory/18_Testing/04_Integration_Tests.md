# 🔗 Интеграционные тесты (Integration Tests)

Интеграционные тесты проверяют, как различные части вашего приложения работают друг с другом. В отличие от unit-тестов, здесь мы не мокаем все подряд, а стараемся использовать реальные зависимости (базы данных, кэш, другие сервисы).


# 1. 🎯 Что тестируем в интеграционных тестах?
*   Взаимодействие с базой данных (правильно ли сохраняются и читаются данные).
*   HTTP API (корректность роутинга, статус-кодов и JSON-ответов).
*   Внешние интеграции (очереди сообщений, gRPC-вызовы).


# 2. 🗄️ Тестирование Базы Данных

Основная сложность здесь — состояние БД. Каждый тест должен начинаться в чистом окружении.


### Подход с Test Containers
[Testcontainers-go](https://golang.testcontainers.org/) — это библиотека, которая позволяет запускать реальную БД в Docker прямо из кода теста.

```go
func TestUserRepository_Save(t *testing.T) {
    ctx := context.Background()
    // Запуск Postgres в контейнере
    container, err := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:15-alpine"))
    if err != nil {
        t.Fatal(err)
    }
    defer container.Terminate(ctx)

    dsn, _ := container.ConnectionString(ctx)
    db, _ := sql.Open("postgres", dsn)
    
    // ... выполнение теста ...
}
```


# 3. 🌐 Тестирование HTTP API

В Go для этого есть отличный пакет `net/http/httptest`. Он позволяет тестировать хендлеры без запуска реального HTTP-сервера.

```go
func TestUserHandler_GetUser(t *testing.T) {
    // 1. Подготовка запроса
    req := httptest.NewRequest("GET", "/users/1", nil)
    
    // 2. Создание ResponseRecorder (имитация записи ответа)
    rr := httptest.NewRecorder()
    
    // 3. Вызов хендлера
    handler := http.HandlerFunc(GetUserHandler)
    handler.ServeHTTP(rr, req)

    // 4. Проверка результата
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `{"id":1, "name":"Gopher"}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
```


# 4. 💡 Советы по написанию

1.  **Изоляция данных**: Если вы используете одну БД для всех тестов, оборачивайте каждый тест в транзакцию и откатывайте ее (`ROLLBACK`) в конце теста.
2.  **Теги сборки (Build Tags)**: Интеграционные тесты медленные. Чтобы они не запускались при обычном `go test ./...`, используйте теги:
    ```go
    // +build integration

    package tests
    ```
    Запуск: `go test -tags=integration ./...`
3.  **Docker Compose**: Для сложных окружений можно использовать `docker-compose` для поднятия всей инфраструктуры перед запуском тестов.

> [!WARNING]
> Никогда не запускайте интеграционные тесты против "живой" (production) или даже "staging" базы данных. Используйте только локальные или специально выделенные тестовые инстансы.

<!-- QUIZ_START 

[
    {
        "question": "Чем интеграционные тесты отличаются от модульных (unit) тестов?",
        "options": [
            "Они работают быстрее",
            "Они не используют реальные зависимости",
            "Они проверяют взаимодействие нескольких компонентов и часто используют реальные базы данных или API",
            "Они пишутся на другом языке"
        ],
        "correctIndex": 2
    },
    {
        "question": "Какой пакет в Go предназначен для тестирования HTTP-хендлеров без запуска реального сервера?",
        "options": [
            "net/http",
            "net/http/httptest",
            "encoding/json",
            "testing/http"
        ],
        "correctIndex": 1
    },
    {
        "question": "Что такое 'Test Containers'?",
        "options": [
            "Коробки для хранения жестких дисков",
            "Библиотека для программного запуска реальных сервисов (например, Postgres) в Docker прямо из кода тестов",
            "Способ упаковки Go-приложения",
            "Тип данных в Go"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

