# 🔗 Integration Tests

Integration tests verify how different parts of your application work with each other. Unlike unit tests, we don't mock everything here; instead, we try to use real dependencies (databases, cache, other services).


# 1. 🎯 What Do We Test in Integration Tests?
*   Database interaction (whether data is correctly saved and read).
*   HTTP API (correctness of routing, status codes, and JSON responses).
*   External integrations (message queues, gRPC calls).


# 2. 🗄️ Database Testing

The main difficulty here is the database state. Each test must start in a clean environment.


### Approach with Test Containers
[Testcontainers-go](https://golang.testcontainers.org/) is a library that allows running a real database in Docker directly from the test code.

```go
func TestUserRepository_Save(t *testing.T) {
    ctx := context.Background()
    // Run Postgres in a container
    container, err := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:15-alpine"))
    if err != nil {
        t.Fatal(err)
    }
    defer container.Terminate(ctx)

    dsn, _ := container.ConnectionString(ctx)
    db, _ := sql.Open("postgres", dsn)
    
    // ... test execution ...
}
```


# 3. 🌐 HTTP API Testing

In Go, there is an excellent package `net/http/httptest` for this. it allows testing handlers without running a real HTTP server.

```go
func TestUserHandler_GetUser(t *testing.T) {
    // 1. Request preparation
    req := httptest.NewRequest("GET", "/users/1", nil)
    
    // 2. Creating ResponseRecorder (simulating response writing)
    rr := httptest.NewRecorder()
    
    // 3. Calling the handler
    handler := http.HandlerFunc(GetUserHandler)
    handler.ServeHTTP(rr, req)

    // 4. Verifying the result
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `{"id":1, "name":"Gopher"}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
```


# 4. 💡 Tips for Writing

1.  **Data Isolation**: If you use one DB for all tests, wrap each test in a transaction and roll it back (`ROLLBACK`) at the end of the test.
2.  **Build Tags**: Integration tests are slow. To prevent them from running during a regular `go test ./...`, use tags:
    ```go
    // +build integration

    package tests
    ```
    Running: `go test -tags=integration ./...`
3.  **Docker Compose**: For complex environments, you can use `docker-compose` to bring up the entire infrastructure before running tests.

> [!WARNING]
> Never run integration tests against a "live" (production) or even "staging" database. Use only local or specifically designated test instances.

<!-- QUIZ_START 

[
    {
        "question": "What is the key difference between Unit tests and Integration tests?",
        "options": [
            "Integration tests are faster",
            "Unit tests test the UI, Integration tests test the database",
            "Integration tests verify how different parts of the application work together using real dependencies instead of mocks",
            "Unit tests only work on Windows"
        ],
        "correctIndex": 2
    },
    {
        "question": "Which Go standard library package is used to test HTTP handlers without running a real server?",
        "options": [
            "net/http",
            "net/http/httptest",
            "testing",
            "os/exec"
        ],
        "correctIndex": 1
    },
    {
        "question": "What is 'Testcontainers-go' used for in integration testing?",
        "options": [
            "To write unit tests",
            "To run real databases or services in Docker directly from test code",
            "To deploy applications to production",
            "To measure code coverage"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

