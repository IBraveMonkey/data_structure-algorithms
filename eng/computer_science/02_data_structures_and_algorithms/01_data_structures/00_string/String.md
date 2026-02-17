# 🧵 Strings

**Description**: 
A string is a fundamental data structure representing a sequence of characters. In modern computer science, strings are viewed not just as text, but as an abstraction over an array of bytes with specific interpretation rules and memory management strategies.

---


## 🏗️ Internal Representation

At a low level, there are two primary approaches to representing strings in memory:

1.  **Null-terminated (C-style strings)**:
    *   **How it works**: Characters are placed sequentially, and the end of the string is marked by a special "null" byte (`\0`).
    *   **Pros**: Memory efficient (only 1 byte for the end marker).
    *   **Cons**: Determining string length requires scanning the entire sequence ($O(n)$ complexity). It also introduces security risks like buffer overflows.

2.  **Length-prefixed (Pascal/Go/Java-style strings)**:
    *   **How it works**: An integer representing the length is stored at the beginning (in a string descriptor).
    *   **Pros**: Instant length retrieval ($O(1)$) and enhanced security.
    *   **Cons**: Requires slightly more memory to store metadata (the descriptor).


### 🧱 Hierarchy: Byte vs. Character
Modern strings are typically sequences of bytes interpreted through an **encoding**:
*   **ASCII**: 1 byte per character. Limited to Latin characters and basic symbols.
*   **Unicode (UTF-8)**: Variable length (1 to 4 bytes). The most popular format for web and modern languages.
*   **UTF-16/UTF-32**: Fixed or semi-fixed width, often used in internal system APIs (e.g., Windows, Java/C#).

> [!IMPORTANT]
> String length in bytes and the number of visible characters are often different. For example, the 🐵 emoji takes 4 bytes in UTF-8 but is considered a single character.

---


## 🔒 Immutability Principle

In many modern languages (Java, Python, Go, JavaScript), strings are **immutable**.

**Why is this necessary?**
1.  **Safety**: A string won't change "under the hood" while being passed between functions or threads.
2.  **Optimization (Interning)**: The system can store only one copy of identical strings in memory to save resources.
3.  **Hash Caching**: Since the string doesn't change, its hash can be computed once, speeding up its use as a key in HashMaps/Dictionaries.

---


## 💻 Implementation (Go)

Go adopts a modern approach: strings are immutable, represented by a descriptor (pointer + length), and use UTF-8 by default.

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// In Go...
	s := "Hello, Gopher! 🐹"
	fmt.Printf("Bytes: %d\n", len(s)) 
	fmt.Printf("Characters: %d\n", utf8.RuneCountInString(s))
}
```

```javascript
// Strings in JS
const s = "Hello, JS! 🐹";

// 1. .length returns the number of UTF-16 code units
// For many emojis, this counts as 2 units
console.log(`Length: ${s.length}`); // 13

// 2. To count actual characters (code points)
const charCount = [...s].length;
console.log(`Characters: ${charCount}`); // 12

// 3. Immutability: operations return NEW strings
const upper = s.toUpperCase();
```


### Efficient Usage
Because strings are immutable, the operation `s = s + "more"` inside a loop creates a new string copy on every iteration. This is extremely slow ($O(n^2)$).
To build strings efficiently, use dedicated buffers (like `strings.Builder` in Go or `StringBuilder` in Java).

---


## 🚀 Algorithms and Complexity

| Operation | Complexity (Length-prefixed) | Complexity (Null-terminated) |
| :--- | :--- | :--- |
| **Get Length** | $O(1)$ | $O(n)$ |
| **Index Access** | $O(1)$ (to byte) | $O(n)$ (must scan to index) |
| **Concatenation** | $O(n+m)$ | $O(n+m)$ |
| **Substring Search** | $O(n \cdot m)$ (naive) | $O(n \cdot m)$ |

**Popular Algorithms**:
*   **Knuth-Morris-Pratt (KMP)**: Efficient substring search in $O(n+m)$.
*   **Boyer-Moore**: Used in most standard libraries for high-performance searching.

<!-- QUIZ_START 
[
    {
        "question": "What is the main advantage of Length-prefixed strings (like in Go or Java) over Null-terminated strings (like in C)?",
        "options": [
            "They occupy less memory", 
            "Retrieving string length is an O(1) operation", 
            "They always use UTF-16", 
            "They are easier to print"
        ],
        "correctIndex": 1
    },
    {
        "question": "Why is performing concatenation (+) in a loop considered bad practice for immutable strings?",
        "options": [
            "It leads to compilation errors", 
            "It is only safe for ASCII characters", 
            "A new string copy is created every iteration, resulting in O(n²) complexity", 
            "It slows down garbage collection"
        ],
        "correctIndex": 2
    },
    {
        "question": "How many bytes can a single character occupy in UTF-8 encoding?",
        "options": [
            "Always 1 byte", 
            "Between 1 and 4 bytes", 
            "Always 2 bytes", 
            "Up to 8 bytes"
        ],
        "correctIndex": 1
    }
]
QUIZ_END -->
