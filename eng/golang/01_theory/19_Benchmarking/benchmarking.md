# ⏱️ Benchmarking and Profiling in Go

**Benchmarking** is the process of measuring code performance. In Go, it is a built-in feature of the `testing` package that allows you to find out how much time an operation takes and how much memory it consumes.

---


# 1. 📝 Benchmarking Basics

Benchmarks are written in `_test.go` files and have the signature `func BenchmarkXxx(b *testing.B)`.


### Example of a simple benchmark:

```go
func BenchmarkMyFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MyFunction()
    }
}
```

- **`b.N`**: This is a magic number that Go picks automatically. It runs the loop as many times as needed to get a statistically significant result (usually it takes about 1 second).


### How to run:
```bash
go test -bench=.          # Run all benchmarks
go test -bench=MyFunc     # Run a specific benchmark
go test -bench=. -benchmem # Show memory allocations
```

---


# 2. 🛠️ Time Management Tools

Sometimes preparation for a test takes more time than the test itself. To get a fair result, use methods of `b`:

- **`b.ResetTimer()`**: Resets the accumulated time. Call after heavy preparation.
- **`b.StopTimer()` / `b.StartTimer()`**: Allow excluding specific code sections from measurements.
- **`b.ReportAllocs()`**: Enables allocation reports directly in the code (equivalent to the `-benchmem` flag).

```go
func BenchmarkComplex(b *testing.B) {
    data := prepareHugeData() // Heavy preparation
    b.ResetTimer()            // Ignore preparation time

    for i := 0; i < b.N; i++ {
        Process(data)
    }
}
```

---


# 3. ⚠️ Best Practices (How not to deceive yourself)


### 1. Avoid Compiler Optimizations 🧠
If the result of a function is not used anywhere, the compiler may simply "discard" the function call, and the benchmark will show 0 ns.

**Bad:**
```go
func BenchmarkBad(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello %d", i) // The compiler can optimize this away
    }
}
```

**Good (use a global variable):**
```go
var result string

func BenchmarkGood(b *testing.B) {
    var r string
    for i := 0; i < b.N; i++ {
        r = fmt.Sprintf("hello %d", i)
    }
    result = r // Write to a global variable so the compiler doesn't delete the code
}
```


### 2. Parallel Benchmarks 🚀
If you want to check how your code behaves under load from multiple goroutines:

```go
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            MyConcurrentFunc()
        }
    })
}
```

---


# 4. 🔍 Profiling (pprof)

If a benchmark shows that your code is slow, you need to understand **why**. Go has `pprof` for this.


### Profile Types:
1. **CPU Profile**: Where the processor spends the most time.
2. **Heap (Memory) Profile**: Who allocates memory and where it lives.
3. **Alloc Profile**: All memory allocations (even if already cleaned up by GC).
4. **Goroutine Profile**: State of all goroutines.
5. **Mutex/Block Profile**: Where the program stops while waiting for locks.


### How to capture a profile from a benchmark:
```bash
go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof
```


### How to view results:
```bash
# Web interface (the most convenient way)
go tool pprof -http=:8080 cpu.prof

# Interactive console mode
go tool pprof cpu.prof
(pprof) top10      # Top 10 functions by time
(pprof) list Func  # Detailed breakdown of a function line by line
```

---


# 5. 📽️ Execution Trace (`runtime/trace`)

While `pprof` shows aggregated data, `trace` shows a **timeline**: how goroutines worked, when GC was triggered, and which cores executed the code.


### Running:
```bash
go test -bench=. -trace=trace.out
go tool trace trace.out
```

---


# 6. 🎓 Advanced Tips

- **Inlining**: Check if your function is too large for inlining (`go build -gcflags="-m"`).
- **Escape Analysis**: Understand if a variable goes to the heap. Stack allocations are practically free.
- **Benches vs Tests**: Benchmarks are NOT tests. They should not check for correctness, only for speed.
- **Environment Stability**: Do not run benchmarks on a laptop that is currently updating its system or connected to different power sources (CPU frequency may jump).

---


### 📊 What to check first?
| Problem | Tool | Flag / Command |
| :--- | :--- | :--- |
| Overall speed | `testing.B` | `go test -bench=.` |
| Leaks / Allocations | Mem Profile | `go test -benchmem` |
| Blockage idle | Mutex Profile | `-mutexprofile=mutex.out` |
| Unclear behavior | Trace | `-trace=trace.out` |

<!-- QUIZ_START 

[
    {
        "question": "What is the correct function signature for a benchmark in Go?",
        "options": [
            "func TestXxx(t *testing.T)",
            "func BenchmarkXxx(b *testing.B)",
            "func RunBenchmark(n int)",
            "func BenchmarkXxx(m *testing.M)"
        ],
        "correctIndex": 1
    },
    {
        "question": "What is 'b.N' in a Go benchmark function?",
        "options": [
            "A constant representing 100",
            "A dynamic number of iterations determined by the Go test runner to get a reliable result",
            "The number of CPU cores",
            "The memory limit in bytes"
        ],
        "correctIndex": 1
    },
    {
        "question": "Why should you use 'b.ResetTimer()' in a benchmark?",
        "options": [
            "To speed up the code",
            "To exclude heavy initialization/setup code from the measured benchmark time",
            "To clean up memory",
            "To stop the benchmark immediately"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

