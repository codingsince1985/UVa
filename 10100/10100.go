// UVa 10100 - Longest Match

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lcs(s1, s2 string) int {
	t1, t2 := strings.Fields(s1), strings.Fields(s2)
	l1, l2 := len(t1), len(t2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if t1[i-1] == t2[j-1] { // dp array is 1-based
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func convertNonLetter(s string) string {
	r := []byte(s)
	for i, v := range s {
		if !(v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z' || v == ' ' || v >= '0' && v <= '9') {
			r[i] = ' '
		}
	}
	return string(r)
}

func main() {
	in, _ := os.Open("10100.in")
	defer in.Close()
	out, _ := os.Create("10100.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for count := 1; s.Scan(); count++ {
		s1 := strings.TrimSpace(convertNonLetter(s.Text()))
		s.Scan()
		s2 := strings.TrimSpace(convertNonLetter(s.Text()))

		fmt.Fprintf(out, "%2d. ", count)
		if len(s1) == 0 || len(s2) == 0 {
			fmt.Fprintln(out, "Blank!")
		} else {
			fmt.Fprintf(out, "Length of longest match: %d\n", lcs(s1, s2))
		}
	}
}
