// UVa 10032 - Tug of War

package main

import (
	"fmt"
	"os"
)

func sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func solve(total int, weights []int) int {
	dp := make([]uint64, total+1)
	dp[0] = 1 << 0
	for i := range weights {
		for j := total; j >= 0; j-- {
			if dp[j] != 0 {
				dp[j+weights[i]] |= (dp[j] << 1)
			}
		}
	}

	var bit uint64 = 1 << (uint(len(weights) / 2))
	for i, j := total/2, total/2; i > -1 && j <= total; i, j = i-1, j+1 {
		if dp[i]&bit != 0 {
			return i
		}
		if dp[j]&bit != 0 {
			return j
		}
	}
	return 0
}

func main() {
	in, _ := os.Open("10032.in")
	defer in.Close()
	out, _ := os.Create("10032.out")
	defer out.Close()

	var kase, n int
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanf(in, "\n%d", &n)
		weights := make([]int, n)
		total := 0
		for i := range weights {
			fmt.Fscanf(in, "%d", &weights[i])
			total += weights[i]
		}
		mid := solve(total, weights)
		team1, team2 := sort(mid, total-mid)
		fmt.Fprintln(out, team1, team2)
		if kase--; kase != 0 {
			fmt.Fprintln(out)
		}
	}
}
