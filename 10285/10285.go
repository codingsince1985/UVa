// UVa 10285 - Longest Run on a Snowboard

package main

import (
	"fmt"
	"os"
)

var (
	r, c       int
	grid, dp   [][]int
	directions = [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
)

func dfs(i, j int) int {
	if dp[i][j] == 0 {
		for _, direction := range directions {
			newI, newJ := i+direction[1], j+direction[0]
			if newI >= 0 && newI < r && newJ >= 0 && newJ < c && grid[i][j] > grid[newI][newJ] {
				dp[i][j] = max(dp[i][j], dfs(newI, newJ)+1)
			}
		}
	}
	return dp[i][j]
}

func longestRun() int {
	dp = make([][]int, r)
	for i := range dp {
		dp[i] = make([]int, c)
	}
	var longest int
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			longest = max(longest, dfs(i, j))
		}
	}
	return longest + 1
}

func main() {
	in, _ := os.Open("10285.in")
	defer in.Close()
	out, _ := os.Create("10285.out")
	defer out.Close()

	var n int
	var name string
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%s%d%d", &name, &r, &c)
		grid = make([][]int, r)
		for i := range grid {
			grid[i] = make([]int, c)
			for j := range grid[i] {
				fmt.Fscanf(in, "%d", &grid[i][j])
			}
		}
		fmt.Fprintf(out, "%s: %d\n", name, longestRun())
	}
}
