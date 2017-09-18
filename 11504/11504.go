// UVa 11504 - Dominos

package main

import (
	"fmt"
	"os"
)

var (
	count   int
	matrix  [][]bool
	visited map[int]bool
	stack   []int
)

func dfs(curr int, track bool) {
	visited[curr] = true
	for i, cell := range matrix[curr] {
		if !visited[i] && cell {
			dfs(i, track)
		}
	}
	if track {
		stack = append([]int{curr}, stack...)
	}
}

func solve() int {
	stack = nil
	visited = make(map[int]bool)
	for i := range matrix {
		if !visited[i] {
			dfs(i, true)
		}
	}
	count = 0
	visited = make(map[int]bool)
	for _, i := range stack {
		if !visited[i] {
			count++
			dfs(i, false)
		}
	}
	return count
}

func main() {
	in, _ := os.Open("11504.in")
	defer in.Close()
	out, _ := os.Create("11504.out")
	defer out.Close()

	var kase, n, m, d1, d2 int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d", &n, &m)
		matrix = make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		for ; m > 0; m-- {
			fmt.Fscanf(in, "%d%d", &d1, &d2)
			matrix[d1-1][d2-1] = true
		}
		fmt.Fprintln(out, solve())
	}
}
