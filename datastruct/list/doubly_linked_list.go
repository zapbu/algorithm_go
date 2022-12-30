package list

import (
	"fmt"
)

type DoublyLinkedList[T any] struct {
	head *DNode[T]
	tail *DNode[T]
	len  int
}

type DNode[T any] struct {
	Val  T
	prev *DNode[T]
	next *DNode[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	head := &DNode[T]{}
	tail := &DNode[T]{}
	head.next, tail.prev = tail, head

	return &DoublyLinkedList[T]{
		head: head,
		tail: tail,
		len:  0,
	}
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.len
}

func (l *DoublyLinkedList[T]) Get(index int) (*DNode[T], error) {
	var cursor *DNode[T]

	// 오류 발생 케이스
	// Error Cases
	if l.len == 0 {
		// 1. 빈 List를 조회할 경우
		// 1. Access to Empty List
		return nil, errEmptyListAccess // All the returning errors defined at ./errors.go
	} else if l.len-1 < index {
		// 2. index is bigger than Last Index
		// 2. index가 Last Index보다 큰 경우
		return nil, fmt.Errorf("%v, last index is %d", errOutOfBounds.Error(), l.len-1)
	} else if index < 0 {
		// 3. index is negative
		// 3. index가 음수인 경우
		return nil, errNegativeIndex
	}

	var currentIndex int

	if index < l.len/2 {
		// 1. index가 length의 전반부를 가리키는 경우
		// currentIndex는 -1부터 시작하여 index가 될 때까지 증가
		// 1회 이동 시 비로소 index 0에 위치한 Node를 가리키게 되므로 -1부터 시작
		currentIndex = -1
		cursor = l.head

		for currentIndex < index {
			cursor = cursor.next
			currentIndex++
		}
	} else {
		// 2. index가 length의 후반부를 가리키는 경우
		// currentIndex는 Last Index부터 시작하여 index가 될 때까지 감소
		// 1회 이동 시 비로소 Last index에 위치한 Node를 가리키게 되므로 l.length(Last Index + 1)부터 시작
		currentIndex = l.len
		cursor = l.tail

		for currentIndex > index {
			cursor = cursor.prev
			currentIndex--
		}
	}

	return cursor, nil
}

func (l *DoublyLinkedList[T]) InsertAt(index int, val T) error {
	var (
		prev *DNode[T]
		err  error
	)

	// 1. length is 0, prev will be `head`
	if index == 0 {
		prev = l.head
	} else {
		// 2. insert at Last Index + 1, prev will be Last Node
		prev, err = l.Get(index - 1)
		if err != nil {
			return err
		}
	}

	new, next := &DNode[T]{Val: val}, prev.next
	new.prev, new.next = prev, next
	prev.next, next.prev = new, new
	l.len++

	return nil
}

func (l *DoublyLinkedList[T]) InsertFirst(val T) error {
	return l.InsertAt(0, val)
}

func (l *DoublyLinkedList[T]) InsertLast(val T) error {
	return l.InsertAt(l.len, val)
}

func (l *DoublyLinkedList[T]) DeleteAt(index int) error {
	var (
		prev *DNode[T]
		err  error
	)

	if index == 0 {
		prev = l.head
	} else {
		prev, err = l.Get(index - 1)
		if err != nil {
			return err
		}
	}

	next := prev.next.next
	prev.next, next.prev = next, prev
	l.len--

	return nil
}

func (l *DoublyLinkedList[T]) DeleteFirst() error {
	return l.DeleteAt(0)
}

func (l *DoublyLinkedList[T]) DeleteLast() error {
	return l.DeleteAt(l.len - 1)
}
