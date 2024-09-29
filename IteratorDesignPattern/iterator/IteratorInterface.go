package iterator

import node "tree/nodeClass"

type IteratorInterface interface {
	HasNext() bool
	Next() *node.Node
}
