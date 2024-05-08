package list

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	data       T
	next, prev *Node[T]
}

func (l *List[T]) InsertHead(data T) {
	newHead := l.newNode(data, l.head, nil)

	if l.head != nil {
		l.head.prev = newHead
	} else {
		l.tail = newHead
	}

	l.head = newHead
}

func (l *List[T]) InsertTail(data T) {
	newTail := l.newNode(data, nil, l.tail)

	if l.tail != nil {
		l.tail.next = newTail
	} else {
		l.head = newTail
	}

	l.tail = newTail
}

func (l *List[T]) InsertAt(index int, data T) bool {
	element := l.getAt(index)
	if element == nil {
		return false
	}

	node := l.newNode(data, element, element.prev)
	element.prev.next = node
	element.prev = node

	return true
}

func (l *List[T]) RemoveHead() bool {
	return l.removeHead()
}

func (l *List[T]) RemoveTail() bool {
	return l.removeTail()
}

func (l *List[T]) RemoveAt(index int) bool {
	element := l.getAt(index)
	if element == nil {
		return false
	}

	if element.prev == nil {
		return l.removeHead()
	}

	if element.next == nil {
		return l.removeTail()
	}

	element.prev.next = element.next

	return true
}

func (l *List[T]) Iterate(fn func(node *Node[T])) {
	current := l.head

	for current != nil {
		fn(current)
		current = current.next
	}
}

func (l *List[T]) Get(index int) *Node[T] {
	return l.getAt(index)
}

func (l *List[T]) removeHead() bool {
	if l.head == nil {
		return false
	}

	l.head = l.head.next
	l.head.prev = nil

	return true
}

func (l *List[T]) removeTail() bool {
	if l.tail == nil {
		return false
	}

	l.tail = l.tail.prev
	l.tail.next = nil

	return true
}

func (l *List[T]) getAt(index int) *Node[T] {
	element := l.head

	for i := 0; i != index; i++ {
		if element == nil {
			return nil
		}

		element = element.next
	}

	return element
}

func (l *List[T]) newNode(
	data T,
	next, prev *Node[T],
) *Node[T] {
	return &Node[T]{
		data: data,
		next: next,
		prev: prev,
	}
}
