# Defer in Go: A Complete Guide ⏱️

`defer` is a keyword in Go that post-pones the execution of a function until the surrounding function finishes (either via `return` or due to a panic).

---

## 1. How it Works ⚙️

### LIFO Queue (Last-In-First-Out) 📚
Deferred functions are called in the reverse order of their declaration. It resembles a stack: the last added function will be executed first. This is extremely convenient for step-by-step resource cleanup.

```go
func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
    // Output: Third, Second, First
}
```

### Argument Evaluation 📝
The arguments of a `defer` function are evaluated **at the time of declaration**, not at the time of execution.

```go
func main() {
    i := 0
    defer fmt.Println(i) // i will be 0, as the value is copied here
    i++
    return
}
```

---

## 2. Defer and Named Return Values 🔄

`defer` can modify the named return values of a function. This happens because `defer` is executed **after** the `return` statement has set the values, but **before** the actual exit from the function.

```go
func result() (x int) {
    defer func() {
        x++ // Modifying the return value
    }()
    return 5 // x becomes 5, then defer makes x = 6
}
```

---

## 3. Panic and Recover 🛡️

`defer` is often used together with `recover()` to intercept panics and prevent the entire program from crashing.

```go
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("something went wrong")
}
```

### Important Panic/Recover Restrictions: ⚠️
1. **Scope**: You cannot catch a panic from another goroutine. Each goroutine must have its own `defer/recover`.
2. **Types of errors that CANNOT be caught**:
    - **Out of Memory (OOM)**: If the system runs out of memory, the Go runtime simply terminates the process. `recover` won't help here.
    - **Stack Overflow**: A stack overflow also cannot be handled via `defer`.
    - **Fatal errors** in the runtime (for example, memory corruption).

### What CAN be caught: ✅
- **Runtime panics**: division by zero, nil pointer dereference, slice index out of range. These errors initiate the standard panic mechanism, which `defer` can intercept.

---

## 4. Why Use `defer`? 🎯

1. **Resource Cleanup**: Closing files (`f.Close()`), network connections, or descriptors.
2. **Mutex Unlocking**: `defer mu.Unlock()` guarantees that a mutex will be released even if a panic occurs.
3. **Logging**: Measuring function execution time or logging the exit result.
4. **Panic Interception**: Preventing a service crash when unforeseen errors occur.

---

## 5. Under the Hood and Performance 🚀

- **Stack**: `defer` objects are stored in a linked list within the goroutine structure (`g`). This requires heap allocation (although Go 1.13+ introduced "open-coded defers" optimizations, making calls almost free for simple cases).
- **Arguments**: All arguments are copied at the time the `defer` is called. It's important to remember this when working with large structures (it's better to pass pointers).

> [!IMPORTANT]
> Do not use `defer` inside loops with a large number of iterations unless justified. All deferred calls will accumulate until the function exits, which can lead to a temporary spike in memory consumption.

---

## Interview Summary

| Error | Can be caught via `defer/recover`? |
| :--- | :--- |
| Division by 0 | ✅ Yes |
| Nil pointer dereference | ✅ Yes |
| Index out of range | ✅ Yes |
| Stack Overflow | ❌ No |
| Out of Memory (OOM) | ❌ No |
| Panic in another goroutine | ❌ No |
