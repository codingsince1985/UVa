// UVa 434 - Matty's Blocks

package main

import (
	"fmt"
	"os"
)

const maxHeight = 9

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(k int, front, right []int) (int, int) {
	var m, n int
	for i, l, r := 0, k, k; i < maxHeight; i++ {
		m += max(front[i], right[i]) * i
		n += l * r
		l -= front[i]
		r -= right[i]
	}
	return m, n - k*k
}

func main() {
	in, _ := os.Open("434.in")
	defer in.Close()
	out, _ := os.Create("434.out")
	defer out.Close()

	var kase, n, k int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d", &k)
		front := make([]int, maxHeight)
		for i := 0; i < k; i++ {
			fmt.Fscanf(in, "%d", &n)
			front[n]++
		}
		right := make([]int, maxHeight)
		for i := 0; i < k; i++ {
			fmt.Fscanf(in, "%d", &n)
			right[n]++
		}
		m, n := solve(k, front, right)
		fmt.Fprintf(out, "Matty needs at least %d blocks, and can add at most %d extra blocks.\n", m, n-m)
	}
}
