package tree

import (
	"fmt"
	"strconv"
	"strings"
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

	t.Walk(func(n *Node) {
		if i == n.val {
			node = n
		}
	})

	if node == nil {
		return nil, fmt.Errorf("NotFound element %d", i)
	}

	return node, nil
}

// Delete the element in the tree
func (t *BinaryTree) Delete(i int) error {
	if t.root == nil {
		return fmt.Errorf("EmptyTree")
	}
	t.root = delete(t.root, i)
	return nil
}

func delete(n *Node, i int) *Node {
	if i < n.val {
		n.left = delete(n.left, i)
	} else if i > n.val {
		n.right = delete(n.right, i)
	} else {
		// found leaf
		if n.left == nil && n.right == nil {
			return nil
		}

		// found one left leaf
		// or one right leaf
		if n.right == nil {
			tmp := n.left
			n = nil
			return tmp
		} else if n.left == nil {
			tmp := n.right
			n = nil
			return tmp
		}

		tmp := predecessor(n.left)
		n.val = tmp.val
		n.left = delete(n.left, n.val)
	}

	return n
}

func predecessor(n *Node) *Node {
	if n.right != nil {
		return predecessor(n.right)
	}

	return n
}

// Walk traverses the tree
func (t *BinaryTree) Walk(c Callback) {
	inOrder(t.root, c)
}

func inOrder(n *Node, c Callback) {
	if n == nil {
		return
	}

	if n.left != nil {
		inOrder(n.left, c)
	}

	c(n)

	if n.right != nil {
		inOrder(n.right, c)
	}
}

func (t *BinaryTree) String() string {
	s := []string{}
	t.Walk(func(n *Node) {
		if n == nil {
			return
		}
		s = append(s, strconv.Itoa(n.val))
	})

	return strings.Join(s, " -> ")
}

func (n *Node) String() string {
	return strconv.Itoa(n.val)
}
