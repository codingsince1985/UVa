// UVa 10566 - Crossed Ladders

package main

import (
	"fmt"
	"math"
	"os"
)

const zero = 0.00001

func binarySearch(x, y, c float64) float64 {
	var mid float64
	low, high := 0.0, min(x, y)
	for high-low > zero {
		if mid = (low + high) / 2; c/math.Sqrt(x*x-mid*mid)+c/math.Sqrt(y*y-mid*mid) < 1 {
			low = mid
		} else {
			high = mid
		}
	}
	return mid
}

func main() {
	in, _ := os.Open("10566.in")
	defer in.Close()
	out, _ := os.Create("10566.out")
	defer out.Close()

	var x, y, c float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f%f", &x, &y, &c); err != nil {
			break
		}
		fmt.Fprintf(out, "%.3f\n", binarySearch(x, y, c))
	}
}
