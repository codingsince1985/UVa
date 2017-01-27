// UVa 10004 - Bicoloring

package main

import (
	"fmt"
	"os"
)

func valid(matrix [][]bool, nodes []int, curr, color int) bool {
	for i := range matrix[curr] {
		if matrix[curr][i] && nodes[i] == color {
			return false
		}
	}
	return true
}

func allVisited(nodes []int) bool {
	for _, v := range nodes {
		if v == 0 {
			return false
		}
	}
	return true
}

func dfs(matrix [][]bool, nodes []int, curr, color int) bool {
	if allVisited(nodes) {
		return true
	}
	for i := range nodes {
		if nodes[i] == 0 {
			var nextColor int
			if matrix[curr][i] {
				nextColor = 3 - color
			} else {
				nextColor = color
			}
			if valid(matrix, nodes, i, nextColor) {
				nodes[i] = nextColor
				if dfs(matrix, nodes, i, nextColor) {
					return true
				}
				nodes[i] = 0
			}
		}
	}
	return false
}

func main() {
	in, _ := os.Open("10004.in")
	defer in.Close()
	out, _ := os.Create("10004.out")
	defer out.Close()

	var n, l, n1, n2 int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		matrix := make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		for fmt.Fscanf(in, "%d", &l); l > 0; l-- {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			matrix[n1][n2], matrix[n2][n1] = true, true
		}
		nodes := make([]int, n)
		nodes[0] = 1
		if dfs(matrix, nodes, 0, 1) {
			fmt.Fprintln(out, "BICOLORABLE.")
		} else {
			fmt.Fprintln(out, "NOT BICOLORABLE.")
		}
	}
}
