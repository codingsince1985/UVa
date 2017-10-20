// UVa 10617 - Again Palindrome

package main

import (
	"fmt"
	"os"
)

func solve(line string) int {
	l := len(line)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
		dp[i][i] = 1
	}
	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if line[i] == line[j] {
				dp[i][j] = dp[i+1][j] + dp[i][j-1] + 1
			} else {
				dp[i][j] = dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1]
			}
		}
	}
	return dp[0][l-1]
}

func main() {
	in, _ := os.Open("10617.in")
	defer in.Close()
	out, _ := os.Create("10617.out")
	defer out.Close()

	var kase int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &line)
		fmt.Fprintln(out, solve(line))
	}
}
