// UVa 10192 - Vacation

package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(c1, c2 []byte) int {
	l1, l2 := len(c1), len(c2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if c1[i-1] == c2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func main() {
	in, _ := os.Open("10192.in")
	defer in.Close()
	out, _ := os.Create("10192.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var c1, c2 []byte
	var count int
	for s.Scan() {
		s1 := s.Text()
		if s1 == "#" {
			break
		}
		s.Scan()
		s2 := s.Text()
		c1 = []byte(s1)
		c2 = []byte(s2)
		count++
		fmt.Fprintf(out, "Case #%d: you can visit at most %d cities.\n", count, lcs(c1, c2))
	}
}
