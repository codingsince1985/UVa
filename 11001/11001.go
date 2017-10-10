// UVa 11001 - Necklace

package main

import (
	"fmt"
	"math"
	"os"
)

const zero = 0.00001

func abs(a float64) float64 {
	if a >= 0 {
		return a
	}
	return -a
}

func solve(vt, v0 float64) int {
	var maxLength float64
	var maxDisc int
	for i := 1; vt > v0*float64(i); i++ {
		length := float64(i) * 0.3 * math.Sqrt(vt/float64(i)-v0)
		if abs(maxLength-length) <= zero {
			return 0
		}
		if length > maxLength {
			maxLength, maxDisc = length, i
		}
	}
	return maxDisc
}

func main() {
	in, _ := os.Open("11001.in")
	defer in.Close()
	out, _ := os.Create("11001.out")
	defer out.Close()

	var vt, v0 float64
	for {
		if fmt.Fscanf(in, "%f%f", &vt, &v0); vt == 0 && v0 == 0 {
			break
		}
		fmt.Fprintln(out, solve(vt, v0))
	}
}
