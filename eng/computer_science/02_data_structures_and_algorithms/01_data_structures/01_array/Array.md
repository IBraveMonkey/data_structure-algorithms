# 📋 Array

**Description**: 
An Array is a fundamental data structure that serves as a container for storing elements of the same type in a **contiguous block of memory**. It is the simplest and most efficient structure for organizing data when access speed is critical.

- **How it works internally**: Since elements are stored directly next to each other, the computer can instantly calculate the address of any element using the formula: `start_address + (index * element_size)`. This is why access by index always takes **O(1)** time.
- **Analogy**: Imagine lockers in a gym's changing room. They are all the same size and follow a strict numerical order. If you know the locker number (index), you can walk directly to the correct door.


### Types of Arrays
- **Classical (Static)**: The size is fixed at creation. In Go, this is represented as `[N]T`.
- **Dynamic (Slice)**: Capable of "growing". When the current capacity is reached, the structure allocates a new, larger memory block and copies the data over. In Go, this is implemented as `[]T`.


### Pros and Cons
✅ **Pros**:
1. **Instant Access**: Reading and writing by index is extremely fast.
2. **Cache Friendly**: Because data is stored sequentially, the CPU can easily prefetch it into the cache, speeding up the processing of large datasets.
3. **Memory Efficiency**: Minimum overhead — memory is spent almost exclusively on the actual data.

❌ **Cons**:
1. **Expensive Insert/Delete**: Inserting or removing an element from the beginning or middle requires shifting all subsequent elements (O(n)).
2. **Contiguous Memory Requirement**: A very large array might fail to allocate if a single continuous block of memory isn't available, even if total free space is sufficient (fragmentation).
3. **Fixed Type**: All elements must be of the same size and type.

---


### Basic Operations


#### Classical Array:
- Access (read/write by index)
- Search (linear, if the array is not sorted)


#### Dynamic Array:
- All operations of classical array plus:
- Insertion (at the end or in the middle)
- Deletion (from the end or the middle)


### Complexity

| Operation | Classical Array | Dynamic Array |
|:---|:---:|:---:|
| Access (by index) | O(1) | O(1) |
| Insertion (at the end) | N/A | O(1) amortized |
| Insertion (in the middle) | N/A | O(n) |
| Deletion (from the end) | N/A | O(1) |
| Deletion (from the middle) | N/A | O(n) |
| Search (unsorted) | O(n) | O(n) |
| Storage | O(n) | O(n) |

> [!NOTE]
> **Access O(1)**: The index directly points to a memory cell. **Insertion/deletion O(n)**: In the middle of the array, it requires shifting elements.
> 
> **N/A** — operation is not applicable (e.g., classical arrays have a fixed size at creation, so insertion is not possible).
> 
> **Amortized Complexity O(1)** — means that most insertions take instant time, but occasionally (when space runs out) an expensive memory reallocation occurs. On average (over time), it still counts as O(1).

> [!TIP]
> In Go, standard arrays have fixed size, but slices (`[]T`) are more commonly used, which implement the dynamic array concept.


### When to Use

**Classical Array:**

✅ Fast access by index is needed  
✅ Data size is exactly known and does not change  
✅ Memory usage control is important  
❌ Not suitable if data size may change  

**Dynamic Array:**

✅ Fast access by index is needed  
✅ Data size may change during execution  
✅ Convenience of working with variable-size collections  
❌ Less efficiency with frequent insertions/deletions in the middle  
❌ Possible overhead for memory reallocation


## 💻 Implementation

```go
package array

import "fmt"

// BasicOperations demonstrates basic operations with arrays and slices
func BasicOperations() {
	// 1. Array declaration (fixed size)
	var arr [5]int
	arr[0] = 10 // O(1)

	// 2. Slice declaration (dynamic array)
	slice := []int{1, 2, 3}

	// 3. Appending to the end (Append)
	// Average O(1), but if capacity is exhausted,
	// reallocation and copying of the array will occur (O(n)).
	slice = append(slice, 4)

	// 4. Insertion in the middle (O(n))
	index := 1
	value := 99
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value

	// 5. Deletion from the middle (O(n))
	// Preserving order by shifting elements
	slice = append(slice[:index], slice[index+1:]...)

	fmt.Println("Final slice:", slice)
}
```

```javascript
// Basic array operations in JS
function basicOperations() {
    // 1. Declaration (JS doesn't have native fixed-size arrays)
    const arr = new Array(5).fill(0);
    arr[0] = 10; // O(1)

    // 2. Creation of a dynamic array
    let array = [1, 2, 3];

    // 3. Insertion at the end (push)
    // O(1) amortized average
    array.push(4);

    // 4. Insertion in the middle (O(n))
    const index = 1;
    const value = 99;
    array.splice(index, 0, value);

    // 5. Deletion from the middle (O(n))
    array.splice(index, 1);

    console.log("Final array:", array);
}
```


## 🚀 Practical Problems

```go
package array

import "fmt"

// Examples on Go...
func Examples() {
    // ...
}
```

```javascript
// Practical techniques (JS)
function examples() {
    // Array Reverse (In-place) - O(n)
    const data = [1, 2, 3, 4, 5];
    data.reverse();
    console.log("Reverse:", data);

    // Finding the maximum - O(n)
    const max = Math.max(...data);

    // Filtering (creating a new array) - O(n)
    const even = data.filter(v => v % 2 === 0);
}
```

<!-- QUIZ_START 
[
    {
        "question": "Why does accessing an array element by index take O(1) time?",
        "options": ["Because the computer searches every element", "Because elements are stored in a contiguous block, allowing instant address calculation", "Because arrays are always sorted", "Because Go optimizes all array operations"],
        "correctIndex": 1
    },
    {
        "question": "What is a major disadvantage of inserting an element into the middle of a large array?",
        "options": ["It requires extra RAM for a new array", "It requires shifting all subsequent elements, taking O(n) time", "It causes the computer to crash", "It changes the data types of other elements"],
        "correctIndex": 1
    },
    {
        "question": "In Go, what is the 'dynamic' version of an array called?",
        "options": ["Map", "Struct", "Slice", "Channel"],
        "correctIndex": 2
    }
]
QUIZ_END -->

```

