# ⚖️ Balanced Trees (AVL or Red-Black)

**Description**: 
Balanced Trees (such as AVL or Red-Black trees) are "smart" versions of standard search trees. They automatically monitor their own shape to ensure the tree remains short and bushy, rather than stretching into a long, inefficient list.

- **How it works internally**: With every insertion or deletion, the tree checks its own height. If one branch becomes significantly longer than the others, the tree performs **"rotations"** — moving nodes around to restore equilibrium.
  - *AVL Tree*: Very strict, maintaining a height difference of no more than 1. Costly to restructure but provides the fastest search times.
  - *Red-Black Tree*: More flexible. It uses node colors to manage balance. Generally faster for frequent data updates.
- **Analogy**: Imagine a gardener who prunes a tree's branches so that it grows outwards and stays lush, rather than letting it grow into a single, thin, weak branch reaching for the sky.


### Pros and Cons
✅ **Pros**:
1. **Guaranteed Speed**: Unlike a standard BST, search will always take **O(log n)**, even if data arrives in the "worst" possible order.
2. **Predictability**: You always know the worst-case performance of the algorithm without surprises.

❌ **Cons**:
1. **High Complexity**: Manually implementing rotations and recoloring rules is one of the most challenging tasks in software engineering.
2. **Overhead**: Every data modification requires extra time to verify the balance and perform the necessary rotations.

---


### Balancing Visualization

```mermaid
%%{init: { 'theme': 'base', 'themeVariables': { 'lineColor': '#009688', 'primaryColor': '#009688', 'primaryTextColor': '#009688', 'attributeBkg': '#009688', 'attributeTextColor': '#009688', 'signalColor': '#009688', 'actorLineColor': '#009688', 'nodeBorder': '#009688', 'clusterBorder': '#009688', 'textColor': '#009688', 'fontSize': '16px' } } }%%
graph TD
    subgraph Unbalanced["Unbalanced"]
        U1[3] --> U2[2]
        U2 --> U3[1]
    end
    
    subgraph Balanced["Balanced (AVL)"]
        B1[2] --> B2[1]
        B1 --> B3[3]
    end
    
    Unbalanced -->|Rotation| Balanced



linkStyle default stroke:#009688,stroke-width:2px;




```


### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Insertion | O(log n) | O(log n) (recursion) |
| Search | O(log n) | O(log n) (recursion) |
| Deletion | O(log n) | O(log n) (recursion) |
| Storage | — | O(n) |

> [!IMPORTANT]
> They guarantee O(log n) through rotations (AVL) or repainting/rotations (Red-Black). AVL is faster for searching, Red-Black is faster for insertion/deletion.


## 💻 Implementation

```go
package red_black_tree

import "fmt"

// Node color
type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

// RBTNode represents a node in the Red-Black Tree
type RBTNode struct {
	Value  int
	Color  Color
	Left   *RBTNode
	Right  *RBTNode
	Parent *RBTNode
}

// RBTree implements a self-balancing Red-Black Tree
type RBTree struct {
	Root *RBTNode
	Nil  *RBTNode // Sentinel node (always black)
}

// NewRBTree creates a new Red-Black Tree
func NewRBTree() *RBTree {
	nilNode := &RBTNode{Color: BLACK}
	return &RBTree{
		Root: nilNode,
		Nil:  nilNode,
	}
}

// LeftRotate performs a left rotation at node x
func (t *RBTree) LeftRotate(x *RBTNode) {
	y := x.Right
	x.Right = y.Left
	if y.Left != t.Nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// RightRotate performs a right rotation at node y
func (t *RBTree) RightRotate(y *RBTNode) {
	x := y.Left
	y.Left = x.Right
	if x.Right != t.Nil {
		x.Right.Parent = y
	}
	x.Parent = y.Parent
	if y.Parent == nil {
		t.Root = x
	} else if y == y.Parent.Right {
		y.Parent.Right = x
	} else {
		y.Parent.Left = x
	}
	x.Right = y
	y.Parent = x
}

// Insert adds a value and maintains tree properties
func (t *RBTree) Insert(val int) {
	node := &RBTNode{
		Value: val,
		Color: RED,
		Left:  t.Nil,
		Right: t.Nil,
	}

	var y *RBTNode = nil
	x := t.Root

	for x != t.Nil {
		y = x
		if node.Value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	node.Parent = y
	if y == nil {
		t.Root = node
	} else if node.Value < y.Value {
		y.Left = node
	} else {
		y.Right = node
	}

	if node.Parent == nil {
		node.Color = BLACK
		return
	}

	if node.Parent.Parent == nil {
		return
	}

	t.fixInsert(node)
}

// fixInsert restores Red-Black Tree properties after insertion
func (t *RBTree) fixInsert(k *RBTNode) {
	for k.Parent.Color == RED {
		if k.Parent == k.Parent.Parent.Right {
			u := k.Parent.Parent.Left // uncle
			if u.Color == RED {
				u.Color = BLACK
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				k = k.Parent.Parent
			} else {
				if k == k.Parent.Left {
					k = k.Parent
					t.RightRotate(k)
				}
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				t.LeftRotate(k.Parent.Parent)
			}
		} else {
			u := k.Parent.Parent.Right // uncle
			if u.Color == RED {
				u.Color = BLACK
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				k = k.Parent.Parent
			} else {
				if k == k.Parent.Right {
					k = k.Parent
					t.LeftRotate(k)
				}
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				t.RightRotate(k.Parent.Parent)
			}
		}
		if k == t.Root {
			break
		}
	}
	t.Root.Color = BLACK
}

// Search finds a value in the tree
func (t *RBTree) Search(val int) bool {
	node := t.Root
	for node != t.Nil {
		if val == node.Value {
			return true
		}
		if val < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return false
}

// PrintTree prints the tree structure (recursively)
func (t *RBTree) PrintTree(node *RBTNode, indent string) {
	if node != t.Nil {
		color := "RED"
		if node.Color == BLACK {
			color = "BLACK"
		}
		fmt.Printf("%s%d (%s)\n", indent, node.Value, color)
		t.PrintTree(node.Left, indent+"  L: ")
		t.PrintTree(node.Right, indent+"  R: ")
	}
}
```

```javascript
/**
 * Node color constants.
 */
const RED = 'RED';
const BLACK = 'BLACK';

/**
 * RBTNode represents a node in the Red-Black Tree.
 */
class RBTNode {
  constructor(value, color = RED) {
    this.value = value;
    this.color = color;
    this.left = null;
    this.right = null;
    this.parent = null;
  }
}

/**
 * RBTree implementation with balancing logic.
 */
class RBTree {
  constructor() {
    this.TNULL = new RBTNode(0, BLACK); // Sentinel node
    this.root = this.TNULL;
  }

  // Left rotation
  leftRotate(x) {
    let y = x.right;
    x.right = y.left;
    if (y.left !== this.TNULL) {
      y.left.parent = x;
    }
    y.parent = x.parent;
    if (x.parent === null) {
      this.root = y;
    } else if (x === x.parent.left) {
      x.parent.left = y;
    } else {
      x.parent.right = y;
    }
    y.left = x;
    x.parent = y;
  }

  // Right rotation
  rightRotate(y) {
    let x = y.left;
    y.left = x.right;
    if (x.right !== this.TNULL) {
      x.right.parent = y;
    }
    x.parent = y.parent;
    if (y.parent === null) {
      this.root = x;
    } else if (y === y.parent.right) {
      y.parent.right = x;
    } else {
      y.parent.left = x;
    }
    x.right = y;
    y.parent = x;
  }

  // Insertion with balancing (fixInsert)
  insert(value) {
    let node = new RBTNode(value);
    node.left = this.TNULL;
    node.right = this.TNULL;

    let y = null;
    let x = this.root;

    while (x !== this.TNULL) {
      y = x;
      if (node.value < x.value) {
        x = x.left;
      } else {
        x = x.right;
      }
    }

    node.parent = y;
    if (y === null) {
      this.root = node;
    } else if (node.value < y.value) {
      y.left = node;
    } else {
      y.right = node;
    }

    if (node.parent === null) {
      node.color = BLACK;
      return;
    }

    if (node.parent.parent === null) {
      return;
    }

    this._fixInsert(node);
  }

  _fixInsert(k) {
    while (k.parent.color === RED) {
      if (k.parent === k.parent.parent.right) {
        let u = k.parent.parent.left; // uncle
        if (u.color === RED) {
          u.color = BLACK;
          k.parent.color = BLACK;
          k.parent.parent.color = RED;
          k = k.parent.parent;
        } else {
          if (k === k.parent.left) {
            k = k.parent;
            this.rightRotate(k);
          }
          k.parent.color = BLACK;
          k.parent.parent.color = RED;
          this.leftRotate(k.parent.parent);
        }
      } else {
        let u = k.parent.parent.right; // uncle
        if (u.color === RED) {
          u.color = BLACK;
          k.parent.color = BLACK;
          k.parent.parent.color = RED;
          k = k.parent.parent;
        } else {
          if (k === k.parent.right) {
            k = k.parent;
            this.leftRotate(k);
          }
          k.parent.color = BLACK;
          k.parent.parent.color = RED;
          this.rightRotate(k.parent.parent);
        }
      }
      if (k === this.root) break;
    }
    this.root.color = BLACK;
  }

  search(value) {
    let current = this.root;
    while (current !== this.TNULL) {
      if (value === current.value) return true;
      current = value < current.value ? current.left : current.right;
    }
    return false;
  }

  printTree(node = this.root, indent = "") {
    if (node !== this.TNULL) {
      console.log(`${indent}${node.value} (${node.color})`);
      this.printTree(node.left, indent + "  L: ");
      this.printTree(node.right, indent + "  R: ");
    }
  }
}

// Usage example
const rbt = new RBTree();
[10, 20, 30, 15].forEach(v => rbt.insert(v));
rbt.printTree();
```
`


## 🚀 Practical Problems
```go
package red_black_tree

import "fmt"

// Example demonstrates the use of a red-black tree
func Example() {
	rbt := &RBTree{}

	fmt.Println("Inserting elements: 10, 20, 30, 15")
	rbt.Insert(10)
	rbt.Insert(20)
	rbt.Insert(30)
	rbt.Insert(15)

	fmt.Println("Tree structure (Val Color):")
	// Note: full balancing is not yet implemented in red_black_tree.go
	if rbt.Root != nil {
		rbt.PrintTree(rbt.Root, "")
	} else {
		fmt.Println("Tree is empty")
	}
}

<!-- QUIZ_START 
[
    {
        "question": "What is the primary purpose of 'rotations' in balanced trees like AVL or Red-Black?",
        "options": ["To encrypt the data by moving nodes", "To restore equilibrium and ensure the tree remains 'bushy' and short", "To sort the elements in O(n log n)", "To delete all red nodes"],
        "correctIndex": 1
    },
    {
        "question": "Which of these balanced trees is generally stricter and provides faster search times?",
        "options": ["Red-Black Tree", "AVL Tree", "Standard BST", "Splay Tree"],
        "correctIndex": 1
    },
    {
        "question": "What is the guaranteed worst-case time complexity for searching in a Balanced Tree?",
        "options": ["O(n)", "O(log n)", "O(1)", "O(n!)"],
        "correctIndex": 1
    }
]
QUIZ_END -->

```

