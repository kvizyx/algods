package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListQueue_Enqueue(t *testing.T) {
	q := LinkedListQueue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if actualValue := q.head; actualValue == nil {
		t.Errorf("head shouldn't be nil")
	}

	if actualValue := q.length; actualValue != 3 {
		t.Errorf("queue length is %d, expected %d", actualValue, 3)
	}

	if actualValue := q.head.data; actualValue != 3 {
		t.Errorf("head data is %d, expected %d", actualValue, 3)
	}

	if actualValue := q.head.next; actualValue != nil {
		t.Errorf("head next value is not nil, expected nil")
	}

	if actualValue := q.head.prev.data; actualValue != 2 {
		t.Errorf("head prev element data is %d, expected %d", actualValue, 2)
	}
}

func TestLinkedListQueue_Dequeue(t *testing.T) {
	q := LinkedListQueue[int]{}

	assert.NotPanics(t, func() { q.Dequeue() }, "dequeue on empty queue panics")

	q.Enqueue(1)
	q.Enqueue(2)
	q.Dequeue()

	if actualValue := q.length; actualValue != 1 {
		t.Errorf("queue length is %d, expected %d", actualValue, 1)
	}

	if actualValue := q.head.data; actualValue != 1 {
		t.Errorf("head data is %d, expected %d", actualValue, 1)
	}

	q.Dequeue()
	q.Dequeue()

	if actualValue := q.head; actualValue != nil {
		t.Errorf("expected head to be nil")
	}

	if actualValue := q.length; actualValue != 0 {
		t.Errorf("queue length is %d, expected %d", actualValue, 0)
	}
}

func TestLinkedListQueue_Peek(t *testing.T) {
	q := LinkedListQueue[int]{}

	if actualValue := q.Peek(); actualValue != 0 {
		t.Errorf("new queue head data is %d, expected %d", actualValue, 0)
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if actualValue := q.Peek(); actualValue != 2 {
		t.Errorf("head data is %d, expected %d", actualValue, 2)
	}

	q.Dequeue()

	if actualValue := q.Peek(); actualValue != 1 {
		t.Errorf("head data after dequeue is %d, expected %d", actualValue, 1)
	}
}

func TestLinkedListQueue_IsEmpty(t *testing.T) {
	q := LinkedListQueue[int]{}

	if actualValue := q.IsEmpty(); actualValue != true {
		t.Errorf("expected new queue length length to be 0")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if actualValue := q.IsEmpty(); actualValue != false {
		t.Errorf("expected queue to be not empty after enqueues")
	}

	q.Dequeue()
	q.Dequeue()

	if actualValue := q.IsEmpty(); actualValue != true {
		t.Errorf("expected queue length to be 0 after dequeues")
	}
}
