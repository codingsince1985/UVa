// UVa 10304 - Optimal Binary Search Tree

package main

import (
	"fmt"
	"math"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(n int, f []int) int {
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + f[i-1]
	}
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := 2; i <= n; i++ {
		for j := 1; j+i-1 <= n; j++ {
			s, e, lowest := j, j+i-1, math.MaxInt32
			for k := s; k <= e; k++ {
				lowest = min(lowest, dp[s][k-1]+dp[k+1][e]+prefixSum[e]-prefixSum[s-1]-f[k-1])
			}
			dp[s][e] = lowest
		}
	}
	return dp[1][n]
}

func main() {
	in, _ := os.Open("10304.in")
	defer in.Close()
	out, _ := os.Create("10304.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		f := make([]int, n)
		for i := range f {
			fmt.Fscanf(in, "%d", &f[i])
		}
		fmt.Fprintln(out, solve(n, f))
	}
}
