# 🧵 Strings in Go

A string in Go is an **immutable sequence of bytes**. It is important to understand: a string is **not a slice** of bytes, although they are similar in behavior. Typically, Go strings store text in **UTF-8** encoding.

---


# 1. 🏗️ Internal Structure

In Go, a string (`string`) is a lightweight structure of two fields (StringHeader):
1. **Pointer**: Pointer to a byte array (read-only).
2. **Length (len)**: The length of the string (number of bytes).

```go
type StringHeader struct {
    Data uintptr
    Len  int
}
```

> [!IMPORTANT]
> Unlike `slice`, a string **has no Capacity (cap) field**. A string always matches the size of the data it points to.


### ⚙️ Key Properties
- **Immutability**: You cannot change a single byte in a string (`s[0] = 'a'` will cause a compilation error). Any "change" creates a new string.
- **Slice Compatibility**: Conversion `[]byte(s)` or `string(bytes)` creates a copy of the data in the heap to guarantee string immutability.
- **Comparison**: Strings can be compared using `==` and `!=`. This compares the content byte by byte, not memory addresses.

---


# 2. 🔄 Iteration (Two Ways)

In Go, iteration over a string works differently depending on the method.


### 🔢 Using a simple `for` (by bytes)
When you use the classic loop `for i := 0; i < len(s); i++`, you retrieve **bytes**.

```go
s := "Hello"
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i]) // Will print UTF-8 byte HEX codes
}
```


### 🔤 Using `range` (by runes)
The `range` loop automatically decodes UTF-8 and yields **runes** (`rune` is an alias type for `int32` representing a Unicode character).

```go
s := "Hello"
for index, runeValue := range s {
    fmt.Printf("%d: %c ", index, runeValue)
}
```

> [!CAUTION]
> When iterating via `range`, indices may "jump" (e.g., 0, 2, 4...), as one character in UTF-8 can take up multiple bytes.

---


# 3. 🧩 Rune

**Rune** is a Unicode character (Code Point).
- In Go `rune` is a synonym for `int32`.
- If you need to work with characters (e.g., reverse a string or count the number of letters), always convert the string to `[]rune`.

```go
s := "Gopher"
runes := []rune(s)
fmt.Println(len(runes)) // Will return the number of CHARACTERS, not bytes
```

---


# 4. 🚀 Efficiency

- **Copying**: Copying a string is very cheap (only the header is copied — pointer and length).
- **Slicing**: Getting a substring `s[1:4]` is also cheap, since the new string points to the same memory area in the underlying byte array.
- **Concat**: Concatenating strings via `+` is efficient for small volumes. To assemble large strings from many parts, use `strings.Builder`.

---


# 📊 Summary

| Operation | Result |
| :--- | :--- |
| `len(s)` | Number of **bytes** |
| `s[i]` | **Byte** by index |
| `for r := range s` | Iterates over **runes** (Unicode characters) |
| `[]rune(s)` | Conversion to character array (expensive) |
| `[]byte(s)` | Conversion to byte array (expensive) |

<!-- QUIZ_START 

[
    {
        "question": "What two fields make up a string header in Go?",
        "options": [
            "Pointer to array, Capacity",
            "Pointer to byte array, Length",
            "Length, Hash",
            "Value, Type"
        ],
        "correctIndex": 1
    },
    {
        "question": "Can you change a specific character in a Go string (e.g., s[0] = 'A')?",
        "options": [
            "Yes, if it's a short string",
            "No, strings in Go are immutable",
            "Yes, using the 'unsafe' package only",
            "Only if the string was created using make()"
        ],
        "correctIndex": 1
    },
    {
        "question": "What is the difference between simple for-loop and for-range loop when iterating over a string?",
        "options": [
            "Simple for iterates over characters, range over bytes",
            "Simple for iterates over bytes, range over runes (Unicode characters)",
            "There is no difference",
            "Range is slower but uses less memory"
        ],
        "correctIndex": 1
    }
]

QUIZ_END -->

