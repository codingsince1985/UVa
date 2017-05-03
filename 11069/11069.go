// UVa 11069 - A Graph Problem

package main

import (
	"fmt"
	"os"
)

const max = 76

var dp = func() []int {
	dp := make([]int, max+1)
	dp[1], dp[2], dp[3] = 1, 2, 2
	for i := 4; i <= max; i++ {
		dp[i] = dp[i-3] + dp[i-2]
	}
	return dp
}()

func main() {
	in, _ := os.Open("11069.in")
	defer in.Close()
	out, _ := os.Create("11069.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, dp[n])
	}
}
