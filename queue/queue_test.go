package queue

import "fmt"

func ExampleQueue_Enqueue() {
	q := New()
	fmt.Println(q)

	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println(q)

	// Output:
	// nil
	// 10-> 20
}

func ExampleQueue_Dequeue() {
	q := New()
	fmt.Println(q)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q)

	// Output:
	// nil
	// 10-> 20-> 30
	// 10
	// 20-> 30
}

func ExampleQueue_Dequeue_letempty() {
	q := New()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	// Output:
	// 10
	// 20
	// 30
	// 0
}

func ExampleQueue_Peek() {
	q := New()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q)

	fmt.Println(q.Peek())
	fmt.Println(q.Peek())

	// Output:
	// 10-> 20-> 30
	// 10
	// 10
}

func ExampleQueue_Size() {
	q := New()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q.Size())

	// Output:
	// 3
}
