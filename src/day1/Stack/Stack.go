package stack

type Node[T any] struct {
	Value T
	Prev  *Node[T]
}

type Stack[T any] struct {
	Length int
	Head   *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Length: 0,
		Head:   nil,
	}
}

func (s *Stack[T]) Push(item T) {
	s.Length += 1
	node := &Node[T]{Value: item, Prev: nil}

	if s.Head != nil {
		node.Prev = s.Head
	}

	s.Head = node
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Head == nil {
		var noop T
		return noop, false
	}

	s.Length -= 1
	node := s.Head
	s.Head = node.Prev
	node.Prev = nil

	return node.Value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.Head == nil {
		var noop T
		return noop, false
	}

	return s.Head.Value, true
}
