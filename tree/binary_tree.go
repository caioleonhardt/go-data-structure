package tree

// BinaryTree binary tree data structure
type BinaryTree struct {
	//root
	root *Node
}

// NewBinaryTree returns a pointer to a new Binary Tree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(i int) {
	panic("not implemented") // TODO: Implement
}

func (t *BinaryTree) Search(i int) (Node, error) {
	panic("not implemented") // TODO: Implement
}

func (t *BinaryTree) Delete(i int) error {
	panic("not implemented") // TODO: Implement
}

func (t *BinaryTree) Walk(c Callback) {
	panic("not implemented") // TODO: Implement
}
