// UVa 11151 - Longest Palindrome

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type byteSlice []byte

func (b byteSlice) Len() int           { return len(b) }
func (b byteSlice) Less(i, j int) bool { return i < j }
func (b byteSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

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

func solve(line string) int {
	var bs byteSlice = []byte(line)
	sort.Sort(sort.Reverse(bs))
	return lcs([]byte(line), bs)
}

func main() {
	in, _ := os.Open("11151.in")
	defer in.Close()
	out, _ := os.Create("11151.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	s.Scan()
	for t, _ := strconv.Atoi(s.Text()); t > 0 && s.Scan(); t-- {
		fmt.Fprintln(out, solve(s.Text()))
	}
}
