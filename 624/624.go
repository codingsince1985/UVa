// UVa 624 - CD

package main

import (
	"fmt"
	"os"
)

func knapsack(n int, tracks []int) (int, []int) {
	dp := make([]int, n+1)
	path := make([][]int, n+1)
	for i, vi := range tracks {
		for j := n; j >= vi; j-- {
			if dp[j] < dp[j-vi]+vi {
				dp[j] = dp[j-vi] + vi
				path[j] = append(path[j-vi], i)
			}
		}
	}
	return dp[n], path[n]
}

func main() {
	in, _ := os.Open("624.in")
	defer in.Close()
	out, _ := os.Create("624.out")
	defer out.Close()

	var n, t int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &t); err != nil {
			break
		}
		tracks := make([]int, t)
		for i := range tracks {
			fmt.Fscanf(in, "%d", &tracks[i])
		}
		max, path := knapsack(n, tracks)
		for _, vi := range path {
			fmt.Fprint(out, tracks[vi], " ")
		}
		fmt.Fprintf(out, "sum:%d\n", max)
	}
}
