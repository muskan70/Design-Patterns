package node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewNode(x int) *Node {
	return &Node{
		Val:   x,
		Left:  nil,
		Right: nil,
	}
}
