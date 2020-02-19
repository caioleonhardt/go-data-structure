package tree

import (
	"fmt"
	"testing"
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

	t.Walk(func(n *Node) {
		fmt.Printf("%d -> ", n.val)
	})
	// Output: 1 -> 10 -> 16 -> 19 -> 23 -> 32 -> 55 -> 79 ->
}

func ExampleBinaryTree_Walk_empty() {
	t := NewBinaryTree()

	t.Walk(func(n *Node) {
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

func TestBinaryTree_Delete_empty(t *testing.T) {
	err := NewBinaryTree().Delete(1)

	if err == nil || err.Error() != "EmptyTree" {
		t.Errorf("Delete() = EmptyTree, got= %v", err)
	}
}

func TestBinaryTree_Delete_leftLeaf(t *testing.T) {
	tr := NewBinaryTree()
	tr.Insert(32)
	tr.Insert(10)
	tr.Insert(1)
	tr.Insert(19)

	// left
	if gotErr := tr.Delete(1); gotErr != nil {
		t.Errorf("Delete() = nil, got= %v", gotErr)
	}

	want := "10 -> 19 -> 32"
	if gotStr := tr.String(); gotStr != want {
		t.Errorf("String() = %s, got= %v", want, gotStr)

	}
}

func TestBinaryTree_Delete_rightLeaf(t *testing.T) {
	tr := NewBinaryTree()
	tr.Insert(32)
	tr.Insert(10)
	tr.Insert(1)
	tr.Insert(19)

	//right
	if gotErr := tr.Delete(19); gotErr != nil {
		t.Errorf("Delete() = nil, got= %v", gotErr)
	}

	want := "1 -> 10 -> 32"
	if gotStr := tr.String(); gotStr != want {
		t.Errorf("String() = %s, got= %v", want, gotStr)

	}
}

func TestBinaryTree_Delete_left_LL(t *testing.T) {
	tr := NewBinaryTree()
	tr.Insert(32)
	tr.Insert(10)
	tr.Insert(1)

	fmt.Println("before:", tr)
	//right
	if gotErr := tr.Delete(10); gotErr != nil {
		t.Errorf("Delete() = nil, got= %v", gotErr)
	}

	want := "1 -> 32"
	if gotStr := tr.String(); gotStr != want {
		t.Errorf("String() = %s, got= %v", want, gotStr)

	}
	fmt.Println("after:", tr)
}

func TestBinaryTree_Delete_left_RL(t *testing.T) {
	tr := NewBinaryTree()
	tr.Insert(32)
	tr.Insert(55)
	tr.Insert(44)

	fmt.Println("before:", tr)
	//right
	if gotErr := tr.Delete(55); gotErr != nil {
		t.Errorf("Delete() = nil, got= %v", gotErr)
	}

	want := "32 -> 44"
	if gotStr := tr.String(); gotStr != want {
		t.Errorf("String() = %s, got= %v", want, gotStr)
	}
	fmt.Println("after:", tr)
}
