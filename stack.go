package caravan

type Stack[T any] struct {
	top    *stackNode[T]
	length int
}

type stackNode[T any] struct {
	value T
	prev  *stackNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

func (s *Stack[T]) Len() int {
	return s.length
}

func (s *Stack[T]) Peek() T {
	if s.length <= 0 {
		var empty T
		return empty
	}

	return s.top.value
}

func (s *Stack[T]) Pop() T {
	if s.length <= 0 {
		var empty T
		return empty
	}

	n := s.top
	s.top = n.prev
	s.length--

	return n.value
}

func (s *Stack[T]) Push(value T) {
	n := &stackNode[T]{value, s.top}
	s.top = n
	s.length++
}
