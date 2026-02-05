### 🗺️ What is a map?

**map** is a built-in data type in Go representing a **hash table**. It is a structure for storing "key-value" pairs that provides fast operations:
- fast adding (`m["key"] = value`).
- Search (`value := m["key"]`).
- Deletion (`delete(m, "key")`).

Ideally, operations are performed in O(1) — constant time.

---

# 1. 🏛️ Traditional Implementation (`map` < Go 1.24)

Prior to version 1.24, Go used a classic hash table architecture based on **buckets** and **overflow chains**.

## ⚙️ `map` Internals (hmap)

Internally, a `map` is implemented as an **hmap** (hash map) structure in the Go runtime:

```go
type hmap struct {
    count     int       // Number of elements
    B         uint8     // Logarithm of the number of buckets (2^B)
    buckets   []bmap    // Array of buckets
    oldbuckets []bmap   // Old array of buckets (during growth)
    hash0     uint32    // Initial value for the hash function (seed)
}
```

- **count**: How many "key-value" pairs are in the table.
- **B**: The size of the bucket array as a power of 2 ($2^B$). For example, $B = 3 \to 8$ buckets.
- **buckets**: A slice pointing to the array of buckets.
- **oldbuckets**: Used during evacuation.
- **hash0**: Random **seed** to protect against collisions.

### Hashing and Bucket Selection 🎯

1. **Hash Function**: Turns the key into a 64-bit number. Go uses optimized functions (e.g., **aeshash** for strings).
2. **Randomness**: `hash0` is added to calculations so that hashes change between runs (protection against collisions).
3. **Bucket Index**: The lower bits of the hash are taken. For 8 buckets ($B=3$) — the last 3 bits.
   - `123456789 → ...101 → bucket 5`.

## 📦 Bucket Internals (bmap)

Each bucket is a container for storing up to **8 "key-value" pairs**.

```go
bmap {
    tophash:  [8]uint8    // Top 8 bits of the hash for fast lookup
    keys:     [8]string   // 8 keys
    values:   [8]int      // 8 values
    overflow: *bmap       // Pointer to the next bucket (on overflow)
}
```

- **tophash**: Allows quickly filtering out keys without performing a full comparison.
- **Keys and Values**: Stored separately (first all 8 keys, then all 8 values) to minimize memory padding.

### Collisions and Overflow

If more than 8 elements fall into a bucket, an **overflow bucket** is created, which is attached to the current one in a chain. This makes searching slower ($O(N)$ in the worst case for one chain).

## 📈 Growth and Evacuation

> [!NOTE]
> **Evacuation** is the gradual transfer of elements to new buckets as the table grows.

1. **When?**:
   - Load factor > 6.5 (on average more than 6.5 elements per bucket).
   - Too many overflow buckets.
2. **How?**:
   - A new array of buckets 2 times larger is created ($B+1$).
   - **Lazy Evacuation**: Elements are transferred in parts during write or delete operations so as not to block the program for a long time.

---

# 2. 🚀 New Implementation (`map` Go 1.24+)

Starting from Go 1.24, the **Swiss Tables** design is used. This is a radical change that makes `map` significantly more efficient.

## What changed?

Instead of buckets, **Groups** and **Swiss Tables** algorithms are now used.

- **hmap**: The main structure now manages tables of groups.
- **Group**: Contains **8 slots** and a special **Control word** (8 bytes).
- **Control word**: Stores 7-bit hash fingerprints (**H2**) or slot statuses (`empty`, `deleted`).

### Why is it faster?

1. **SIMD**: Go checks for the presence of the desired hash in the entire group (8 slots) in **one CPU instruction**.
2. **Cache Locality**: No overflow bucket chains. Data lies denser in memory.
3. **Linear Probing**: If a group is full, we simply go to the next neighboring group.

## Operations in Swiss Tables

### Search
1. The hash is split into **H1** (points to the group) and **H2** (7-bit fingerprint).
2. Using **SIMD**, we look for **H2** in the group's `control word`.
3. If a match is found, we check the key itself.
4. If we encounter an `empty` status, the key is definitely not there.

### Deletion 🖇️
Instead of just clearing the slot, it is marked as `deleted`. This is important for linear probing so that the search does not break off prematurely.

---

## 📊 Comparison of Implementations

| Feature | Traditional (pre 1.24) | Swiss Tables (1.24+) |
| :--- | :--- | :--- |
| **Structure** | Buckets + Overflow buckets | Compact groups |
| **Search** | Linear bucket traversal | **SIMD** group check |
| **Collisions** | Chaining | Linear probing |
| **Performance** | Average | **+20-50%** to search speed |
| **Memory** | More (padding, overflow) | **-20-28%** consumption |

---

---

# 3. 🔑 Keys and Values

### What types can be keys?
In Go, a `map` key can be any type for which a comparison operation is defined (`==` and `!=`). These types are called **comparable**.

| Can be used | Cannot be used |
| :--- | :--- |
| **Simple types**: `int`, `float`, `string`, `bool` | **Slices** (`slice`) |
| **Pointers** | **Maps** (`map`) |
| **Channels** (`chan`) | **Functions** (`func`) |
| **Interfaces** (if the dynamic type is comparable) | |
| **Structs** (if all fields are comparable) | |
| **Arrays** (if elements are comparable) | |

> [!WARNING]
> If you use an **interface** as a key and put a non-comparable type (e.g., `slice`) there, the program will crash with a **panic** at runtime.

### Why can't slice or map be used? ❓
Because they are "reference" types with shallow comparison. Their contents can change, which invalidates the hash calculated earlier. The key must be "stable".

---

# 4. ⚙️ Main Properties

1. **Non-thread-safe**: `map` cannot be simultaneously read and written from different goroutines (will cause a fatal error). For concurrent work, use `sync.Mutex` or `sync.Map`.
2. **Random Order**: Iteration via `range` always returns elements in random order. This is done intentionally so that developers do not rely on the order.
3. **Reference type (almost)**: In fact, `map` is a pointer to an `hmap` structure. Therefore, when passing to a function, no copy is created, and the function can modify the original map.
4. **Zero Value**: The zero value for `map` is `nil`.
    - You can read from a `nil` map (will always return the zero value of the type).
    - You **cannot write** to a `nil` map (will cause a panic).

---

# 5. 💡 Usage Example

```go
package main

import "fmt"

func main() {
    // Creation
    m := make(map[string]int, 10) // 10 - initial capacity

    // Writing
    m["Go"] = 2009

    // Reading (comma-ok)
    val, ok := m["Go"]
    if ok {
        fmt.Println("Found:", val)
    }

    // Deletion
    delete(m, "Go")
}
```
