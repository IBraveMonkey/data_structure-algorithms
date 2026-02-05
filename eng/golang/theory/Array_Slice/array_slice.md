### 📦 Arrays and Slices in Go

Arrays and slices are fundamental data structures in Go for working with sequences of elements.

---

# 1. 🏗️ Array

An **Array** is a numbered collection of elements of a **fixed** length of the same type.

### ⚙️ Key Features
- **Fixed Size**: The size of an array is determined when it is created and is part of its type (for example, `[3]int` and `[5]int` are different types).
- **Pass by Value**: When passing an array to a function or assigning it, it is **copied in its entirety**.
- **Memory**: Elements are stored in memory as a continuous block.

```go
arr := [3]int{1, 2, 3}
var small [2]string = [2]string{"hi", "go"}
matrix := [2][2]int{{1, 2}, {3, 4}}
```

> [!TIP]
> Use `[...]int{1, 2, 3}` if you want the compiler to determine the array size automatically based on the number of elements.

---

# 2. 🧩 Slice

A **Slice** is a powerful and flexible wrapper over an array. Unlike arrays, slices have a dynamic length.

### ⚙️ Internal Structure (Slice Header)
Internally, a slice is represented by a structure consisting of three fields (3 machine words):
1. **Pointer**: A pointer to the start of the data in the underlying array.
2. **Length (len)**: The number of elements in the slice.
3. **Capacity (cap)**: Capacity — the maximum number of elements the slice can hold without allocating new memory.

```go
type slice struct {
    array unsafe.Pointer // Pointer to array
    len   int            // Current length
    cap   int            // Capacity
}
```

### 🚀 `append` Mechanics
When you call `append`, Go checks if there is enough capacity (`cap`):
1. If there is enough space, the element is simply added to the underlying array, and `len` increases.
2. If there is **not enough space**:
   - A new, larger underlying array is allocated (typically the size doubles, with a smaller coefficient for very large slices).
   - The old data is copied to the new array.
   - A new slice descriptor with an updated pointer is returned.

> [!IMPORTANT]
> Always assign the result of `append` back to the variable: `s = append(s, val)`.

---

# 3. ⚠️ Important Nuances (Gotchas)

### 🔗 Shared Underlying Array
Multiple slices can point to the same array. Changing an element in one of them affects the others.

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3] // [2 3], refers to the same array
b[0] = 99   // Now a[1] is also 99!
```

### 💧 Memory Leaks
If you create a small slice based on a huge array and store it for a long time, the GC won't be able to clean up the entire array because there is a reference to it.
**Solution**: Use `copy()` into a new slice.

---

# 4. 📊 Comparison

| Feature | Array | Slice |
| :--- | :--- | :--- |
| **Size** | Fixed (constant) | Dynamic (changeable) |
| **Transmission** | Copying data (Costly) | Copying header (Cheap) |
| **Type** | Size is part of the type | Only the element type |
| **Usage** | Rare (mostly low-level) | The main tool in Go |

---

# 💡 Use Example
```go
// Creating with memory pre-allocation (efficient)
s := make([]int, 0, 10) // len=0, cap=10

// Adding
for i := 0; i < 5; i++ {
    s = append(s, i)
}

// Copying
dst := make([]int, len(s))
copy(dst, s)
```
