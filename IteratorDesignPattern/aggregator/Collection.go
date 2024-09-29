package aggregator

import iterator "tree/iterator"

type BinaryTreeCollection interface {
	CreatePreOrderIterator() iterator.IteratorInterface
	CreateInOrderIterator() iterator.IteratorInterface
	CreatePostOrderIterator() iterator.IteratorInterface
}
