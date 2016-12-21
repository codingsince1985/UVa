// UVa 10050 - Hartals

package main

import (
	"fmt"
	"os"
)

func solve(parties []int, days int) int {
	hartal := make(map[int]bool)
	for _, v := range parties {
		round := 1
		for v*round <= days {
			day := v * round
			if !hartal[day] && day%7 != 6 && day%7 != 0 {
				hartal[day] = true
			}
			round++
		}
	}
	return len(hartal)
}

func main() {
	in, _ := os.Open("10050.in")
	defer in.Close()
	out, _ := os.Create("10050.out")
	defer out.Close()

	var kase, days, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d\n%d", &days, &n)
		parties := make([]int, n)
		for i := range parties {
			fmt.Fscanf(in, "%d", &parties[i])
		}
		fmt.Fprintln(out, solve(parties, days))
	}
}
