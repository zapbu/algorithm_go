package sorting

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}

	m := len(slice) / 2
	left, right := MergeSort(slice[:m]), MergeSort(slice[m:])
	return merge(left, right)
}

func merge[T constraints.Ordered](left, right []T) []T {
	new, n := make([]T, len(left)+len(right)), 0
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		fmt.Println(new)

		if left[l] < right[r] {
			new[n] = left[l]
			l++
		} else {
			new[n] = right[r]
			r++
		}
		n++
	}

	for l < len(left) {
		fmt.Println(new)
		new[n] = left[l]
		l++
		n++
	}

	for r < len(right) {
		fmt.Println(new)
		new[n] = right[r]
		r++
		n++
	}

	return new
}
