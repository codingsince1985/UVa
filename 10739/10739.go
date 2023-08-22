// UVa 10739 - String to Palindrome

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	word string
	dp   [][]int
)

func find(i, j int) int {
	switch {
	case i > j:
		return 0
	case dp[i][j] != -1:
		return dp[i][j]
	case word[i] == word[j]:
		return find(i+1, j-1)
	default:
		return min(min(find(i+1, j), find(i, j-1)), find(i+1, j-1)) + 1
	}
}

func solve() int {
	length := len(word)
	dp = make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return find(0, length-1)
}

func main() {
	in, _ := os.Open("10739.in")
	defer in.Close()
	out, _ := os.Create("10739.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var t int
	s.Scan()
	fmt.Sscanf(s.Text(), "%d", &t)
	for kase := 1; kase <= t && s.Scan(); kase++ {
		word = s.Text()
		fmt.Fprintf(out, "Case %d: %d\n", kase, solve())
	}
}
