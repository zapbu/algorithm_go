package list

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type SinglyLinkedList[T constraints.Ordered] struct {
	head   *SNode[T]
	tail   *SNode[T]
	length int
}

type SNode[T constraints.Ordered] struct {
	Val  T
	next *SNode[T]
}

func NewSinglyLinkedList[T constraints.Ordered]() *SinglyLinkedList[T] {
	head := &SNode[T]{}
	tail := &SNode[T]{}
	head.next = tail

	return &SinglyLinkedList[T]{
		head:   head,
		tail:   tail,
		length: 0,
	}
}

func (l *SinglyLinkedList[T]) Len() int {
	return l.length
}

func (l *SinglyLinkedList[T]) Get(index int) (*SNode[T], error) {
	if l.length == 0 {
		return nil, fmt.Errorf("empty list")
	}
	if l.length-1 < index {
		return nil, fmt.Errorf("allowed to access to index at up to %d", l.length-1)
	}

	cursor := l.head

	for index > -1 {
		cursor = cursor.next
		index--
	}

	return cursor, nil
}

func (l *SinglyLinkedList[T]) InsertAt(index int, val T) error {
	prev, err := l.Get(index - 1)
	if err != nil {
		if err.Error() == "empty list" {
			prev = l.head
		} else {
			return fmt.Errorf("allowed to insert at up to %d", l.length)
		}
	}

	newNode := &SNode[T]{Val: val}

	next := prev.next
	prev.next = newNode
	newNode.next = next
	l.length++

	return nil
}

func (l *SinglyLinkedList[T]) InsertFirst(val T) error {
	return l.InsertAt(0, val)
}

func (l *SinglyLinkedList[T]) InsertLast(val T) error {
	return l.InsertAt(l.length, val)
}

func (l *SinglyLinkedList[T]) DeleteAt(index int) error {
	prev, err := l.Get(index - 1)
	if err != nil {
		return err
	}

	target := prev.next
	prev.next = target.next
	l.length--

	return nil
}

func (l *SinglyLinkedList[T]) DeleteFirst() error {
	return l.DeleteAt(0)
}

func (l *SinglyLinkedList[T]) DeleteLast() error {
	return l.DeleteAt(l.length - 1)
}
