package singlylinkedlist

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	head   *Node[T]
	length int
}

func (l *SinglyLinkedList[T]) Len() int { return l.length }

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head:   nil,
		length: 0,
	}
}

func (l *SinglyLinkedList[T]) Prepend(item T) {
	node := &Node[T]{Value: item, Next: nil}
	l.length += 1

	if l.head == nil {
		l.head = node
	} else {
		node.Next = l.head
		l.head = node
	}
}

func (l *SinglyLinkedList[T]) InsertAt(item T, idx int) {
	node := &Node[T]{Value: item, Next: nil}
	l.length += 1

	if l.head == nil {
		l.head = node
	} else {
		var prev *Node[T]
		current := l.head
		for i := 0; current != nil && i < l.length; prev, current, i = current, current.Next, i+1 {
			if i == idx {
				if prev != nil {
					prev.Next = node
					node.Next = current
				} else {
					node.Next = l.head
					l.head = node
				}
				return
			}
		}
		if prev != nil {
			prev.Next = node
		}
	}
}

func (l *SinglyLinkedList[T]) Append(item T) {
	node := &Node[T]{Value: item, Next: nil}

	if l.head == nil {
		l.head = node
	} else {
		current := l.head

		for ; current.Next != nil; current = current.Next {
		}

		current.Next = node
	}

	l.length += 1
}

func (l *SinglyLinkedList[T]) Remove(item T) (T, bool) {
	var prev *Node[T]
	current := l.head
	for ; current != nil; prev, current = current, current.Next {
		if current.Value == item {
			l.length -= 1
			if prev != nil {
				prev.Next = current.Next
			} else {
				l.head = current.Next
			}
			current.Next = nil
			return current.Value, true
		}
	}

	var noop T
	return noop, false
}

func (l *SinglyLinkedList[T]) Get(idx int) (T, bool) {
	for current, i := l.head, 0; current != nil && i < l.length; current, i = current.Next, i+1 {
		if i == idx {
			return current.Value, true
		}
	}

	var noop T
	return noop, false
}

func (l *SinglyLinkedList[T]) RemoveAt(idx int) (T, bool) {
	var prev *Node[T]
	current := l.head
	for i := 0; current != nil && i < l.length; prev, current, i = current, current.Next, i+1 {
		if i == idx {
			l.length -= 1
			if prev != nil {
				prev.Next = current.Next
			} else {
				l.head = current.Next
			}
			current.Next = nil
			return current.Value, true
		}
	}

	var noop T
	return noop, false
}
