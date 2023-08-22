// UVa 10003 - Cutting Sticks

package main

import (
	"fmt"
	"math"
	"os"
)

var dp [][]int

func solve(c []int, l, r int) int {
	if dp[l][r] != -1 {
		return dp[l][r]
	}
	if l+1 == r {
		dp[l][r] = 0
		return 0
	}
	dp[l][r] = math.MaxInt32
	for i := l + 1; i < r; i++ {
		dp[l][r] = min(dp[l][r], solve(c, l, i)+solve(c, i, r)+(c[r]-c[l]))
	}
	return dp[l][r]
}

func main() {
	in, _ := os.Open("10003.in")
	defer in.Close()
	out, _ := os.Create("10003.out")
	defer out.Close()

	var l, n int
	for {
		if fmt.Fscanf(in, "%d", &l); l == 0 {
			break
		}
		fmt.Fscanf(in, "%d", &n)
		newl := n + 2
		c := make([]int, newl)
		c[0], c[n+1] = 0, l
		for i := 1; i <= n; i++ {
			fmt.Fscanf(in, "%d", &c[i])
		}

		dp = make([][]int, newl)
		for i := range dp {
			dp[i] = make([]int, newl)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		fmt.Fprintf(out, "The minimum cutting is %d.\n", solve(c, 0, newl-1))
	}
}
