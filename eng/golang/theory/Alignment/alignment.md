# 📏 Memory Alignment and Padding in Go

## 1. 🧩 What is Alignment

**Alignment** is a rule for placing data in memory, according to which a value of a certain type must start at an address that is a multiple of its **alignment**.

**Example:**
- `int32` must start at an address that is a multiple of **4**.
- `int64` — a multiple of **8**.

> [!NOTE]
> Alignment is directly related to processor architecture and affects the correctness and performance of the program.

## 2. ❓ Why is Alignment Needed

### 2.1 Processor Architecture
Processors read memory not byte by byte, but in blocks of fixed size (words). If data is aligned, **one memory read** is required. If not, the processor has to perform multiple reads and "glue" the results together, which slows down operation.

### 2.2 Atomic Operations
For many processors, atomic operations (the `sync/atomic` package) work correctly **only** if the data is aligned to its size. On 32-bit architectures, unaligned `int64` will trigger a panic during an atomic access attempt.

## 3. 🧱 What is Padding

**Padding** are empty bytes that the compiler automatically inserts between fields or at the end of a struct to satisfy alignment requirements.

## 4. 🔢 Base Type Alignment (64-bit)

| Type | Size (bytes) | Alignment |
| :--- | :---: | :---: |
| `bool`, `int8`, `uint8` | 1 | 1 |
| `int16`, `uint16` | 2 | 2 |
| `int32`, `uint32`, `float32` | 4 | 4 |
| `int64`, `uint64`, `float64`, `pointer` | 8 | 8 |
| `string` (struct with 2 fields) | 16 | 8 |
| `slice` (struct with 3 fields) | 24 | 8 |

> [!TIP]
> The size of a type and its alignment are different things. For example, a `string` occupies 16 bytes, but its **alignment** is 8, because the largest field inside (a pointer) requires 8-byte alignment.

## 5. 🏢 Struct Alignment

### 5.1 Basic Rule
Each field in a struct must start at an address that is a multiple of its own **alignment**. The total size of the struct must also be a multiple of its overall **alignment** (which is the maximum among all fields).

### 5.2 Example: Suboptimal Order
```go
type Bad struct {
    a bool  // 1 byte
            // [7 bytes padding]
    b int64 // 8 bytes
} // Total: 16 bytes
```
**Memory:** `[a][p][p][p][p][p][p][p] [b][b][b][b][b][b][b][b]`

### 5.3 Example: Optimal Order
```go
type Good struct {
    b int64 // 8 bytes
    a bool  // 1 byte
            // [7 bytes padding at the end]
} // Total: 16 bytes
```
**Memory:** `[b][b][b][b][b][b][b][b] [a][p][p][p][p][p][p][p]`

> [!TIP]
> **Why are both examples 16 bytes?**
> The total size of a struct must always be a multiple of its maximum alignment (in this case, **8** for `int64`).
> In both cases, `8 + 1 = 9`, and the nearest multiple of 8 is **16**.
>
> **Why is `Good` better?**
> In `Good`, the empty bytes (padding) are collected at the end. If you add another field `c bool`, it will take the place of the existing padding:
> - `Good` + `c bool` = **16 bytes**
> - `Bad` + `c bool` (at the end) = **24 bytes**

> [!IMPORTANT]
> **Rule of Thumb**: Arrange fields in descending order of their size. This minimizes the amount of **padding** inserted.

## 6. 🚀 Advanced Topics

### 6.1 Zero-sized types
The types `struct{}` and `[0]int` have a size of **0**. However, they have a peculiarity when placed in structs:

```go
type S1 struct {
    a struct{}
    b int64
} // Size: 8, 'a' takes no space

type S2 struct {
    a int64
    b struct{}
} // Size: 16! (Padding added at the end)
```

> [!WARNING]
> If an empty field is the **last** one in a struct, the compiler adds **padding** so that the address of this field does not point to the next object in memory (this is important for correct GC operation).

### 6.2 Cache Line and False Sharing
Memory is read from RAM into processor cache in blocks of **64 bytes** (**Cache Line**). 
- If two goroutines frequently write to different fields of the same struct that happen to be in the same cache line, the processor will constantly invalidate the cache. This is called **False Sharing**.
- Solution: add "empty" fields (padding) to move active data into different cache lines.

### 6.3 Checking via `unsafe`
```go
import "unsafe"

fmt.Println(unsafe.Sizeof(s))   // Total size
fmt.Println(unsafe.Alignof(s))  // Struct alignment
fmt.Println(unsafe.Offsetof(s.b)) // Offset of field 'b' from the beginning
```

---

## 🎯 Summary: Practical Tips
1. **Group fields**: from largest to smallest (`int64` -> `int32` -> `bool`).
2. **Mind the tail**: do not put `struct{}` as the last field if minimal size is important.
3. **Atomics**: ensure variables for `sync/atomic` are 8-byte aligned.
4. **Tools**: use linters (e.g., `fieldalignment`) to automatically find unnecessary padding.