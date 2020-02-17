package linkedlist

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	l := NewLinkedList()

	tests := []struct {
		name string
		n    int
		want int
	}{
		{"", 1, 1},
		{"", 10, 2},
		{"", 20, 3},
		{"", 30, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Add(tt.n)
			if l.size != tt.want {
				t.Errorf("list.size = %v, want %v", l.size, tt.want)
			}
		})
	}
}

func ExampleLinkedList_String() {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	fmt.Println(l)
	//Output: 1-> 2-> 3-> 4
}

func ExampleLinkedList_Get() {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	fmt.Println(l.Get(0))
	fmt.Println(l.Get(3))
	// Output:
	// 1
	// 4
}

func TestRemoveHeadTail(t *testing.T) {
	l := NewLinkedList()
	l.Add(1)
	l.Remove(0)
	if l.size != 0 {
		t.Errorf("list.size = %v, want %v", l.size, 0)
	}

	l.Add(1)
	l.Add(2)
	l.Remove(0)
	if l.size != 1 {
		t.Errorf("list.size = %v, want %v", l.size, 1)
	}

	if l.Get(0) != 2 {
		t.Errorf("list.size = %v, want %v", l.size, 1)
	}
}

func TestRemoveHead(t *testing.T) {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Remove(0)

	want := "2"
	got := l.String()
	if got != want {
		t.Errorf("list = %v, want %v", got, want)
	}
}

func TestRemoveTail(t *testing.T) {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Remove(1)

	want := "1"
	got := l.String()
	if got != want {
		t.Errorf("list = %v, want %v", got, want)
	}
}

func TestRemoveMiddle(t *testing.T) {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	l.Remove(1)
	want := "1-> 3"
	got := l.String()
	if got != want {
		t.Errorf("list = %v, want %v", got, want)
	}
}

func ExampleLinkedList_Size() {
	l := NewLinkedList()
	fmt.Println(l.Size())

	l.Add(1)
	l.Add(2)
	fmt.Println(l.Size())

	l.Add(3)
	l.Add(4)
	fmt.Println(l.Size())

	l.Remove(0)
	l.Remove(0)
	l.Remove(0)
	l.Remove(0)
	fmt.Println(l.Size())

	// Output:
	// 0
	// 2
	// 4
	// 0
}
