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

func ExampleBinaryTree_Walk_empty() {
	t := NewBinaryTree()

	t.Walk(func(n Node) {
		fmt.Printf("%d -> ", n.val)
	})
	// Output:
}

func ExampleBinaryTree_Search_find() {
	t := NewBinaryTree()

	t.Insert(32)
	t.Insert(55)
	t.Insert(79)
	t.Insert(10)
	t.Insert(1)
	t.Insert(19)
	t.Insert(23)
	t.Insert(16)

	n, err := t.Search(10)
	fmt.Println(n.val, err)

	// Output: 10 <nil>
}

func ExampleBinaryTree_Search_notfound() {
	t := NewBinaryTree()

	t.Insert(32)
	t.Insert(55)
	t.Insert(79)
	t.Insert(10)
	t.Insert(1)
	t.Insert(19)
	t.Insert(23)
	t.Insert(16)

	n, err := t.Search(22)
	fmt.Println(n, err)

	// Output: <nil> NotFound element 22
}

func ExampleBinaryTree_Search_empty() {
	t := NewBinaryTree()

	n, err := t.Search(22)
	fmt.Println(n, err)

	// Output: <nil> NotFound element 22
}
