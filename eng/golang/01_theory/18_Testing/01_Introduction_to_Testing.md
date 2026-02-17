# 🧪 Introduction to Testing

Testing is an essential part of the software development process. In Go, testing is built into the language itself from day one, emphasizing its importance for building reliable systems.


# 1. ❓ Why Do We Need Tests?

Many beginners ask: "Why spend time writing tests if I can already see that the program works?". Here are the main reasons:

1.  **Preventing Regressions**: When you make changes to your code, tests ensure that old functionality hasn't broken.
2.  **Confidence During Refactoring**: You can boldly change the structure of your code, knowing that if the logic is violated, the tests will show it.
3.  **Documentation**: Well-written tests serve as living examples of how to use your code.
4.  **Improving Code Design**: Code that is easy to test is usually better designed (fewer coupled components, clearer interfaces).
5.  **Long-term Time Savings**: Finding a bug in production is dozens of times more expensive and time-consuming than discovering it during the development stage.


# 2. 🔍 Overview of Test Types

In Go, it's common to divide tests into several levels (we will explore them in more detail in the following chapters):

*   **Unit Tests**: Verify minimum parts of code (functions, methods) in isolation. They are very fast and cheap to write.
*   **Integration Tests**: Verify how several system components work together (e.g., interaction between a service and a database).
*   **E2E (End-to-End) Tests**: Verify the entire system operation, simulating real user actions.


# 3. 🛡️ Testing Philosophy in Go

Unlike many other languages, Go doesn't require third-party libraries to start testing. The `testing` package provides everything you need "out of the box".

Core rules:
*   Test files always end with `_test.go`.
*   Test functions must start with `Test` and accept `*testing.T`.
*   Tests reside in the same package as the tested code (or in a `xxx_test` package for external testing).

> [!TIP]
> In Go, it's standard practice to write tests as simply and clearly as the main code. Avoid "magic" and complex frameworks where possible.

<!-- QUIZ_START 

[
    {
        "question": "What is the primary reason for 'refactoring confidence' provided by tests?",
        "options": [
            "Tests make the code run faster",
            "Tests allow you to change code structure knowing that logic violations will be detected",
            "Tests automatically rewrite broken code",
            "Tests prevent other developers from editing the file"
        ],
        "correctIndex": 1
    },
    {
        "question": "According to the 'Long-term Time Savings' principle, when is it cheapest to find a bug?",
        "options": [
            "In production",
            "During the development stage",
            "After the project is closed",
            "During a security audit"
        ],
        "correctIndex": 1
    },
    {
        "question": "What suffix must Go test files have?",
        "options": [
            ".test.go",
            "_test.go",
            "_spec.go",
            ".go_test"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

