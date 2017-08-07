// UVa 116 - Unidirectional TSP

package main

import (
	"fmt"
	"math"
	"os"
)

var m, n int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func tsp(out *os.File, matrix [][]int) {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][n-1] = matrix[i][n-1]
	}
	path := make([][]int, m)
	for i := range path {
		path[i] = make([]int, n-1)
		for j := range path[i] {
			path[i][j] = math.MaxInt32
		}
	}

	for j := n - 2; j >= 0; j-- {
		for i := 0; i < m; i++ {
			fromUp := matrix[i][j] + dp[(i-1+m)%m][j+1]
			fromMid := matrix[i][j] + dp[i][j+1]
			fromDown := matrix[i][j] + dp[(i+1)%m][j+1]
			least := min(fromUp, min(fromMid, fromDown))
			dp[i][j] = least
			if least == fromUp {
				path[i][j] = min(path[i][j], (i-1+m)%m)
			}
			if least == fromMid {
				path[i][j] = min(path[i][j], i)
			}
			if least == fromDown {
				path[i][j] = min(path[i][j], (i+1)%m)
			}
		}
	}
	least := math.MaxInt32
	idx := 0
	for i := range dp {
		if dp[i][0] < least {
			least = dp[i][0]
			idx = i
		}
	}
	fmt.Fprint(out, idx+1)
	for i := 0; i < n-1; i++ {
		fmt.Fprint(out, " ", path[idx][i]+1)
		idx = path[idx][i]
	}
	fmt.Fprintf(out, "\n%d\n", least)
}

func main() {
	in, _ := os.Open("116.in")
	defer in.Close()
	out, _ := os.Create("116.out")
	defer out.Close()

	for {
		if _, err := fmt.Fscanf(in, "%d%d", &m, &n); err != nil {
			break
		}
		matrix := make([][]int, m)
		for i := range matrix {
			matrix[i] = make([]int, n)
			for j := range matrix[i] {
				fmt.Fscanf(in, "%d", &matrix[i][j])
			}
		}
		tsp(out, matrix)
	}
}
