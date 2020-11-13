package main

import "fmt"

// e.g. using binary Tree.
//
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

type NodeIter struct {
	node *Node
	root *Node
	flag bool
}

func NewNodeIter(root *Node) *NodeIter {
	it := &NodeIter{root, root, false}
	/*
	 * Since inorder traversal start from leftmost node, so make sure
	 * you are always at the leftmost node of any given subtree
	 */
	for it.node.Left != nil {
		it.node = it.node.Left
	}
	return it
}

func (it *NodeIter) Reset() {
	it.node = it.root
	it.flag = false
}

func (it *NodeIter) Next() bool {
	if it.node == nil {
		return false
	}
	if !it.flag {
		it.flag = true
		return true
	}
	if it.node.Right != nil {
		it.node = it.node.Right
		for it.node.Left != nil {
			it.node = it.node.Left
		}
		return true
	} else {
		p := it.node.Parent
		for p != nil && it.node == p.Right {
			it.node = p
			p = p.Parent
		}
		it.node = p
		return it.node != nil
	}
}

func (it *NodeIter) Value() *Node {
	return it.node
}

func NewNode(val int) *Node {
	return &Node{val, nil, nil, nil}
}

func (n *Node) AddNode(val int) *Node {
	if n == nil {
		return NewNode(val)
	}
	var node *Node = (*Node)(nil)
	if val < n.Val {
		node = n.Left.AddNode(val)
		n.Left = node
		node.Parent = n
	} else {
		node = n.Right.AddNode(val)
		n.Right = node
		node.Parent = n
	}
	return n
}

func (n *Node) dfs() {
	if n == nil {
		return
	}
	n.Left.dfs()
	fmt.Println(n.Val)
	n.Right.dfs()
}

func main() {
	fmt.Println("vim-go")
	root := NewNode(5)
	for i := 1; i < 10; i++ {
		root.AddNode(i)
	}

	root.dfs()

	for it := NewNodeIter(root); it.Next(); {
		fmt.Println("====>", it.node.Val)
	}
	// o/p of above:-
	// vim-go
	// 1
	// 2
	// 3
	// 4
	// 5
	// 5
	// 6
	// 7
	// 8
	// 9
	// ====> 1
	// ====> 2
	// ====> 3
	// ====> 4
	// ====> 5
	// ====> 5
	// ====> 6
	// ====> 7
	// ====> 8
	// ====> 9
}
