// UVa 674 - Coin Change

package main

import (
	"fmt"
	"os"
)

var cs = []int{1, 5, 10, 25, 50}

func main() {
	in, _ := os.Open("674.in")
	defer in.Close()
	out, _ := os.Create("674.out")
	defer out.Close()

	var v int
	var dp []uint64
	for {
		if _, err := fmt.Fscanf(in, "%d", &v); err != nil {
			break
		}
		dp = make([]uint64, v+1)
		dp[0] = 1

		for i := range cs {
			c := cs[i]
			for j := c; j <= v; j++ {
				dp[j] += dp[j-c]
			}
		}
		fmt.Fprintln(out, dp[v])
	}
}
