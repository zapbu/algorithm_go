package sorting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func createSlice() []int {
	return []int{4, 3, 2, 7, 4, 11, 33, 44, 55}
}
func TestBubbleSort(t *testing.T) {
	slice := createSlice()
	BubbleSort(slice)
	fmt.Println(slice)
	for i := 0; i < len(slice)-1; i++ {
		require.True(t, slice[i] <= slice[i+1])
	}
}
