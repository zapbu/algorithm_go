package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	require.True(t, stack.IsEmpty())
	require.NotNil(t, stack)
	require.Nil(t, stack.head)

	for i := 0; i < 10; i++ {
		stack.Push(i)
		require.Equal(t, i, stack.Top().Val)
		require.False(t, stack.IsEmpty())
	}

	for i := 9; i >= 0; i-- {
		require.Equal(t, i, stack.Pop().Val)
	}
}
