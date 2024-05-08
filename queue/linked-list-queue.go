package queue

type LinkedListQueue[T any] struct {
	head   *Node[T]
	length int
}

type Node[T any] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

func (q *LinkedListQueue[T]) Enqueue(data T) {
	node := &Node[T]{data: data}

	q.length++

	if q.head == nil {
		q.head = node
		return
	}

	node.prev = q.head
	q.head.next = node
	q.head = node
}

func (q *LinkedListQueue[T]) Dequeue() bool {
	if q.head == nil {
		return false
	}

	q.length--

	if q.head.prev == nil {
		q.head = nil
		return true
	}

	q.head.prev.next = nil
	q.head = q.head.prev

	return true
}

func (q *LinkedListQueue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *LinkedListQueue[T]) Peek() T {
	if q.head == nil {
		var emptyValue T
		return emptyValue
	}

	return q.head.data
}
