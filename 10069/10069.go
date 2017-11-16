// UVa 10069 - Distinct Subsequences

package main

import (
	"fmt"
	"math/big"
	"os"
)

func solve(x, z string) *big.Int {
	dp := make([][]big.Int, len(z)+1)
	for i := range dp {
		dp[i] = make([]big.Int, len(x)+1)
	}
	var cnt int64
	for i := range z {
		for j := range x {
			if i == 0 {
				if z[i] == x[j] {
					cnt++
				}
				dp[i+1][j+1].SetInt64(cnt)
			} else {
				if z[i] == x[j] {
					dp[i+1][j+1].Add(&dp[i+1][j], &dp[i][j])
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}
	return &dp[len(z)][len(x)]
}

func main() {
	in, _ := os.Open("10069.in")
	defer in.Close()
	out, _ := os.Create("10069.out")
	defer out.Close()

	var n int
	var x, z string
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%s", &x)
		fmt.Fscanf(in, "%s", &z)
		fmt.Fprintln(out, solve(x, z))
	}
}
