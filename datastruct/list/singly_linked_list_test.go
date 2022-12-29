package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// INSERT
func TestInsertLastWorks(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	require.Equal(t, 0, list.Len())
	firstVal, secondVal := 1, 2
	require.NoError(t, list.InsertLast(firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertLast(secondVal))
	require.Equal(t, 2, list.Len())

	firstNode, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, firstVal, firstNode.Val)

	secondNode, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, secondVal, secondNode.Val)
}

func TestInsertAtCanAddToFirst(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	require.NoError(t, list.InsertAt(0, firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertAt(0, secondVal))
	require.Equal(t, 2, list.Len())

	firstNode, err := list.Get(0)
	require.NoError(t, err)
	secondNode, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, secondVal, firstNode.Val)
	require.Equal(t, firstVal, secondNode.Val)
}

func TestInsertAtCanAddToLast(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	require.NoError(t, list.InsertAt(0, firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertAt(1, secondVal))
	require.Equal(t, 2, list.Len())

	firstNode, err := list.Get(0)
	require.NoError(t, err)
	secondNode, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, firstVal, firstNode.Val)
	require.Equal(t, secondVal, secondNode.Val)
}

func TestInsertAtCanAddToMid(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal, thirdVal := 1, 2, 3
	require.NoError(t, list.InsertAt(0, firstVal))
	require.Equal(t, 1, list.Len())
	require.NoError(t, list.InsertAt(1, secondVal))
	require.Equal(t, 2, list.Len())
	require.NoError(t, list.InsertAt(1, thirdVal))
	require.Equal(t, 3, list.Len())

	firstNode, err := list.Get(0)
	require.NoError(t, err)
	secondNode, err := list.Get(1)
	require.NoError(t, err)
	thirdNode, err := list.Get(2)
	require.NoError(t, err)
	require.Equal(t, firstVal, firstNode.Val)
	require.Equal(t, thirdVal, secondNode.Val)
	require.Equal(t, secondVal, thirdNode.Val)
}

// DELETE

func TestDeleteFirstWorks(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertAt(0, firstVal)
	list.InsertAt(1, secondVal)

	require.NoError(t, list.DeleteFirst())
	require.Equal(t, 1, list.Len())

	firstNode, _ := list.Get(0)
	require.Equal(t, secondVal, firstNode.Val)

	secondNode, err := list.Get(1)
	require.Nil(t, secondNode)
	require.Error(t, err)
}

func TestDeleteLastWorks(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertAt(0, firstVal)
	list.InsertAt(1, secondVal)
	require.Equal(t, 2, list.Len())

	require.NoError(t, list.DeleteLast())
	require.Equal(t, 1, list.Len())

	firstNode, _ := list.Get(0)
	require.Equal(t, firstVal, firstNode.Val)

	secondNode, err := list.Get(1)
	require.Nil(t, secondNode)
	require.Error(t, err)
}

func TestDeleteAtCanDeleteFirst(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)
	require.Equal(t, 2, list.Len())

	require.NoError(t, list.DeleteAt(0))
	firstNode, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, secondVal, firstNode.Val)
	require.Equal(t, 1, list.Len())
}

func TestDeleteAtCanDeleteLast(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)
	require.Equal(t, 2, list.Len())

	require.NoError(t, list.DeleteAt(1))
	firstNode, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, firstVal, firstNode.Val)
	require.Equal(t, 1, list.Len())
}

func TestDeleteAtCanDeleteMid(t *testing.T) {
	list := NewSinglyLinkedList[int]()
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

func TestCreateStringList(t *testing.T) {
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
func TestGetInIntList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	firstVal, secondVal := 1, 2
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)

	firstNode, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, firstVal, firstNode.Val)

	secondNode, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, secondVal, secondNode.Val)
}

func TestGetInStringList(t *testing.T) {
	list := NewSinglyLinkedList[string]()
	firstVal, secondVal := "a", "b"
	list.InsertLast(firstVal)
	list.InsertLast(secondVal)

	firstNode, err := list.Get(0)
	require.NoError(t, err)
	require.Equal(t, firstVal, firstNode.Val)

	secondNode, err := list.Get(1)
	require.NoError(t, err)
	require.Equal(t, secondVal, secondNode.Val)
}
