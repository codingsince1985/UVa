// UVa 10446 - The Marriage Interview :-)

package main

import (
	"fmt"
	"os"
)

const max = 61

var dp [][]int64

func trib(n, b int) int64 {
	if n < 0 {
		n = 0
	}
	if dp[n][b] == 0 {
		dp[n][b] = 1
		if n > 1 {
			for i := 1; i <= b; i++ {
				dp[n][b] += trib(n-i, b)
			}
		}
	}
	return dp[n][b]
}

func main() {
	in, _ := os.Open("10446.in")
	defer in.Close()
	out, _ := os.Create("10446.out")
	defer out.Close()

	dp = make([][]int64, max)
	for i := range dp {
		dp[i] = make([]int64, max)
	}
	var n, b int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &n, &b); n >= max {
			break
		}
		fmt.Fprintf(out, "Case %d: %d\n", kase, trib(n, b))
	}
}
