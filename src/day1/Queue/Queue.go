package queue

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Length int
	Head   *Node[T]
	Tail   *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Length: 0,
		Head:   nil,
		Tail:   nil,
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.Length += 1
	node := &Node[T]{Value: item, Next: nil}

	if q.Head == nil && q.Tail == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func (q *Queue[T]) Deque() (T, bool) {
	if q.Head == nil {
		var noop T
		return noop, false
	}

	q.Length -= 1
	node := q.Head
	q.Head = node.Next
	node.Next = nil

    if q.Length == 0 {
        q.Tail = nil
    }

	return node.Value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.Head == nil {
		var noop T
		return noop, false
	}

	return q.Head.Value, true
}
