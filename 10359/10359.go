// UVa 10359 - Tiling

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 250

var cache = func() []big.Int {
	dp := make([]big.Int, max+1)
	dp[0].SetInt64(1)
	dp[1].SetInt64(1)
	for i := 2; i <= max; i++ {
		tmp := big.NewInt(2)
		dp[i] = *tmp.Mul(tmp, &dp[i-2]).Add(tmp, &dp[i-1])
	}
	return dp
}()

func main() {
	in, _ := os.Open("10359.in")
	defer in.Close()
	out, _ := os.Create("10359.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, &cache[n])
	}
}
