package search_three

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func NewNode(val int) *Node {

	return &Node{Value: val}
}
