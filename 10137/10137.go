// UVa 10137 - The Trip

package main

import (
	"fmt"
	"math"
	"os"
)

func round(amt float64) float64 { return math.Floor(amt*100+0.5) / 100 }

func solve(n int, lst []float64) float64 {
	var total float64
	for _, l := range lst {
		total += l
	}
	avg := total / float64(n)
	var diffPos, diffNeg float64
	for _, l := range lst {
		if l > avg {
			diffPos += round(l - avg)
		} else {
			diffNeg += round(avg - l)
		}
	}
	return min(diffPos, diffNeg)
}

func main() {
	in, _ := os.Open("10137.in")
	defer in.Close()
	out, _ := os.Create("10137.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		lst := make([]float64, n)
		for i := range lst {
			fmt.Fscanf(in, "%f", &lst[i])
		}
		fmt.Fprintf(out, "$%.2f\n", solve(n, lst))
	}
}
