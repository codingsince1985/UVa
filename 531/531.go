// UVa 531 - Compromise

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func nextLine(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(a, b []string) string {
	l1, l2 := len(a), len(b)
	dp := make([][][]string, l1 + 1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([][]string, l2 + 1)
		for j := 0; j <= l2; j++ {
			dp[i][j] = make([]string, 0)
		}
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if a[i - 1] == b[j - 1] {
				dp[i][j] = append(dp[i - 1][j - 1], a[i - 1])
			} else {
				if len(dp[i - 1][j]) > len(dp[i][j - 1]) {
					dp[i][j] = dp[i - 1][j]
				} else {
					dp[i][j] = dp[i][j - 1]
				}
			}
		}
	}
	return strings.Join(dp[l1][l2], " ")
}

func main() {
	in, _ := os.Open("531.in")
	defer in.Close()
	out, _ := os.Create("531.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	a := make([]string, 0)
	b := make([]string, 0)
	var l string
	first := true
	for {
		if l = nextLine(s); l == "" {
			break
		}
		if l == "#" {
			first = !first
			if first {
				fmt.Fprintln(out, lcs(a, b))
				a = make([]string, 0)
				b = make([]string, 0)
			}
			continue
		}
		tokens := strings.Split(l, " ")
		if first {
			a = append(a, tokens...)
		} else {
			b = append(b, tokens...)
		}
	}
}
