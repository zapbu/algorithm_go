package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// INSERT
func TestInsertLastWorksSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	require.Equal(t, 0, list.Len())
	firstVal, secondVal := 1, 2
	require.NoError(t, list.InsertLast(firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertLast(secondVal))
	require.Equal(t, 2, list.Len())

	first, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, firstVal, first.Val)

	second, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, secondVal, second.Val)
}

func TestInsertAtCanAddToFirstSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	require.NoError(t, list.InsertAt(0, firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertAt(0, secondVal))
	require.Equal(t, 2, list.Len())

	first, err := list.Get(0)
	require.NoError(t, err)
	second, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, secondVal, first.Val)
	require.Equal(t, firstVal, second.Val)
}

func TestInsertAtCanAddToLastSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	require.NoError(t, list.InsertAt(0, firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertAt(1, secondVal))
	require.Equal(t, 2, list.Len())

	first, err := list.Get(0)
	require.NoError(t, err)
	second, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, firstVal, first.Val)
	require.Equal(t, secondVal, second.Val)
}

func TestInsertAtCanAddToMidSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal, thirdVal := 1, 2, 3
	require.NoError(t, list.InsertAt(0, firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertAt(1, secondVal))
	require.Equal(t, 2, list.Len())
	require.NoError(t, list.InsertAt(1, thirdVal))
	require.Equal(t, 3, list.Len())

	first, err := list.Get(0)
	require.NoError(t, err)
	second, err := list.Get(1)
	require.NoError(t, err)
	third, err := list.Get(2)
	require.NoError(t, err)
	require.Equal(t, firstVal, first.Val)
	require.Equal(t, thirdVal, second.Val)
	require.Equal(t, secondVal, third.Val)
}

// DELETE

func TestDeleteFirstWorksSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertAt(0, firstVal)
	list.InsertAt(1, secondVal)

	require.NoError(t, list.DeleteFirst())
	require.Equal(t, 1, list.Len())

	first, _ := list.Get(0)
	require.Equal(t, secondVal, first.Val)

	second, err := list.Get(1)
	require.Nil(t, second)
	require.Error(t, err)
}

func TestDeleteLastWorksSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertAt(0, firstVal)
	list.InsertAt(1, secondVal)
	require.Equal(t, 2, list.Len())

	require.NoError(t, list.DeleteLast())
	require.Equal(t, 1, list.Len())

	first, _ := list.Get(0)
	require.Equal(t, firstVal, first.Val)

	second, err := list.Get(1)
	require.Nil(t, second)
	require.Error(t, err)
}

func TestDeleteAtCanDeleteFirstSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)
	require.Equal(t, 2, list.Len())

	require.NoError(t, list.DeleteAt(0))
	first, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, secondVal, first.Val)
	require.Equal(t, 1, list.Len())
}

func TestDeleteAtCanDeleteLastSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)
	require.Equal(t, 2, list.Len())

	require.NoError(t, list.DeleteAt(1))
	first, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, firstVal, first.Val)
	require.Equal(t, 1, list.Len())
}

func TestDeleteAtCanDeleteMidSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal, thirdVal := 1, 2, 3
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)
	list.InsertLast(thirdVal)
	require.Equal(t, 3, list.Len())

	require.NoError(t, list.DeleteAt(1))
	mid, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, thirdVal, mid.Val)
	require.Equal(t, 2, list.Len())
}

// CREATE
func TestCreateIntList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal, thirdVal := 1, 2, 3
	require.Equal(t, 0, list.Len())
	list.InsertLast(firstVal)
	require.Equal(t, 1, list.Len())
	list.InsertLast(secondVal)
	require.Equal(t, 2, list.Len())
	list.InsertLast(thirdVal)
	require.Equal(t, 3, list.Len())

	t.Logf("Create int List Result: %+v", list)
}

func TestCreateStringListSLL(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	firstVal, secondVal, thirdVal := "a", "b", "c"
	require.Equal(t, 0, list.Len())
	list.InsertLast(firstVal)
	require.Equal(t, 1, list.Len())
	list.InsertLast(secondVal)
	require.Equal(t, 2, list.Len())
	list.InsertLast(thirdVal)
	require.Equal(t, 3, list.Len())

	t.Logf("Create string List Result: %+v", list)
}

// GET
func TestGetInIntListSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)

	first, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, firstVal, first.Val)

	second, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, secondVal, second.Val)
}

func TestGetInStringListSLL(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	firstVal, secondVal := "a", "b"
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)

	first, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, firstVal, first.Val)

	second, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, secondVal, second.Val)
}

// Error Cases
func TestGetFromEmptyListSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	node, err := list.Get(0)
	require.Nil(t, node)
	require.Error(t, err)
	require.EqualError(t, errEmptyListAccess, err.Error())
}

func TestGetAtOutOfBoundsSLL(t *testing.T) {
	index := 1
	list := NewSinglyLinkedList[int]()
	list.InsertLast(9999999)
	node, err := list.Get(index)
	require.Nil(t, node)
	require.Error(t, err)
	t.Logf("ok: %v", err.Error())
	// require.EqualError(t, err.Error(), fmt.Errorf("%v, last index is %d", errOutOfBounds.Error(), 2).Error())
}

func TestGetAtNegativeIndexSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.InsertFirst(1)
	node, err := list.Get(-1)
	require.Nil(t, node)
	require.Error(t, err)
	require.EqualError(t, errNegativeIndex, err.Error())
}

func TestInsertAtOutOfBoundsSLL(t *testing.T) {
	index := 1
	list := NewSinglyLinkedList[int]()
	list.InsertAt(999999, 1)
	node, err := list.Get(index)
	require.Nil(t, node)
	require.Error(t, err)
	// require.EqualError(t, err.Error(), fmt.Errorf("%v, last index is %d", errOutOfBounds.Error(), 2).Error())
}

func TestDeleteAtOutOfBoundsSLL(t *testing.T) {
	index := 1
	list := NewSinglyLinkedList[int]()
	list.DeleteAt(999999)
	node, err := list.Get(index)
	require.Nil(t, node)
	require.Error(t, err)
	// require.EqualError(t, err.Error(), fmt.Errorf("%v, last index is %d", errOutOfBounds.Error(), 2).Error())
}
