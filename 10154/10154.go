// UVa 10154 - Weights and Measures

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type turtle struct{ weight, strength int }

func lis(turtles []turtle) int {
	sort.Slice(turtles, func(i, j int) bool { return turtles[i].strength < turtles[j].strength })
	length := len(turtles)
	dp := make([]int, length+1)
	for i := 1; i <= length; i++ {
		dp[i] = math.MaxInt32
	}
	maxLen := 0
	for i := 1; i <= length; i++ {
		for j := length; j >= 1; j-- {
			if turtles[i-1].strength-turtles[i-1].weight >= dp[j-1] {
				if dp[j]-dp[j-1] >= turtles[i-1].weight {
					dp[j] = dp[j-1] + turtles[i-1].weight
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

	var turtles []turtle
	var t turtle
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &t.weight, &t.strength); err != nil {
			break
		}
		turtles = append(turtles, t)
	}
	fmt.Fprintln(out, lis(turtles))
}
