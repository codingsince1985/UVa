// UVa 10066 - The Twin Towers

package main

import (
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(r1, r2 []int) int {
	var dp [101][101]int
	n, m := len(r1), len(r2)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if r1[i - 1] == r2[j - 1] { // dp array is 1-based
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[n][m]
}

func main() {
	in, _ := os.Open("10066.in")
	defer in.Close()
	out, _ := os.Create("10066.out")
	defer out.Close()

	var n1, n2, count int
	var r1, r2 []int
	for {
		fmt.Fscanf(in, "%d%d", &n1, &n2)
		if n1 == 0 && n2 == 0 {
			break
		}
		count ++
		r1 = make([]int, n1)
		r2 = make([]int, n2)
		for i := 0; i < n1; i++ {
			fmt.Fscanf(in, "%d", &r1[i])
		}
		for i := 0; i < n2; i++ {
			fmt.Fscanf(in, "%d", &r2[i])
		}

		fmt.Fprintf(out, "Twin Towers #%d\n", count)
		fmt.Fprintf(out, "Number of Tiles : %d\n\n", lcs(r1, r2))
	}
}
