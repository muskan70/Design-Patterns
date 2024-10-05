package iterator

import node "tree/nodeClass"

type PreOrderIterator struct {
	nodeList []*node.Node
	index    int
}

func NewPreOrderIterator(root *node.Node) *PreOrderIterator {
	preorder := &PreOrderIterator{index: 0}
	preorder.PreOrderTraversal(root)
	return preorder
}

func (ord *PreOrderIterator) PreOrderTraversal(root *node.Node) {
	if root == nil {
		return
	}
	ord.nodeList = append(ord.nodeList, root)
	ord.PreOrderTraversal(root.Left)
	ord.PreOrderTraversal(root.Right)
}

func (p *PreOrderIterator) HasNext() bool {
	return p.index < len(p.nodeList)
}

func (p *PreOrderIterator) Next() *node.Node {
	if p.HasNext() {
		p.index++
		return p.nodeList[p.index-1]
	}
	return nil
}
