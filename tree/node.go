package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("setting Value to nil node. Ignored")
		return
	}
	node.Value = Value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func CreateTreeNode(Value int) *Node {
	return &Node{Value: Value}
}
