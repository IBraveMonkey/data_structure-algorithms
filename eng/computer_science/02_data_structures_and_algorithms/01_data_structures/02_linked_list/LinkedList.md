# 🔗 Linked List

**Description**: 
A Linked List is a dynamic data structure consisting of **nodes**. Each node contains two fields: the data itself and a **reference (pointer)** to the next node in the sequence. Unlike an array, list elements can be scattered across different parts of RAM.

- **How it works internally**: The list starts with a **Head**. To find the 5th element, the computer cannot simply calculate its address like in an array. It must literally "walk" through the chain: from the head to the first, from the first to the second, and so on until it reaches the desired node. This is why access by index takes **O(n)**.
- **Analogy**: Imagine a "Treasure Hunt" game. You have the first note (head), which tells you where the clue is hidden and where to find the next note. You don't know where the final prize is until you follow the entire chain of notes one by one.


### Types of Lists
- **Singly Linked**: Each node only knows about the *next* one.
- **Doubly Linked**: Each node knows about both the *next* and the *previous* one. This allows for efficient movement in both directions.
- **Circular**: The last node references the first node, closing the loop.


### Pros and Cons
✅ **Pros**:
1. **Dynamic Size**: The list grows and shrinks easily without needing to reallocate large contiguous blocks of memory.
2. **Fast Insert/Delete**: If you already have a pointer to the desired position, insertion or deletion happens instantly (O(1)) because you just need to reassign the "links" without shifting the actual data.
3. **Efficient Use of Fragmented Memory**: The list doesn't require a single large block of memory; it fills up available "holes".

❌ **Cons**:
1. **Slow Access**: To reach the middle or end, you must traverse all elements from the beginning.
2. **Extra Memory**: For each element, you must store one or two pointers, which increases memory consumption.
3. **Poor Cache Friendliness**: Nodes are scattered across memory, so the CPU spends more time fetching them.

---


### Visualization

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph LR
    Head[Head] --> N1[Node 1<br/>data: 10<br/>next: →]
    N1 --> N2[Node 2<br/>data: 20<br/>next: →]
    N2 --> N3[Node 3<br/>data: 30<br/>next: nil]



linkStyle default stroke:#009688,stroke-width:2px;




```


### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Access (by index) | O(n) | O(1) |
| Insertion (at the beginning) | O(1) | O(1) |
| Insertion (at the end) | O(n)* | O(1) |
| Insertion (in the middle) | O(n) | O(1) |
| Deletion (from the beginning) | O(1) | O(1) |
| Deletion (from the end) | O(n)* | O(1) |
| Deletion (from the middle) | O(n) | O(1) |
| Search | O(n) | O(1) |
| Storage | — | O(n) |

> [!NOTE]
> *For a doubly linked list, insertion/deletion at the end is O(1) if a pointer to the tail is maintained.


### When to Use

✅ Frequent insertions/deletions at the beginning  
✅ Direct access by index is not needed  
✅ Doubly linked — if operations from the end are needed


## 💻 Implementation
```go
package linked_list

// Implementation...
func (l *LinkedList) AddToFront(value int) {
    // ...
}
```

```javascript
// LinkedList Implementation (JS)
class Node {
  constructor(value) {
    this.value = value;
    this.next = null;
  }
}

class LinkedList {
  constructor() {
    this.head = null;
    this.tail = null;
    this.size = 0;
  }

  addToFront(value) {
    const newNode = new Node(value);
    newNode.next = this.head;
    this.head = newNode;
    if (!this.tail) this.tail = newNode;
    this.size++;
  }

  addToBack(value) {
    const newNode = new Node(value);
    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      this.tail.next = newNode;
      this.tail = newNode;
    }
    this.size++;
  }

  reverse() {
    let prev = null, current = this.head;
    this.tail = this.head;
    while (current) {
      let nextTemp = current.next;
      current.next = prev;
      prev = current;
      current = nextTemp;
    }
    this.head = prev;
  }
}
```

```


## 🚀 Practical Problems
```go
package linked_list

// Problems on Go...
func HasCycle(l *LinkedList) bool {
    // ...
}
```

```javascript
// Algorithmic Problems (JS)

// 1. Cycle Detection
function hasCycle(head) {
  let slow = head, fast = head;
  while (fast && fast.next) {
    slow = slow.next;
    fast = fast.next.next;
    if (slow === fast) return true;
  }
  return false;
}

// 2. Find Middle
function middleNode(head) {
  let slow = head, fast = head;
  while (fast && fast.next) {
    slow = slow.next;
    fast = fast.next.next;
  }
  return slow;
}

// 3. Remove N-th from End
function removeNthFromEnd(head, n) {
  let dummy = new Node(0);
  dummy.next = head;
  let fast = dummy, slow = dummy;
  for (let i = 0; i <= n; i++) fast = fast.next;
  while (fast) {
    fast = fast.next;
    slow = slow.next;
  }
  slow.next = slow.next.next;
  return dummy.next;
}

// 4. Merge Two Sorted Lists
function mergeTwoLists(l1, l2) {
  let dummy = new Node(0);
  let curr = dummy;
  while (l1 && l2) {
    if (l1.value <= l2.value) {
      curr.next = l1;
      l1 = l1.next;
    } else {
      curr.next = l2;
      l2 = l2.next;
    }
    curr = curr.next;
  }
  curr.next = l1 || l2;
  return dummy.next;
}
```

<!-- QUIZ_START 
[
    {
        "question": "How does a Linked List differ from an Array in terms of memory allocation?",
        "options": ["Lists use a contiguous block of RAM", "Array elements can be scattered, while List elements must be sequential", "Linked List elements (nodes) can be scattered across different parts of RAM, connected by pointers", "There is no difference in memory allocation"],
        "correctIndex": 2
    },
    {
        "question": "What is the time complexity of accessing an element by index in a singly linked list?",
        "options": ["O(1)", "O(log n)", "O(n)", "O(n²)"],
        "correctIndex": 2
    },
    {
        "question": "What is a significant advantage of a Doubly Linked List over a Singly Linked List?",
        "options": ["It uses less memory", "It allows efficient traversal and deletion in both directions (forward and backward)", "It is faster for searching random elements", "It doesn't require pointers"],
        "correctIndex": 1
    }
]
QUIZ_END -->

```

