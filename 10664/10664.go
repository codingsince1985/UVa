// UVa 10664 - Luggage

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func knapsack(luggage []int, sum int) bool {
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := range luggage {
		for j := sum; j >= luggage[i]; j-- {
			dp[j] = dp[j-luggage[i]]
		}
	}
	return dp[sum/2]
}

func solve(token []string) bool {
	var sum int
	luggage := make([]int, len(token))
	for i := range luggage {
		luggage[i], _ = strconv.Atoi(token[i])
		sum += luggage[i]
	}
	return sum%2 == 0 && knapsack(luggage, sum)
}

func main() {
	in, _ := os.Open("10664.in")
	defer in.Close()
	out, _ := os.Create("10664.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	for m, _ := strconv.Atoi(s.Text()); m > 0 && s.Scan(); m-- {
		if solve(strings.Split(s.Text(), " ")) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
