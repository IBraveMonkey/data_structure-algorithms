# ✍️ Writing Tests (Best Practices)

Writing tests is an art. Good tests help development; bad ones only get in the way, becoming a burden with every code change.


# 1. 🏷️ Naming Conventions

In Go, strict rules are adopted:
1.  **File Name**: Always `file_name_test.go`.
2.  **Function Name**: Always `TestXxx`, where `Xxx` starts with a capital letter.
    *   Recommended: `TestMethodName_Scenario` (e.g., `TestUser_Save_ValidData`).
3.  **Package**:
    *   `package user`: If the test should have access to unexported (private) fields.
    *   `package user_test`: If you are testing only the public API (black-box testing). This is the preferred option.


# 2. 🧱 AAA Pattern (Arrange, Act, Assert)

This is the standard structure for any test, which makes it readable.

```go
func TestCalculateBonus(t *testing.T) {
    // 1. Arrange
    // Preparing input data and environment
    salary := 1000.0
    performance := 1.5
    expectedBonus := 1500.0

    // 2. Act
    // Calling the function under test
    result := CalculateBonus(salary, performance)

    // 3. Assert
    // Verifying the result
    if result != expectedBonus {
        t.Errorf("got %f, want %f", result, expectedBonus)
    }
}
```


# 3. 📂 Test Organization


### Setup and Teardown
Sometimes you need to run code before or after all tests in a package (e.g., to set up a database). the `TestMain` function is used for this.

```go
func TestMain(m *testing.M) {
    setup() // preparation
    code := m.Run() // running all tests
    teardown() // cleanup
    os.Exit(code)
}
```

For individual tests, use `t.Cleanup`:
```go
func TestWithFile(t *testing.T) {
    f := createTempFile()
    t.Cleanup(func() {
        os.Remove(f.Name()) // will execute at the end of the test
    })
}
```


# 4. ✨ Best Practices

1.  **Test Behavior, Not Implementation**: Don't check how a function works internally; check WHAT it returns as output.
2.  **Minimum Logic in Tests**: There should be no `for` loops in tests (except for table-driven tests) and no complex `if/else` conditions. A test should be "flat" and understandable at first glance.
3.  **Isolation**: One test should not depend on the result of another.
4.  **Determinism**: A test should always return the same result with the same input data. Avoid using `time.Now()` directly (use time injection or mocks).

> [!IMPORTANT]
> The best test is the one you wrote BEFORE you started writing the code itself (TDD — Test Driven Development). This forces you to think about the interface and requirements first.

<!-- QUIZ_START 

[
    {
        "question": "What is the standard naming convention for test files in Go?",
        "options": [
            "test_file.go",
            "file_test.go",
            "file.test",
            "fileTest.go"
        ],
        "correctIndex": 1
    },
    {
        "question": "Which package name is generally preferred for testing the public API of a package 'user' (black-box testing)?",
        "options": [
            "package user",
            "package testing",
            "package user_test",
            "package main"
        ],
        "correctIndex": 2
    },
    {
        "question": "What does the AAA pattern in testing stand for?",
        "options": [
            "Always Add Assertions",
            "Arrange, Act, Assert",
            "Apply, Assess, Adjust",
            "Analyze, Action, Audit"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

