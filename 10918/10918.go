// UVa 10918 - Tri Tiling

package main

import (
	"fmt"
	"os"
)

const max = 30

var dp = func() []int {
	dp := make([]int, max+1)
	dp[0], dp[2] = 1, 3
	for i := 4; i <= max; i = i + 2 {
		dp[i] = 4*dp[i-2] - dp[i-4]
	}
	return dp
}()

func main() {
	in, _ := os.Open("10918.in")
	defer in.Close()
	out, _ := os.Create("10918.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == -1 {
			break
		}
		fmt.Fprintln(out, dp[n])
	}
}
