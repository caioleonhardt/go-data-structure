package tree

// Node is the node in the tree
type Node struct {
	val int
	//left
	left *Node
	//right
	right *Node
}

// Callback to be executed for each node in Walk
type Callback func(n Node)

// Tree data structure
type Tree interface {
	Insert(i int)
	Search(i int) (Node, error)
	Delete(i int) error
	Walk(c Callback)
}
