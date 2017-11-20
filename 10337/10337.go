// UVa 10337 - Flight Planner

package main

import (
	"fmt"
	"math"
	"os"
)

const level = 10

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(steps int, matrix [][]int) int {
	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, level)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	for i := range dp {
		if i != steps {
			for j, cell := range dp[i] {
				wind := matrix[i+1][j]
				if j+1 < level {
					dp[i+1][j+1] = min(dp[i+1][j+1], cell+60-wind)
				}
				if j > 0 {
					dp[i+1][j-1] = min(dp[i+1][j-1], cell+20-wind)
				}
				dp[i+1][j] = min(dp[i+1][j], cell+30-wind)
			}
		}
	}
	return dp[steps][0]
}

func main() {
	in, _ := os.Open("10337.in")
	defer in.Close()
	out, _ := os.Create("10337.out")
	defer out.Close()

	var n, x int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "\n%d", &x)
		steps := x / 100
		matrix := make([][]int, steps+1)
		for i := range matrix {
			matrix[i] = make([]int, level)
		}
		for i := 0; i < level; i++ {
			for j := 1; j <= steps; j++ {
				fmt.Fscanf(in, "%d", &matrix[j][level-i-1])
			}
		}
		fmt.Fprintf(out, "%d\n\n", solve(steps, matrix))
	}
}
