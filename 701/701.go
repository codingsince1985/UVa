// UVa 701 - The Archeologists' Dilemma

package main

import (
	"fmt"
	"math"
	"os"
)

var log2_10 = math.Log2(10.0)

func solve(n float64) int {
	for digits := int(math.Log10(n)) + 2; ; digits++ {
		low := int(math.Log2(n) + float64(digits)*log2_10)
		high := int(math.Log2(n+1) + float64(digits)*log2_10)
		if low+1 == high {
			return high
		}
	}
}

func main() {
	in, _ := os.Open("701.in")
	defer in.Close()
	out, _ := os.Create("701.out")
	defer out.Close()

	var n float64
	for {
		if _, err := fmt.Fscanf(in, "%f", &n); err != nil {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}
