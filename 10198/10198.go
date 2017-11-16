// UVa 10198 - Counting

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 1000

var (
	two = big.NewInt(2)
	dp  = func() []big.Int {
		dp := make([]big.Int, max+1)
		dp[0].SetInt64(1)
		dp[1].SetInt64(2)
		dp[2].SetInt64(5)
		for i := 3; i <= max; i++ {
			dp[i].Mul(&dp[i-1], two).Add(&dp[i], &dp[i-2]).Add(&dp[i], &dp[i-3])
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
