package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func CreateExampleNode() *Node {
	return &Node{
		Value: 3,
		Left: &Node{
			Value: 0,
			Left:  nil,
			Right: &Node{Value: 2},
		},
		Right: &Node{
			Value: 5,
			Left:  &Node{Value: 0},
			Right: nil,
		},
	}
}
