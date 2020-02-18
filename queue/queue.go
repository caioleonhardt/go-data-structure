package queue

import (
	"strconv"
	"strings"
)

// Queue data structure
type Queue struct {
	head *Node
	tail *Node
	size int
}

// Node represents all element on linked list
type Node struct {
	data int
	next *Node
}

// New instance of queue data structure
func New() *Queue {
	return &Queue{}
}

// Enqueue put the elements in end of the list
func (q *Queue) Enqueue(v int) {
	q.size++

	if q.head == nil {
		q.head = &Node{v, nil}
		q.tail = q.head
		return
	}

	node := &Node{v, nil}
	q.tail.next = node
	q.tail = node
}

// Dequeue retrieves the element of the beggining
// the list and removes it
func (q *Queue) Dequeue() int {
	q.size--

	if q.head == nil {
		return 0
	}

	old := q.head
	q.head = q.head.next

	return old.data
}

// Peek retrieves the element from the beginning
// without removing it of the list
func (q *Queue) Peek() int {
	if q.head == nil {
		return 0
	}
	return q.head.data
}

// Size of the list
func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) String() string {
	if q.head == nil {
		return "nil"
	}

	var str []string

	for cur := q.head; cur != nil; cur = cur.next {
		str = append(str, strconv.Itoa(cur.data))
	}
	return strings.Join(str, "-> ")
}
