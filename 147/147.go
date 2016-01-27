// UVa 147 - Dollars

package main

import (
	"fmt"
	"os"
)

var cs = []int{5, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}

func main() {
	in, _ := os.Open("147.in")
	defer in.Close()
	out, _ := os.Create("147.out")
	defer out.Close()

	var v float64
	var vi int
	var dp []int
	for {
		if fmt.Fscanf(in, "%f", &v); v == 0 {
			break
		}
		vi = int(v * 100);
		dp = make([]int, vi + 1)
		dp[0] = 1

		for i := range cs {
			c := cs[i]
			for j := c; j <= vi; j++ {
				dp[j] += dp[j - c]
			}
		}
		fmt.Fprintf(out, "%6.2f%17d\n", v, dp[vi])
	}
}
