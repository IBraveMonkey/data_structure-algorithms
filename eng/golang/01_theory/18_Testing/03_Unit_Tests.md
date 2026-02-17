# 🧩 Unit Tests

Unit testing is the foundation of your testing pyramid. In Go, it is implemented as simply and effectively as possible through the standard library.


# 1. 📦 `testing` Package Basics

Any test in Go is a regular function in a `*_test.go` file.

Example of a simple test:
```go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```


### Main `testing.T` Methods:
*   `t.Error(args...)`: Reports an error but continues test execution.
*   `t.Errorf(format, args...)`: Formatted error message.
*   `t.Fatal(args...)`: Reports an error and immediately stops execution of the current test.
*   `t.Log(args...)`: Equivalent to `fmt.Println` for outputting information only when a test fails or the `-v` flag is used.


# 2. 📊 Table-Driven Tests

This is the de facto standard in the Go community. Instead of writing 10 functions for different inputs, you describe a structure with data and iterate through it in a loop.

```go
func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -1, -2},
        {"zeroes", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```
> [!NOTE]
> The `t.Run` method creates a "subtest", which allows running them individually and seeing the result for each case.


# 3. 🛠️ Test Fixtures and Helpers

If you need to create a complex object or prepare the environment, use helper functions.

```go
func setupTestDB(t *testing.T) *Database {
    t.Helper() // Indicates that if an error occurs, it should show the line where this function was called, not the line inside the helper
    db := NewDatabase()
    // data preparation...
    return db
}
```


# 4. 🎭 Mocking and Test Doubles

In Go, it's not common to use heavy mocking frameworks if you can get by with interfaces.

**Step 1: Describe the interface**
```go
type OrderRepository interface {
    Save(order *Order) error
}
```

**Step 2: Create a mock structure in the test**
```go
type mockRepo struct {
    saveCalled bool
}

func (m *mockRepo) Save(o *Order) error {
    m.saveCalled = true
    return nil
}

func TestOrderService_Create(t *testing.T) {
    repo := &mockRepo{}
    service := NewOrderService(repo)
    // ...
}
```


### Popular Libraries:
If the project is large and there are many mocks, you can use:
*   [stretchr/testify](https://github.com/stretchr/testify): The most popular set of utilities for assertions and mocks.
*   [golang/mock (gomock)](https://github.com/golang/mock): Mock generation based on interfaces.

> [!TIP]
> Try to use interfaces where you expect complex logic or external dependencies. This will make your code "testable".

<!-- QUIZ_START 

[
    {
        "question": "Which method on 'testing.T' should you use to report an error but allow the test to continue executing?",
        "options": [
            "t.Fatal()",
            "t.Error()",
            "t.Log()",
            "t.Skip()"
        ],
        "correctIndex": 1
    },
    {
        "question": "What is the 'Table-Driven Test' pattern in Go?",
        "options": [
            "A test that uses a physical table",
            "A pattern where you iterate through a slice of structs containing input and expected output to run multiple test cases",
            "A test that only works with databases",
            "A way to sort test results"
        ],
        "correctIndex": 1
    },
    {
        "question": "What does the 't.Helper()' method do in a test helper function?",
        "options": [
            "It calls for help from the OS",
            "It tells the test runner to skip this function",
            "It marks the function as a helper so that error reports show the caller's line number instead of the helper's line number",
            "It automatically mocks all interfaces"
        ],
        "correctIndex": 2
    }
]

QUIZ_END -->

