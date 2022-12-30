package list

import (
	"fmt"
)

type SinglyLinkedList[T any] struct {
	head *SNode[T]
	tail *SNode[T]
	len  int
}

// Each Elements of Linked List normally called a `Node`.
//
// prefix 'S' is added but it's just for avoiding duplicate type error.
// because I'll also make DoublyLinkedList in the same package.(Go doesn't allow duplicate type names in the same package.)
// and both of'em have `Node` for representing their Elements.
// but it's a bit different.(DoublyLinkedList also has `prev` field additionally.)
// So I distinguish them with the prefix `S` and `D`
type SNode[T any] struct {
	Val  T
	next *SNode[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	head := &SNode[T]{}
	tail := &SNode[T]{}
	head.next = tail

	return &SinglyLinkedList[T]{
		head: head,
		tail: tail,
		len:  0,
	}
}

func (l *SinglyLinkedList[T]) Len() int {
	return l.len
}

func (l *SinglyLinkedList[T]) Get(index int) (*SNode[T], error) {
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

	// A pointer to search Node at input index
	// 입력된 index에 위치한 Node를 찾기 위한 포인터
	cursor := l.head

	// index를 1씩 감소시키면서 총 index + 1회만큼 next로 이동
	// index + 1인 이유는 첫 Node가 Dummy Node인 l.head이기 때문
	// 1회 이동 해야 cursor가 index 0에 위치한 Node를 가리키게 됨
	for index > -1 {
		cursor = cursor.next
		index--
	}

	return cursor, nil
}

func (l *SinglyLinkedList[T]) InsertAt(index int, val T) error {
	var (
		prev *SNode[T]
		err  error
	)

	if index == 0 {
		// 1. index is 0, prev will be `head`
		prev = l.head
	} else {
		// 2. insert at Last Index + 1, prev will be Last Node
		prev, err = l.Get(index - 1)
		if err != nil {
			return err
		}
	}

	new, next := &SNode[T]{Val: val}, prev.next
	prev.next, new.next = new, next
	l.len++

	return nil
}

func (l *SinglyLinkedList[T]) InsertFirst(val T) error {
	return l.InsertAt(0, val)
}

func (l *SinglyLinkedList[T]) InsertLast(val T) error {
	return l.InsertAt(l.len, val)
}

func (l *SinglyLinkedList[T]) DeleteAt(index int) error {
	var (
		prev *SNode[T]
		err  error
	)

	if index == 0 {
		// Case 1. delete at index 0, set prev to `head` Node
		prev = l.head
	} else {
		// Case 2. else, get Node from `.Get` Method
		prev, err = l.Get(index - 1)
		if err != nil {
			return err
		}
	}

	prev.next = prev.next.next
	l.len--

	return nil
}

func (l *SinglyLinkedList[T]) DeleteFirst() error {
	return l.DeleteAt(0)
}

func (l *SinglyLinkedList[T]) DeleteLast() error {
	return l.DeleteAt(l.len - 1)
}
