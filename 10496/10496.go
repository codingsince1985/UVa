// UVa 10496 - Collecting Beepers

package main

import (
	"fmt"
	"math"
	"os"
)

type point struct{ x, y int }

var (
	n          int
	matrix, dp [][]int
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func tsp(state, last, n int) int {
	if state == 0 && last != 0 {
		return math.MaxInt32
	}
	if dp[state][last] != -1 {
		return dp[state][last]
	}
	min := math.MaxInt32
	for i := range matrix {
		if (state & (1 << uint32(i))) != 0 {
			if tmp := tsp(state-1<<uint32(i), i, n) + matrix[i][last]; tmp < min {
				min = tmp
			}
		}
	}
	dp[state][last] = min
	return dp[state][last]
}

func solve(beepers []point) int {
	matrix = make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j] = abs(beepers[i].x-beepers[j].x) + abs(beepers[i].y-beepers[j].y)
			matrix[j][i] = matrix[i][j]
		}
	}
	dp = make([][]int, 1<<uint32(n))
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	return tsp(1<<uint32(n)-1, 0, n)
}

func main() {
	in, _ := os.Open("10496.in")
	defer in.Close()
	out, _ := os.Create("10496.out")
	defer out.Close()

	var kase, x, y int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d", &x, &y)
		fmt.Fscanf(in, "%d%d", &x, &y)
		fmt.Fscanf(in, "%d", &n)
		n++
		beepers := make([]point, n)
		beepers[0] = point{x, y}
		for i := 1; i < len(beepers); i++ {
			fmt.Fscanf(in, "%d%d", &beepers[i].x, &beepers[i].y)
		}
		fmt.Fprintf(out, "The shortest path has length %d\n", solve(beepers))
	}
}
