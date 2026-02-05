package red_black_tree

import "fmt"

/* Red-Black Tree
It is a self-balancing binary search tree. It ensures that the tree height remains logarithmic even after multiple insertions and deletions.

It is a regular binary search tree (BST), BUT with an additional balancing property that ensures uniform growth.

- Each node stores one extra bit - color (red or black).

Main properties of a red-black tree:
	1️⃣ Every node is either RED or BLACK.
	2️⃣ The root is always black.
	3️⃣ Every path from the root to a Nil node contains the same number of black nodes (this is called the black height).
	4️⃣ A red node cannot have red children (i.e., red nodes must be separated by black nodes).
	5️⃣ Every new node is added as RED.

🔹 These rules ensure balancing, preventing skewing.

### Complexity

| Operation | Average (O) | Worst (O) | Space Complexity (O) |
|:---|:---:|:---:|:---:|
| Insertion | O(log n) | O(log n) | O(log n) |
| Search | O(log n) | O(log n) | O(1) iterative / O(log n) recursive |
| Deletion | O(log n) | O(log n) | O(log n) |
| Storage | — | — | O(n) |

*Guarantees logarithmic height h < 2 * log2(n + 1).
*/

const (
	RED   = true
	BLACK = false
)

type RBTNode struct {
	Val    int
	Color  bool
	Left   *RBTNode
	Right  *RBTNode
	Parent *RBTNode
}

type RBTree struct {
	Root *RBTNode
}

func NewNode(val int) *RBTNode {
	return &RBTNode{Val: val, Color: RED} // Always insert as red
}

// LeftRotate
func (t *RBTree) LeftRotate(x *RBTNode) {
	y := x.Right
	x.Right = y.Left

	if y.Left != nil {
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

// RightRotate
func (t *RBTree) RightRotate(x *RBTNode) {
	y := x.Left
	x.Left = y.Right

	if y.Right != nil {
		y.Right.Parent = x
	}

	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

// Insert - inserts a new node
func (t *RBTree) Insert(val int) {
	newNode := NewNode(val)
	if t.Root == nil {
		newNode.Color = BLACK // Root is always black
		t.Root = newNode
	} else {
		t.insertRecursive(t.Root, newNode)
	}
	t.fixInsert(newNode) // Balancing
}

// Recursive insertion in BST
func (t *RBTree) insertRecursive(root, node *RBTNode) {
	if node.Val < root.Val {
		if root.Left == nil {
			root.Left = node
			node.Parent = root
		} else {
			t.insertRecursive(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
		} else {
			t.insertRecursive(root.Right, node)
		}
	}
}

// Balancing after insertion
func (t *RBTree) fixInsert(node *RBTNode) {
	// TODO: implement balancing rules
}

// PrintTree - tree output (simplified)
func (t *RBTree) PrintTree(node *RBTNode, indent string) {
	if node != nil {
		t.PrintTree(node.Right, indent+"   ")
		color := "R"
		if node.Color == BLACK {
			color = "B"
		}
		fmt.Println(indent, node.Val, color)
		t.PrintTree(node.Left, indent+"   ")
	}
}
