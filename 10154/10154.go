// UVa 10154 - Weights and Measures

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type (
	turtle  struct{ weight, strength int }
	turtles []turtle
)

func lis(ts turtles) int {
	sort.Slice(ts, func(i, j int) bool { return ts[i].strength < ts[j].strength })
	dp := make([]int, len(ts)+1)
	for i := 1; i <= len(ts); i++ {
		dp[i] = math.MaxInt32
	}
	maxLen := 0
	for i := 1; i <= len(ts); i++ {
		for j := len(ts); j >= 1; j-- {
			if ts[i-1].strength-ts[i-1].weight >= dp[j-1] {
				if dp[j]-dp[j-1] >= ts[i-1].weight {
					dp[j] = dp[j-1] + ts[i-1].weight
					if j > maxLen {
						maxLen = j
					}
				}
			}
		}
	}
	return maxLen
}

func main() {
	in, _ := os.Open("10154.in")
	defer in.Close()
	out, _ := os.Create("10154.out")
	defer out.Close()

	var ts turtles
	var t turtle
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &t.weight, &t.strength); err != nil {
			break
		}
		ts = append(ts, t)
	}
	fmt.Fprintln(out, lis(ts))
}
