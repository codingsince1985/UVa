// UVa 10908 - Largest Square

package main

import (
	"fmt"
	"os"
)

var m, n int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initialize(grid [][]byte) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == grid[i-1][j-1] && grid[i][j] == grid[i-1][j] && grid[i][j] == grid[i][j-1] {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp
}

func solve(dp [][]int, y, x int) int {
	center := dp[y][x]
	for i := 1; i < center; i++ {
		if ny, nx := y+i, x+i; !(ny < m && nx < n && dp[ny][nx] == center+i) {
			return i
		}
	}
	return center
}

func main() {
	in, _ := os.Open("10908.in")
	defer in.Close()
	out, _ := os.Create("10908.out")
	defer out.Close()

	var kase, q, y, x int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d%d", &m, &n, &q)
		grid := make([][]byte, m)
		for i := range grid {
			fmt.Fscanf(in, "%s", &line)
			grid[i] = []byte(line)
		}
		fmt.Fprintln(out, m, n, q)
		dp := initialize(grid)
		for ; q > 0; q-- {
			fmt.Fscanf(in, "%d%d", &y, &x)
			fmt.Fprintln(out, (solve(dp, y, x)-1)*2+1)
		}
	}
}
