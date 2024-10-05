package iterator

import (
	node "tree/nodeClass"
)

type InOrderIterator struct {
	nodeList []*node.Node
	index    int
}

func NewInOrderIterator(root *node.Node) *InOrderIterator {
	inorder := &InOrderIterator{index: 0}
	inorder.InOrderTraversal(root)
	return inorder
}

func (inOrd *InOrderIterator) InOrderTraversal(root *node.Node) {
	if root == nil {
		return
	}
	inOrd.InOrderTraversal(root.Left)
	inOrd.nodeList = append(inOrd.nodeList, root)
	inOrd.InOrderTraversal(root.Right)
}

func (i *InOrderIterator) HasNext() bool {
	return i.index < len(i.nodeList)
}

func (i *InOrderIterator) Next() *node.Node {
	if i.HasNext() {
		i.index++
		return i.nodeList[i.index-1]
	}
	return nil
}
