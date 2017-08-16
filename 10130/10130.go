// UVa 10130 - SuperSale

package main

import (
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsack(l, tw int, p, w []int) int {
	dp := make([][]int, l+1)
	for i := range dp {
		dp[i] = make([]int, tw+1)
	}

	for i := 1; i <= l; i++ {
		for j := 1; j <= tw; j++ {
			// p and w are 0-based
			if w[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i-1]]+p[i-1])
			}
		}
	}
	return dp[l][tw]
}

func main() {
	in, _ := os.Open("10130.in")
	defer in.Close()
	out, _ := os.Create("10130.out")
	defer out.Close()

	var t, n, g, tw int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d", &n)
		p, w := make([]int, n), make([]int, n)
		for j := range p {
			fmt.Fscanf(in, "%d%d", &p[j], &w[j])
		}
		var total int
		for fmt.Fscanf(in, "%d", &g); g > 0; g-- {
			fmt.Fscanf(in, "%d", &tw)
			total += knapsack(n, tw, p, w)
		}
		fmt.Fprintln(out, total)
	}
}
