# 🏁 E2E Tests (End-to-End Tests)

End-to-End tests are the tip of our pyramid. They verify the entire system, from the interface (or API) to the deepest level of the database, simulating the behavior of a real user.


# 1. 🎯 Purpose of E2E Testing

E2E tests answer the main question: **"Does the entire business scenario work?"**.

Example chain:
1.  User registers.
2.  Receives a confirmation email.
3.  Logs into the system.
4.  Creates an order.

If a problem occurs at any of the stages (database, email service, frontend, backend), an E2E test will detect it.


# 2. 📅 When to Use E2E Tests?

Since they are slow and expensive, they should not cover everything.

**Good candidates for E2E:**
*   **Critical Path**: The most important scenarios (payment, registration).
*   **Complex Integrations**: Where more than 3-4 microservices are involved.
*   **Smoke tests**: A quick check after deployment that the main functions are "alive".


# 3. 🛠️ Tools in Go

Although Go is often used for server-side logic, there are tools for full testing:

1.  **Playwright / Selenium / Cypress**: Usually used for Web interface testing. There are bindings for Go, for example [playwright-go](https://github.com/playwright-community/playwright-go).
2.  **Godog (Cucumber)**: Allows writing tests in the Gherkin language (Given/When/Then), which is convenient for aligning business logic.


### Smoke Test Example in Go

For an API service, an E2E test might look like starting the entire application and making real HTTP requests to it.

```go
func TestEndToEnd_ComplexScenario(t *testing.T) {
    // 1. Set up infrastructure (Docker Compose or TestContainers)
    // 2. Start the application itself
    app := StartApp()
    defer app.Stop()

    client := &http.Client{}

    // Step 1: Registration
    resp, _ := client.Post("http://localhost:8080/register", "application/json", body)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)

    // Step 2: Checking state in the database (via API)
    resp, _ = client.Get("http://localhost:8080/profile")
    // ...
}
```


# 4. ⚠️ Problems with E2E Tests

*   **Brittleness**: Changing text on a button or a network delay of 1 second can "drop" a test.
*   **Setup Complexity**: You need to prepare the state of all dependencies (databases, cache, external API simulation).
*   **Execution Time**: They can take minutes or even hours.

> [!TIP]
> Try to keep the number of E2E tests to a minimum. If you can verify logic at the Unit or Integration test level, do it there. E2E is the "last line of defense".

<!-- QUIZ_START 

[
    {
        "question": "What is the main question that End-to-End (E2E) tests aim to answer?",
        "options": [
            "Is the function well-optimized?",
            "Does the entire business scenario work (e.g., registration to order completion)?",
            "Does the code follow style guides?",
            "How many goroutines are running?"
        ],
        "correctIndex": 1
    },
    {
        "question": "What are good candidates for E2E testing?",
        "options": [
            "Formatting of log messages",
            "Critical paths (e.g., payment, registration) and complex multi-service integrations",
            "Simple math functions",
            "Private methods of a struct"
        ],
        "correctIndex": 1
    },
    {
        "question": "Which of these is a common problem with E2E tests?",
        "options": [
            "They are too simple to write",
            "Brittleness (small UI or network changes can cause failures) and long execution time",
            "They only test the database",
            "They are too fast for developers to see"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

