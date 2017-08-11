// UVa 10910 - Marks Distribution

package main

import (
	"fmt"
	"os"
)

func solve(n, t, p int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}
	for i := p; i <= t; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for x := p; x <= t-p; x++ {
			for y := p; y+x <= t; y++ {
				dp[i][x+y] += dp[i-1][x]
			}
		}
	}
	return dp[n][t]
}

func main() {
	in, _ := os.Open("10910.in")
	defer in.Close()
	out, _ := os.Create("10910.out")
	defer out.Close()

	var k, n, t, p int
	for fmt.Fscanf(in, "%d", &k); k > 0; k-- {
		fmt.Fscanf(in, "%d%d%d", &n, &t, &p)
		fmt.Fprintln(out, solve(n, t, p))
	}
}
