// UVa 11063 - B2-Sequence

package main

import (
	"fmt"
	"os"
)

func solve(n int, b []int) bool {
	for i, bi := range b {
		if bi < 1 || i > 0 && b[i-1] >= b[i] {
			return false
		}
	}
	var bij int
	sum := make(map[int]bool)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if bij = b[i] + b[j]; sum[bij] {
				return false
			}
			sum[bij] = true
		}
	}
	return true
}

func main() {
	in, _ := os.Open("11063.in")
	defer in.Close()
	out, _ := os.Create("11063.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		b := make([]int, n)
		for i := range b {
			fmt.Fscanf(in, "%d", &b[i])
		}
		if solve(n, b) {
			fmt.Fprintf(out, "Case #%d: It is a B2-Sequence.\n\n", kase)
		} else {
			fmt.Fprintf(out, "Case #%d: It is not a B2-Sequence.\n\n", kase)
		}
		fmt.Fscanln(in)
	}
}
