// UVa 11450 - Wedding shopping

package main

import (
	"fmt"
	"os"
)

func solve(m, c int, garments [][]int) int {
	dp := make([][]int, c+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1
	for i, garment := range garments {
		for _, price := range garment {
			for k := price; k <= m; k++ {
				dp[i+1][k] += dp[i][k-price]
			}
		}
	}
	for ; m >= 0; m-- {
		if dp[c][m] != 0 {
			return m
		}
	}
	return 0
}

func main() {
	in, _ := os.Open("11450.in")
	defer in.Close()
	out, _ := os.Create("11450.out")
	defer out.Close()

	var kase, m, c, n, price int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d", &m, &c)
		garments := make([][]int, c)
		for i := range garments {
			for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
				fmt.Fscanf(in, "%d", &price)
				garments[i] = append(garments[i], price)
			}
		}
		if max := solve(m, c, garments); max == 0 {
			fmt.Fprintln(out, "no solution")
		} else {
			fmt.Fprintln(out, max)
		}
	}
}
