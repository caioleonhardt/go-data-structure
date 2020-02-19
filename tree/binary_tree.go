package tree

import (
	"fmt"
)

// BinaryTree binary tree data structure
type BinaryTree struct {
	//root
	root *Node
}

// NewBinaryTree returns a pointer to a new Binary Tree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// Insert the element in the tree
func (t *BinaryTree) Insert(i int) {
	if t.root == nil {
		t.root = &Node{i, nil, nil}
		return
	}

	insert(t.root, i)
}

func insert(n *Node, i int) {
	if i < n.val {
		if n.left == nil {
			n.left = &Node{i, nil, nil}
		} else {
			insert(n.left, i)
		}
	} else if i > n.val {
		if n.right == nil {
			n.right = &Node{i, nil, nil}
		} else {
			insert(n.right, i)
		}
	}
}

// Search the element in the tree
func (t *BinaryTree) Search(i int) (*Node, error) {
	var node *Node

	t.Walk(func(n Node) {
		if i == n.val {
			node = &n
		}
	})

	if node == nil {
		return nil, fmt.Errorf("NotFound element %d", i)
	}

	return node, nil
}

// Delete the element in the tree
func (t *BinaryTree) Delete(i int) error {
	panic("not implemented") // TODO: Implement
}

// Walk traverses the tree
func (t *BinaryTree) Walk(c Callback) {
	inOrder(t.root, c)
}

func inOrder(n *Node, c Callback) {
	if n.left != nil {
		inOrder(n.left, c)
	}

	c(*n)

	if n.right != nil {
		inOrder(n.right, c)
	}
}
