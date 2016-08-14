// UVa 10276 - Hanoi Tower Troubles Again!

package main

import (
	"fmt"
	"math"
	"os"
)

func isSquareNumber(n int) bool {
	sqrt := int(math.Sqrt(float64(n)) + 0.5)
	return sqrt*sqrt == n
}

func hanoi(n int) int {
	pegs := make([]int, n)
	ball := 1
here:
	for {
		for i, vi := range pegs {
			if vi == 0 || isSquareNumber(vi+ball) {
				pegs[i] = ball
				ball++
				continue here
			}
		}
		return ball - 1
	}
}

func main() {
	in, _ := os.Open("10276.in")
	defer in.Close()
	out, _ := os.Create("10276.out")
	defer out.Close()

	var t, n int
	fmt.Fscanf(in, "%d", &t)
	for t > 0 {
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintln(out, hanoi(n))
		t--
	}
}
