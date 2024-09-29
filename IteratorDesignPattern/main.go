package main

import (
	"fmt"
	"tree/aggregator"
	node "tree/nodeClass"
)

func main() {
	node0 := node.NewNode(0)
	node1 := node.NewNode(1)
	node2 := node.NewNode(2)
	node3 := node.NewNode(3)
	node4 := node.NewNode(4)
	node5 := node.NewNode(5)
	node6 := node.NewNode(6)
	node7 := node.NewNode(7)

	node0.Left = node1
	node0.Right = node2

	node1.Left = node3
	node1.Right = node4

	node2.Left = node5
	node2.Right = node6

	node3.Right = node7

	tree := aggregator.NewBinaryTree(node0)

	fmt.Println("Inorder Traversal:")

	itr1 := tree.CreateInOrderIterator()
	for itr1.HasNext() {
		fmt.Print(itr1.Next().Val, " ")
	}

	fmt.Println()
	fmt.Println("Preorder Traversal:")

	itr2 := tree.CreatePreOrderIterator()
	for itr2.HasNext() {
		fmt.Print(itr2.Next().Val, " ")
	}

	fmt.Println()
	fmt.Println("Postorder Traversal:")

	itr3 := tree.CreatePostOrderIterator()
	for itr3.HasNext() {
		fmt.Print(itr3.Next().Val, " ")
	}
}
