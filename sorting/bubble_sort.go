package sorting

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
