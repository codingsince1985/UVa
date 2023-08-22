// UVa 10306 - e-Coins

package main

import (
	"fmt"
	"math"
	"os"
)

type ecoin struct{ conventional, infoTechnological int }

func solve(s int, ecoins []ecoin) int {
	dp := make([][]int, s+1)
	for i := range dp {
		dp[i] = make([]int, s+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	for _, coin := range ecoins {
		for i := coin.conventional; i <= s; i++ {
			for j := coin.infoTechnological; j <= s; j++ {
				dp[i][j] = min(dp[i][j], dp[i-coin.conventional][j-coin.infoTechnological]+1)
			}
		}
	}
	coins := math.MaxInt32
	for i := 0; i <= s; i++ {
		for j := 0; j <= s; j++ {
			if i*i+j*j == s*s {
				coins = min(coins, dp[i][j])
			}
		}
	}
	return coins
}

func main() {
	in, _ := os.Open("10306.in")
	defer in.Close()
	out, _ := os.Create("10306.out")
	defer out.Close()

	var n, m, s int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &m, &s)
		ecoins := make([]ecoin, m)
		for i := range ecoins {
			fmt.Fscanf(in, "%d%d", &ecoins[i].conventional, &ecoins[i].infoTechnological)
		}
		if coins := solve(s, ecoins); coins == math.MaxInt32 {
			fmt.Fprintln(out, "not possible")
		} else {
			fmt.Fprintln(out, coins)
		}
	}
}
