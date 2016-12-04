// UVa 10810 - Ultra-QuickSort

package main

import (
	"fmt"
	"os"
)

func merge(num []int, mid int) int {
	tmp := make([]int, len(num))
	l, r, idx, steps := 0, mid, 0, 0
	for l < mid || r < len(num) {
		if l < mid && r < len(num) {
			if num[l] > num[r] {
				tmp[idx], r = num[r], r+1
				steps += mid - l
			} else {
				tmp[idx], l = num[l], l+1
			}
		} else {
			if l < mid {
				tmp[idx], l = num[l], l+1
			} else {
				tmp[idx], r = num[r], r+1
			}
		}
		idx++
	}
	copy(num, tmp)
	return steps
}

func sort(num []int) int {
	if len(num) <= 1 {
		return 0
	}
	mid := len(num) / 2
	return sort(num[:mid]) + sort(num[mid:]) + merge(num, mid)
}

func main() {
	in, _ := os.Open("10810.in")
	defer in.Close()
	out, _ := os.Create("10810.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		num := make([]int, n)
		for i := range num {
			fmt.Fscanf(in, "%d", &num[i])
		}
		fmt.Fprintln(out, sort(num))
	}
}
