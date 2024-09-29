package iterator

import node "tree/nodeClass"

type PostOrderIterator struct {
	nodeList []*node.Node
	index    int
}

func NewPostOrderIterator(root *node.Node) *PostOrderIterator {
	postorder := &PostOrderIterator{index: 0}
	postorder.PostOrderTraversal(root)
	return postorder
}

func (p *PostOrderIterator) PostOrderTraversal(root *node.Node) {
	if root == nil {
		return
	}
	p.PostOrderTraversal(root.Left)
	p.PostOrderTraversal(root.Right)
	p.nodeList = append(p.nodeList, root)
}

func (p *PostOrderIterator) HasNext() bool {
	return p.index < len(p.nodeList)
}

func (p *PostOrderIterator) Next() *node.Node {
	if p.HasNext() {
		p.index++
		return p.nodeList[p.index-1]
	}
	return nil
}
