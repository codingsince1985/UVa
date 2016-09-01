// UVa 10198 - Counting

package main

import (
	"fmt"
	"math/big"
	"os"
)

var (
	big2 = big.NewInt(2)
	dp   = func() []big.Int {
		dp := make([]big.Int, 1001)
		dp[0].SetInt64(1)
		dp[1].SetInt64(2)
		dp[2].SetInt64(5)
		for i := 3; i <= 1000; i++ {
			dp[i].Mul(&dp[i-1], big2)
			dp[i].Add(&dp[i], &dp[i-2])
			dp[i].Add(&dp[i], &dp[i-3])
		}
		return dp
	}()
)

func main() {
	in, _ := os.Open("10198.in")
	defer in.Close()
	out, _ := os.Create("10198.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, &dp[n])
	}
}
