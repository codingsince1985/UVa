// UVa 10453 - Make Palindrome

package main

import (
	"fmt"
	"os"
	"strings"
)

func output(l, count int, line string, dp [][]int) string {
	var firstBuilder strings.Builder
	for i, j := 0, l-1; i <= j; {
		switch {
		case line[i] == line[j]:
			firstBuilder.WriteByte(line[i])
			i++
			j--
		case dp[i][j] == dp[i+1][j]+1:
			firstBuilder.WriteByte(line[i])
			i++
		default:
			firstBuilder.WriteByte(line[j])
			j--
		}
	}
	var secondBuilder strings.Builder
	first := firstBuilder.String()
	for i := len(first) - (l+count)%2 - 1; i >= 0; i-- {
		secondBuilder.WriteByte(first[i])
	}
	return first + secondBuilder.String()
}

func solve(line string) (int, string) {
	l := len(line)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
	}
	for i := 1; i <= l; i++ {
		for j := 0; j+i < l; j++ {
			switch {
			case line[j] == line[j+i]:
				dp[j][j+i] = dp[j+1][j+i-1]
			case dp[j+1][j+i] <= dp[j][j+i-1]:
				dp[j][j+i] = dp[j+1][j+i] + 1
			default:
				dp[j][j+i] = dp[j][j+i-1] + 1
			}
		}
	}
	return dp[0][l-1], output(l, dp[0][l-1], line, dp)
}

func main() {
	in, _ := os.Open("10453.in")
	defer in.Close()
	out, _ := os.Create("10453.out")
	defer out.Close()

	var line string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		count, palindrome := solve(line)
		fmt.Fprintln(out, count, palindrome)
	}
}
