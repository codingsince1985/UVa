// UVa 10404 - Bachet's Game

package main

import (
	"fmt"
	"os"
)

func knapsack(n int, stones []int) bool {
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for _, stone := range stones {
			if i-stone >= 0 && !dp[i-stone] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func main() {
	in, _ := os.Open("10404.in")
	defer in.Close()
	out, _ := os.Create("10404.out")
	defer out.Close()

	var n, m int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &m); err != nil {
			break
		}
		stones := make([]int, m)
		for i := range stones {
			fmt.Fscanf(in, "%d", &stones[i])
		}
		if knapsack(n, stones) {
			fmt.Fprintln(out, "Stan wins")
		} else {
			fmt.Fprintln(out, "Ollie wins")
		}
	}
}
