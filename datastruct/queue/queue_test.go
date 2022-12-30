package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {
	queue := NewQueue[int]()
	require.True(t, queue.IsEmpty())

	queue.Enqueue(1)
	require.False(t, queue.IsEmpty())
	require.Equal(t, 1, queue.Len())
}

func TestDequeue(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1)

	front, ok := queue.Dequeue()
	require.True(t, ok)
	require.Equal(t, 1, front)

	front, ok = queue.Dequeue()
	require.False(t, ok)
	require.Equal(t, 0, front)
}

func TestFront(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1)

	front, ok := queue.Front()
	require.True(t, ok)
	require.Equal(t, 1, front)
}

func TestLen(t *testing.T) {
	queue := NewQueue[int]()

	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		require.Equal(t, i+1, queue.Len())
	}
}
