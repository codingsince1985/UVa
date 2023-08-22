// UVa 562 - Dividing coins

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("562.in")
	defer in.Close()
	out, _ := os.Create("562.out")
	defer out.Close()

	var n, m int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		sum := 0
		fmt.Fscanf(in, "%d", &m)
		coins := make([]int, m)
		for i := range coins {
			fmt.Fscanf(in, "%d", &coins[i])
			sum += coins[i]
		}
		avg := sum / 2
		dp := make([]int, avg+1)
		for _, coin := range coins {
			for j := avg; j >= coin; j-- {
				dp[j] = max(dp[j], dp[j-coin]+coin)
			}
		}
		fmt.Fprintln(out, sum-2*dp[avg])
	}
}
