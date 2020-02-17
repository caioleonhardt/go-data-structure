package stack

import (
	"strconv"
	"strings"
)

// Node represents the elements of the stack
type Node struct {
	val  int
	next *Node
}

// Stack data structure
type Stack struct {
	head *Node
	size int
}

// New returns an instance of Stack
func New() *Stack {
	return &Stack{}
}

// Push inserts elements into the stack
func (s *Stack) Push(v int) {
	if s.head == nil {
		s.head = &Node{v, nil}
		s.size++
		return
	}

	n := &Node{v, s.head}
	s.head = n
	s.size++
}

// Peek retrives the element from the top of the stack
// without remove from the stack
func (s *Stack) Peek() int {
	if s.head != nil {
		return s.head.val
	}
	return 0
}

// Pop retrives the element from the top of the stack.
// this method removes from the stack
func (s *Stack) Pop() int {
	if s.head == nil {
		return 0
	}

	res := s.head
	s.head = s.head.next
	s.size--

	return res.val
}

// Size of the stack
func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) String() string {
	if s.head == nil {
		return "nil"
	}

	str := []string{}
	s.transverse(func(n Node) {
		str = append(str, strconv.Itoa(n.val))
	})
	return strings.Join(str, "-> ")
}

func (s *Stack) transverse(f func(n Node)) {
	for cur := s.head; cur != nil; cur = cur.next {
		f(*cur)
	}
}
