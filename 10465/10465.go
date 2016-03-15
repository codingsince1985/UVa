// UVa 10465 - Homer Simpson

package main

import (
	"fmt"
	"os"
)

func knapsack(b []int, t int) (int, int) {
	dp := make([]int, t+1)
	num := make([]int, t+1)

	for _, v := range b {
		for j := v; j <= t; j++ {
			if dp[j] < dp[j-v]+v {
				dp[j] = dp[j-v] + v
				num[j] = num[j-v] + 1
			}
		}
	}
	return dp[t], num[t]
}

func main() {
	in, _ := os.Open("10465.in")
	defer in.Close()
	out, _ := os.Create("10465.out")
	defer out.Close()

	var m, n, t int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d", &m, &n, &t); err != nil {
			break
		}
		dp, num := knapsack([]int{m, n}, t)
		if dp == t {
			fmt.Fprintln(out, num)
		} else {
			fmt.Fprintln(out, num, t-dp)
		}
	}
}
