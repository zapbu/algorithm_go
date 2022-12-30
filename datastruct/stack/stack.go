package stack

type Stack[T any] struct {
	head *Node[T]
}

type Node[T any] struct {
	Val  T
	next *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	new := &Node[T]{Val: val, next: s.head}
	s.head = new
}

func (s *Stack[T]) Pop() *Node[T] {
	popped := s.head
	s.head = s.head.next
	return popped
}

func (s Stack[T]) Top() *Node[T] {
	return s.head
}

func (s Stack[T]) IsEmpty() bool {
	return s.head == nil
}
