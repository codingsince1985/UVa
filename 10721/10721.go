// UVa 10721 - Bar Codes

package main

import (
	"fmt"
	"os"
)

const max = 50

var dp = func() [][][]int64 {
	dp := make([][][]int64, max+1)
	for i := range dp {
		dp[i] = make([][]int64, max+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, max+1)
			if i == 0 && j == 0 {
				for k := range dp[i][j] {
					dp[i][j][k] = 1
				}
			}
		}
	}
	for i := 1; i <= max; i++ {
		for j := 1; j <= max; j++ {
			for k := 1; k <= max; k++ {
				for l := 1; l <= i && l <= k; l++ {
					dp[i][j][k] += dp[i-l][j-1][k]
				}
			}
		}
	}
	return dp
}()

func main() {
	in, _ := os.Open("10721.in")
	defer in.Close()
	out, _ := os.Create("10721.out")
	defer out.Close()

	var n, k, m int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d", &n, &k, &m); err != nil {
			break
		}
		fmt.Fprintln(out, dp[n][k][m])
	}
}
