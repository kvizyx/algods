// This package implements LIFO stack based on singly-linked list

package stack

type Stack[T any] struct {
	top *Entry[T]
}

type Entry[T any] struct {
	data T
	next *Entry[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		top: nil,
	}
}

func (s *Stack[T]) Push(data T) {
	newEntry := &Entry[T]{
		data: data,
		next: s.top,
	}

	s.top = newEntry
}

func (s *Stack[T]) Pop() (t T) {
	if s.top == nil {
		return t
	}

	removed := s.top.data
	s.top = s.top.next

	return removed
}

func (s *Stack[T]) Peek() (t T) {
	if s.top == nil {
		return t
	}

	return s.top.data
}

func (s *Stack[T]) Empty() bool {
	return s.top == nil
}
