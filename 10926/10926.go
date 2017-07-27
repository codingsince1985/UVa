// UVa 10926 - How Many Dependencies?

package main

import (
	"fmt"
	"os"
)

var (
	count  int
	matrix [][]bool
)

func dfs(curr int) {
	for i, cell := range matrix[curr] {
		if cell {
			count++
			dfs(i)
		}
	}
}

func solve() int {
	idx, max := -1, 0
	for i := range matrix {
		count = 0
		if dfs(i); count > max {
			idx = i
			max = count
		}
	}
	return idx + 1
}

func main() {
	in, _ := os.Open("10926.in")
	defer in.Close()
	out, _ := os.Create("10926.out")
	defer out.Close()

	var n, d, dep int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		matrix = make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
			for fmt.Fscanf(in, "%d", &d); d > 0; d-- {
				fmt.Fscanf(in, "%d", &dep)
				matrix[i][dep-1] = true
			}
		}
		fmt.Fprintln(out, solve())
	}
}
