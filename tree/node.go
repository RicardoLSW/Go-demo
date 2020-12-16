package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("setting Value to nil node. Ignored")
		return
	}
	node.Value = Value
}

func CreateTreeNode(Value int) *Node {
	return &Node{Value: Value}
}
