package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// INSERT
func TestInsertLastWorksDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestInsertFirstWorksDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	require.Equal(t, 0, list.Len())
	firstVal, secondVal := 1, 2
	require.NoError(t, list.InsertFirst(firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertFirst(secondVal))
	require.Equal(t, 2, list.Len())

	first, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, secondVal, first.Val)

	second, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, firstVal, second.Val)
}

func TestInsertAtCanAddToFirstDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestInsertAtCanAddToLastDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestInsertAtCanAddToMidDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDeleteFirstWorksDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDeleteLastWorksDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDeleteAtCanDeleteFirstDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDeleteAtCanDeleteLastDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestDeleteAtCanDeleteMidDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	firstVal, secondVal, thirdVal := 1, 2, 3
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)
	list.InsertLast(thirdVal)
	require.Equal(t, 3, list.Len())

	require.NoError(t, list.DeleteAt(1))
	midNode, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, thirdVal, midNode.Val)
	require.Equal(t, 2, list.Len())
}

// CREATE
func TestCreateIntListDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestCreateStringListDLL(t *testing.T) {
	list := NewDoublyLinkedList[string]()
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
func TestGetInIntListDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
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

func TestGetInStringListDLL(t *testing.T) {
	list := NewDoublyLinkedList[string]()
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
func TestGetFromEmptyListDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	node, err := list.Get(0)
	require.Nil(t, node)
	require.Error(t, err)
	require.ErrorIs(t, err, errEmptyListAccess)
}

func TestGetAtOutOfBoundsDLL(t *testing.T) {
	index := 999
	list := NewDoublyLinkedList[int]()
	list.InsertLast(1)
	node, err := list.Get(index)
	require.Nil(t, node)
	require.Error(t, err)

	t.Log(err)
	// require.EqualError(t, err, fmt.Sprintf("%v, ", errOutOfBounds, index))
}

func TestGetAtNegativeIndexDLL(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.InsertLast(1)

	node, err := list.Get(-1)
	require.Nil(t, node)
	require.Error(t, err)
	require.ErrorIs(t, err, errNegativeIndex)
}

func TestInsertAtOutOfBoundsDLL(t *testing.T) {
	index := 1
	list := NewDoublyLinkedList[int]()
	list.InsertAt(999999, 1)
	node, err := list.Get(index)
	require.Nil(t, node)
	require.Error(t, err)
	// require.EqualError(t, err.Error(), fmt.Errorf("%v, last index is %d", errOutOfBounds.Error(), 2).Error())
}

func TestDeleteAtOutOfBoundsDLL(t *testing.T) {
	index := 1
	list := NewDoublyLinkedList[int]()
	list.DeleteAt(999999)
	node, err := list.Get(index)
	require.Nil(t, node)
	require.Error(t, err)
	// require.EqualError(t, err.Error(), fmt.Errorf("%v, last index is %d", errOutOfBounds.Error(), 2).Error())
}
