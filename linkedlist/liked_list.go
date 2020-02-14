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
