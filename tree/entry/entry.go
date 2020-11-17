package main

import (
	"fmt"
	"go-demo/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This is a shadowed methods")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right.Left.SetValue(4)
	fmt.Println("root.Traverse")
	root.Traverse()
	fmt.Println("root.Node.Traverse")
	root.Node.Traverse()
	fmt.Println()
	root.postOrder()

}
