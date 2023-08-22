// UVa 111 - History Grading

package main

import (
	"fmt"
	"os"
)

func lcs(s1, s2 []int) int {
	l := len(s1)
	dp := make([][]int, l+1)
	for i := range dp {
		dp[i] = make([]int, l+1)
	}

	for i := 1; i <= l; i++ {
		for j := 1; j <= l; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l][l]
}

func main() {
	in, _ := os.Open("111.in")
	defer in.Close()
	out, _ := os.Create("111.out")
	defer out.Close()

	var n, tmp int
	fmt.Fscanf(in, "%d", &n)

	s1 := make([]int, n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(in, "%d", &tmp)
		s1[tmp-1] = i
	}

	s2 := make([]int, n)
here:
	for {
		for i := 1; i <= n; i++ {
			if _, err := fmt.Fscanf(in, "%d", &tmp); err != nil {
				break here
			}
			s2[tmp-1] = i
		}
		fmt.Fprintln(out, lcs(s1, s2))
	}
}
