package linkedlist

import (
	"strconv"
	"strings"
)

// List contains all methods that an user
// can do to handle a collection.
type List interface {
	Add(n int)
	Get(i int) int
	Remove(i int)
	Size() int
}

// LinkedList data structure
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Node represents all element on linked list
type Node struct {
	data int
	next *Node
}

// NewLinkedList returns a new pointer to a linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add elements to the beginning of the list.
// default algorithm used is append.
func (l *LinkedList) Add(n int) {
	l.append(n)
}

func (l *LinkedList) append(n int) {
	l.size++
	if l.head == nil {
		head := &Node{n, nil}
		l.head = head
		l.tail = head
		return
	}

	node := &Node{n, nil}
	l.tail.next = node
	l.tail = node
}

// Get returns the n element in the collection.
func (l *LinkedList) Get(i int) int {
	if i == 0 {
		return l.head.data
	}

	count := 0
	var res int
	l.transverse(func(n Node) {
		if i == count {
			res = n.data
		}
		count++
	})

	return res
}

// Remove removes the n element in the collection.
func (l *LinkedList) Remove(i int) {
	var prev *Node
	var next *Node
	count := 0
	for cur := l.head; cur != nil; cur = cur.next {
		next = cur.next

		if count == i {
			// have only one node
			if l.head == l.tail {
				l.head = nil
				l.tail = nil
				l.size--
				return
			}

			if cur == l.head {

				l.head = next
				cur.next = nil
				l.size--
				return
			}

			if cur == l.tail {
				l.tail = prev
				prev.next = nil
				l.size--
				return
			}

			prev.next = next
			cur.next = nil
		}

		count++
		prev = cur
	}
}

// Size returns how many elements there are in the collection.
func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) transverse(f func(n Node)) {
	for cur := l.head; cur != nil; cur = cur.next {
		f(*cur)
	}
}

func (l *LinkedList) String() string {
	if l.head == nil {
		return "nil"
	}

	str := []string{}
	l.transverse(func(n Node) {
		str = append(str, n.String())
	})

	return strings.Join(str, "-> ")
}

func (n *Node) String() string {
	return strconv.Itoa(n.data)
}
