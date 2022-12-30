package sorting

import "golang.org/x/exp/constraints"

func SelectionSort[T constraints.Ordered](slice []T) {
	for i := 0; i < len(slice)-1; i++ {
		minAt := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minAt] {
				minAt = j
			}
		}
		slice[i], slice[minAt] = slice[minAt], slice[i]
	}
}
