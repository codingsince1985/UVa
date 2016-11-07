// UVa 10137 - The Trip

package main

import (
	"fmt"
	"math"
	"os"
)

func round(amt float64) float64 {
	return math.Floor(amt*100+0.5) / 100
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
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
		total := 0.0
		lst := make([]float64, n)
		for i := range lst {
			fmt.Fscanf(in, "%f", &lst[i])
			total += lst[i]
		}
		avg := total / float64(n)

		var diffPos, diffNeg float64
		for i := range lst {
			if lst[i] > avg {
				diffPos += round(lst[i] - avg)
			} else {
				diffNeg += round(avg - lst[i])
			}
		}
		fmt.Fprintf(out, "$%.2f\n", min(diffPos, diffNeg))
	}
}
