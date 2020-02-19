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

	prev := t.root
	for cur := t.root; cur != nil; {
		if i == cur.val {
			// found leaf
			if cur.left == nil && cur.right == nil {
				if prev.left == cur {
					prev.left = nil
				} else if prev.right == cur {
					prev.right = nil
				}
				return nil
			}

			if cur.right == nil {
				if prev.left == cur {
					prev.left = cur.left
				} else if prev.right == cur {
					prev.right = cur.left
				}
				return nil
			} else if cur.left == nil {
				return nil
			}
		}

		// update pointers
		prev = cur
		if i < cur.val {
			cur = cur.left
		} else if i > cur.val {
			cur = cur.right
		}
	}

	return fmt.Errorf("NotFound")
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
