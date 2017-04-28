// UVa 10943 - How do you add?

package main

import (
	"fmt"
	"os"
)

const max = 200

var dp = func() [][]int {
	dp := make([][]int, max+1)
	for i := range dp {
		dp[i] = make([]int, max+1)
		dp[i][0] = 1
	}
	for i := 1; i <= max; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = (dp[i-1][j-1] + dp[i-1][j]) % 1000000
		}
	}
	return dp
}()

func main() {
	in, _ := os.Open("10943.in")
	defer in.Close()
	out, _ := os.Create("10943.out")
	defer out.Close()

	var n, k int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &k); n == 0 && k == 0 {
			break
		}
		fmt.Fprintln(out, dp[n+k-1][k-1])
	}
}
