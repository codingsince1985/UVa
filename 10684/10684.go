// UVa 10684 - The jackpot

package main

import (
	"fmt"
	"os"
)

func maxSubArray(bets []int) int {
	max, meh := 0, 0
	for _, bet := range bets {
		meh += bet
		if meh < 0 {
			meh = 0
		}
		if meh > max {
			max = meh
		}
	}
	return max
}

func main() {
	in, _ := os.Open("10684.in")
	defer in.Close()
	out, _ := os.Create("10684.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		bets := make([]int, n)
		for i := range bets {
			fmt.Fscanf(in, "%d", &bets[i])
		}
		max := maxSubArray(bets)
		if max > 0 {
			fmt.Fprintf(out, "The maximum winning streak is %d.\n", max)
		} else {
			fmt.Fprintln(out, "Losing streak.")
		}
	}
}
