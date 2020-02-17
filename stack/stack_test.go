package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		{"empty stack", &Stack{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleStack_Push() {
	s := New()

	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s)
	// Output: 30-> 20-> 10
}

func ExampleStack_Peek() {
	s := New()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println(s.Peek())
	fmt.Println(s.Peek())
	// Output:
	// 30
	// 30
}

func ExampleStack_Pop() {
	s := New()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	// Output:
	// 30
	// 20
}

func ExampleStack_String_nil() {
	s := New()

	fmt.Println(s)
	// Output:
	// nil
}

func ExampleStack_Size() {
	s := New()
	fmt.Println(s.Size())
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Size())
	s.Pop()
	fmt.Println(s.Size())

	// Output:
	// 0
	// 3
	// 2
}
