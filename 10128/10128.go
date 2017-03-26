// UVa 10128 - Queue

package main

import (
	"fmt"
	"os"
)

const max = 14

var dp = func() [][][]int {
	dp := make([][][]int, max)
	for i := range dp {
		dp[i] = make([][]int, max)
		for j := range dp[i] {
			dp[i][j] = make([]int, max)
		}
	}
	dp[1][1][1] = 1
	for n := 2; n < max; n++ {
		for p := 1; p <= n; p++ {
			for r := 1; r <= n; r++ {
				dp[n][p][r] = dp[n-1][p-1][r] + dp[n-1][p][r-1] + dp[n-1][p][r]*(n-2)
			}
		}
	}
	return dp
}()

func main() {
	in, _ := os.Open("10128.in")
	defer in.Close()
	out, _ := os.Create("10128.out")
	defer out.Close()

	var kase, n, p, r int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d%d", &n, &p, &r)
		fmt.Fprintln(out, dp[n][p][r])
	}
}
