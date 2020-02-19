package tree

import (
	"fmt"
)

func ExampleBinaryTree_Walk() {
	t := NewBinaryTree()

	t.Insert(32)
	t.Insert(55)
	t.Insert(79)
	t.Insert(10)
	t.Insert(1)
	t.Insert(19)
	t.Insert(23)
	t.Insert(16)

	t.Walk(func(n Node) {
		fmt.Printf("%d -> ", n.val)
	})
	// Output: 1 -> 10 -> 16 -> 19 -> 23 -> 32 -> 55 -> 79 ->
}
