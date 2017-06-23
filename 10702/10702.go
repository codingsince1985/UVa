// UVa 10702 - Travelling Salesman

package main

import (
	"fmt"
	"os"
)

var (
	c, t       int
	matrix, dp [][]int
	targets    map[int]bool
)

func solve() int {
	var max int
	for step := 2; step <= t; step++ {
		for i := 0; i < c; i++ {
			for j := 0; j < c; j++ {
				if newEarn := dp[step-1][i] + matrix[i][j]; newEarn > dp[step][j] {
					dp[step][j] = newEarn
					if targets[j] && newEarn > max {
						max = newEarn
					}
				}
			}
		}
	}
	return max
}

func main() {
	in, _ := os.Open("10702.in")
	defer in.Close()
	out, _ := os.Create("10702.out")
	defer out.Close()

	var s, e, tmp int
	for {
		if fmt.Fscanf(in, "%d%d%d%d", &c, &s, &e, &t); c == 0 && s == 0 && e == 0 && t == 0 {
			break
		}
		dp = make([][]int, t+1)
		for i := range dp {
			dp[i] = make([]int, c)
		}
		matrix = make([][]int, c)
		for i := range matrix {
			matrix[i] = make([]int, c)
			for j := range matrix[i] {
				fmt.Fscanf(in, "%d", &matrix[i][j])
				if i+1 == s {
					dp[1][j] = matrix[i][j]
				}
			}
		}
		targets = make(map[int]bool)
		for ; e > 0; e-- {
			fmt.Fscanf(in, "%d", &tmp)
			targets[tmp-1] = true
		}
		fmt.Fprintln(out, solve())
		fmt.Fscanln(in)
	}
}
