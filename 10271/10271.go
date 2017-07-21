// UVa 10271 - Chopsticks

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	bad []int
	dp  [][]int
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func doDP(k, n int) int {
	if dp[k][n] == -1 {
		switch {
		case n < 3*k:
			dp[k][n] = math.MaxInt32
		case k == 0:
			dp[k][n] = 0
		default:
			dp[k][n] = min(doDP(k, n-1), doDP(k-1, n-2)+bad[n-1])
		}
	}
	return dp[k][n]
}

func solve(k, n int, chopsticks []int) int {
	bad = make([]int, n)
	for i := 1; i < n; i++ {
		bad[n-i] = (chopsticks[i] - chopsticks[i-1]) * (chopsticks[i] - chopsticks[i-1])
	}
	dp = make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return doDP(k, n)
}

func main() {
	in, _ := os.Open("10271.in")
	defer in.Close()
	out, _ := os.Create("10271.out")
	defer out.Close()

	var t, k, n int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d%d", &k, &n)
		chopsticks := make([]int, n)
		for i := range chopsticks {
			fmt.Fscanf(in, "%d", &chopsticks[i])
		}
		fmt.Fprintln(out, solve(k+8, n, chopsticks))
	}
}
