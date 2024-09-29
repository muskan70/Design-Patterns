package aggregator

import (
	"tree/iterator"
	node "tree/nodeClass"
)

type BinaryTree struct {
	root *node.Node
}

func NewBinaryTree(temp *node.Node) *BinaryTree {
	return &BinaryTree{
		root: temp,
	}
}

func (b *BinaryTree) CreateInOrderIterator() iterator.IteratorInterface {
	return iterator.NewInOrderIterator(b.root)
}

func (b *BinaryTree) CreatePreOrderIterator() iterator.IteratorInterface {
	return iterator.NewPreOrderIterator(b.root)
}

func (b *BinaryTree) CreatePostOrderIterator() iterator.IteratorInterface {
	return iterator.NewPostOrderIterator(b.root)
}
