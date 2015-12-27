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

func knapsack(p, w []int, tw int) int {
	l := len(p)
	dp := make([][]int, l + 1)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, tw + 1)
	}

	for i := 1; i <= l; i++ {
		for j := 1; j <= tw; j++ {
			// p and w are 0-based
			if w[i - 1] > j {
				dp[i][j] = dp[i - 1][j]
			} else {
				dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - w[i - 1]] + p[i - 1])
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

	var t int
	fmt.Fscanf(in, "%d", &t)
	for i := 0; i < t; i++ {
		var n, g, tw, total int
		fmt.Fscanf(in, "%d", &n)
		p := make([]int, n)
		w := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(in, "%d%d", &p[j], &w[j])
		}
		fmt.Fscanf(in, "%d", &g)
		for j := 0; j < g; j++ {
			fmt.Fscanf(in, "%d", &tw)
			total += knapsack(p, w, tw)
		}
		fmt.Fprintln(out, total)
	}
}
