// UVa 590 - Always on the run

package main

import (
	"fmt"
	"math"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initialise(n, k int) [][]int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	return dp
}

func solve(n, k int, flights [][][]int) int {
	dp := initialise(n, k)
	for i := 1; i <= k; i++ {
		for j := 0; j < n; j++ {
			for l := 0; l < n; l++ {
				if j != l {
					if idx := (i - 1) % len(flights[l][j]); flights[l][j][idx] != 0 && dp[i-1][l] != math.MaxInt32 {
						dp[i][j] = min(dp[i][j], dp[i-1][l]+flights[l][j][idx])
					}
				}
			}
		}
	}
	return dp[k][n-1]
}

func main() {
	in, _ := os.Open("590.in")
	defer in.Close()
	out, _ := os.Create("590.out")
	defer out.Close()

	var n, k, f int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &n, &k); n == 0 && k == 0 {
			break
		}
		flights := make([][][]int, n)
		for i := range flights {
			flights[i] = make([][]int, n)
			for j := range flights[i] {
				if i != j {
					fmt.Fscanf(in, "%d", &f)
					flights[i][j] = make([]int, f)
					for k := range flights[i][j] {
						fmt.Fscanf(in, "%d", &flights[i][j][k])
					}
				}
			}
		}
		fmt.Fprintf(out, "Scenario #%d\n", kase)
		if cost := solve(n, k, flights); cost != math.MaxInt32 {
			fmt.Fprintf(out, "The best flight costs %d.\n", cost)
		} else {
			fmt.Fprintln(out, "No flight possible.")
		}
		fmt.Fprintln(out)
	}
}
