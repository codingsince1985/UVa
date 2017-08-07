// UVa 10081 - Tight Words

package main

import (
	"fmt"
	"os"
)

func solve(k, n int) float64 {
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}
	for i := range dp[1] {
		dp[1][i] = 100.0 / float64(k+1)
	}
	for i := 2; i <= n; i++ {
		for j := range dp[i] {
			dp[i][j] = dp[i-1][j]
			if j != 0 {
				dp[i][j] += dp[i-1][j-1]
			}
			if j != k {
				dp[i][j] += dp[i-1][j+1]
			}
			dp[i][j] /= float64(k + 1)
		}
	}
	var ret float64
	for i := range dp[n] {
		ret += dp[n][i]
	}
	return ret
}

func main() {
	in, _ := os.Open("10081.in")
	defer in.Close()
	out, _ := os.Create("10081.out")
	defer out.Close()

	var k, n int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &k, &n); err != nil {
			break
		}
		fmt.Fprintf(out, "%.5f\n", solve(k, n))
	}
}
