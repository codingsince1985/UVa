// UVa 11137 - Ingenuous Cubrency

package main

import (
	"fmt"
	"os"
)

var coins = func() []int {
	coins := make([]int, 21)
	for i := range coins {
		coins[i] = (i + 1) * (i + 1) * (i + 1)
	}
	return coins
}()

func knapsack(n int) int64 {
	dp := make([]int64, n+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= n; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[n]
}

func main() {
	in, _ := os.Open("11137.in")
	defer in.Close()
	out, _ := os.Create("11137.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, knapsack(n))
	}
}
