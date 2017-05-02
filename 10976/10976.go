// UVa 10976 - Fractions Again?!

package main

import (
	"fmt"
	"os"
)

var out *os.File

type solution struct{ x, y int }

func solve(k int) {
	var solutions []solution
	for y := k + 1; y <= k*2; y++ {
		if (y*k)%(y-k) == 0 {
			solutions = append(solutions, solution{(y * k) / (y - k), y})
		}
	}
	fmt.Fprintln(out, len(solutions))
	for _, s := range solutions {
		fmt.Fprintf(out, "1/%d = 1/%d + 1/%d\n", k, s.x, s.y)
	}
}

func main() {
	in, _ := os.Open("10976.in")
	defer in.Close()
	out, _ = os.Create("10976.out")
	defer out.Close()

	var k int
	for {
		if _, err := fmt.Fscanf(in, "%d", &k); err != nil {
			break
		}
		solve(k)
	}
}
