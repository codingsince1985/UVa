// UVa 539 - The Settlers of Catan

package main

import (
	"fmt"
	"os"
)

var (
	matrix [][]bool
	max    int
)

func dfs(node, depth int) {
	if depth > max {
		max = depth
	}
	for i := range matrix {
		if matrix[i][node] {
			matrix[i][node], matrix[node][i] = false, false
			dfs(i, depth+1)
			matrix[i][node], matrix[node][i] = true, true
		}
	}
}

func backtracking() {
	for i := range matrix {
		dfs(i, 0)
	}
}

func main() {
	in, _ := os.Open("539.in")
	defer in.Close()
	out, _ := os.Create("539.out")
	defer out.Close()

	var n, m, n1, n2 int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		max = 0
		matrix = make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		for i := 0; i < m; i++ {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			matrix[n1][n2], matrix[n2][n1] = true, true
		}
		backtracking()
		fmt.Fprintln(out, max)
	}
}
