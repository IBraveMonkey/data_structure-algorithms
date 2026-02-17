# 🏃 Running Tests

The `go test` command is a powerful tool with many flags that allow you to control the testing process.


# 1. 💻 Basic Commands

*   `go test .`: Running tests in the current directory.
*   `go test ./...`: Running tests in all subdirectories of the project.
*   `go test -v ./...`: Verbose output. Shows the names of all tests and the output from `t.Log`.


# 2. 🎯 Running Specific Tests

Sometimes you need to run only one test or a group of tests to avoid waiting for others.

*   `go test -run TestName`: Runs a test with a specific name.
*   `go test -run TestUser/`: Runs all tests whose name starts with `TestUser`.
*   `go test -run /ValidData`: Runs all subtests matching a regular expression.


# 3. ⚡ Parallel Execution

Go can run tests in parallel, which significantly speeds up verification of large projects.

1.  In the test code, you need to call `t.Parallel()`:
    ```go
    func TestMyFunction(t *testing.T) {
        t.Parallel()
        // ...
    }
    ```
2.  The `-parallel N` flag limits the number of simultaneously running tests (defaults to `GOMAXPROCS`).


# 4. 🚩 Useful Flags

*   `-count N`: Run each test N times. Useful for finding "flaky" bugs (the `-count 1` flag also disables test caching).
*   `-timeout 30s`: Set a time limit. If tests take longer, they will be forcibly stopped.
*   `-short`: Tells tests to skip long scenarios. Verified in code via `testing.Short()`.
*   `-failfast`: Stop execution at the first failed test.


# 5. 🏷️ Build Tags

Allow separating tests into groups (e.g., unit and integration).

Add a line to the beginning of the file:
```go
// +build integration
```

Run with the flag:
```bash
go test -tags=integration ./...
```


# 6. 💾 Test Caching

Go caches test results. If the code hasn't changed, you'll see `(cached)` in the output.
To **clear the cache**, use:
```bash
go clean -testcache
```

> [!TIP]
> Use `go test -v -race ./...` during development. The `-race` flag enables the race detector, which helps find complex bugs in concurrent code.

<!-- QUIZ_START 

[
    {
        "question": "Which command is used to run all tests in all subdirectories of a project?",
        "options": [
            "go test .",
            "go test -all",
            "go test ./...",
            "go run test"
        ],
        "correctIndex": 2
    },
    {
        "question": "What does the '-v' flag do when running tests?",
        "options": [
            "Verifies the code version",
            "Ensures 'verbose' output, showing all test names and logs",
            "Runs tests very fast",
            "Validates the compiler version"
        ],
        "correctIndex": 1
    },
    {
        "question": "How can you run only a specific test function named 'TestUserLogin'?",
        "options": [
            "go test -only TestUserLogin",
            "go test -run TestUserLogin",
            "go test -name TestUserLogin",
            "go run TestUserLogin"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

