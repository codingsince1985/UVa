// UVa 836 - Largest Submatrix

package main

import (
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(m [][]byte) int {
	l := len(m)
	dp := make([][]int, l+1)
	for i := range dp {
		dp[i] = make([]int, l+1)
	}

	max := 0
	for i := 1; i <= l; i++ {
		for j := 1; j <= l; j++ {
			if m[i-1][j-1] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}

func main() {
	in, _ := os.Open("836.in")
	defer in.Close()
	out, _ := os.Create("836.out")
	defer out.Close()

	var c int
	var m [][]byte
	for fmt.Fscanf(in, "%d\n\n", &c); c > 0; c-- {
		first := true
		for count := 0; ; count++ {
			var l string
			if fmt.Fscanf(in, "%s", &l); l == "" {
				break
			}
			if first {
				first = false
				n := len(l)
				m = make([][]byte, n)
				for i := range m {
					m[i] = make([]byte, n)
				}
			}
			m[count] = []byte(l)
		}
		fmt.Fprintln(out, solve(m))
	}
}
