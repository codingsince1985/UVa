// UVa 580 - Critical Mass

package main

import (
	"fmt"
	"math"
	"os"
)

func solve(s int) int {
	if s < 3 {
		return 0
	}
	dp := make([]int, s+1)
	dp[0], dp[1], dp[2], dp[3] = 1, 1, 2, 4
	for i := 4; i <= s; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return int(math.Pow(2, float64(s))) - dp[s] - dp[s-1] - dp[s-2]
}

func main() {
	in, _ := os.Open("580.in")
	defer in.Close()
	out, _ := os.Create("580.out")
	defer out.Close()

	var s int
	for {
		if fmt.Fscanf(in, "%d", &s); s == 0 {
			break
		}
		fmt.Fprintln(out, solve(s))
	}
}
