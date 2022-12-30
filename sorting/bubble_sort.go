package sorting

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](slice []T) {
	for i := 0; i < len(slice)-1; i++ {
		a := slice[i]
		for j := i + 1; j < len(slice); j++ {
			b := slice[j]
			if a > b {
				slice[i], slice[j] = slice[j], slice[i]
				continue
			}
		}
	}
}
