// UVa 10100 - Longest Match

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(t1, t2 []string) int {
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

	var s1, s2 string
	var t1, t2 []string
	var count int
	for s.Scan() {
		s1 = s.Text()

		s.Scan()
		s2 = s.Text()

		s1 = convertNonLetter(s1)
		s2 = convertNonLetter(s2)
		s1 = strings.TrimSpace(s1)
		s2 = strings.TrimSpace(s2)

		count++
		fmt.Fprintf(out, "%2d. ", count)

		if len(s1) == 0 || len(s2) == 0 {
			fmt.Fprintln(out, "Blank!")
			continue
		}
		t1 = strings.Split(s1, " ")
		t2 = strings.Split(s2, " ")
		fmt.Fprintf(out, "Length of longest match: %d\n", lcs(t1, t2))
	}
}
