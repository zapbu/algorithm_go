package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	slice  []int
	cnt    int
	ANS    int
	N, K   int
)

func main() {

	fmt.Scanf("%d %d", &N, &K)
	slice = make([]int, N)
	line, _ := reader.ReadString('\n')
	for i, numStr := range strings.Fields(line) {
		num, _ := strconv.Atoi(numStr)
		slice[i] = num
	}

	MergeSort(slice, 0, N-1)
	if ANS > 0 {
		fmt.Println(ANS)
	} else {
		fmt.Println(-1)
	}
}

func MergeSort(A []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		MergeSort(A, l, m)
		MergeSort(A, m+1, r)
		merge(A, l, m, r)
	}
}

func merge(A []int, l, m, r int) {
	i, j := l, m+1
	tmp := []int{}

	for i <= l && j <= r {
		if A[i] <= A[j] {
			tmp = append(tmp, A[i])
			i++
		} else {
			tmp = append(tmp, A[j])
			j++
		}
	}

	for i <= m {
		tmp = append(tmp, A[i])
		i++
	}

	for j <= r {
		tmp = append(tmp, A[j])
		j++
	}

	i = l
	tmpIdx := 0

	for i <= r {
		A[i] = tmp[tmpIdx]
		cnt += 1

		if cnt == K {
			ANS = A[i]
		}
		i++
		tmpIdx++
	}
}
