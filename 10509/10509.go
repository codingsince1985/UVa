// UVa 10509 - R U Kidding Mr. Feynman?

package main

import (
	"fmt"
	"os"
)

func binarySearch(n float64) int {
	low, high := 1, 1000
	for low+1 < high {
		mid := (low + high) / 2
		if float64(mid*mid*mid) > n {
			high = mid - 1
		} else {
			low = mid
		}
	}
	if float64(high*high*high) > n {
		return low
	} else {
		return high
	}
}

func main() {
	in, _ := os.Open("10509.in")
	defer in.Close()
	out, _ := os.Create("10509.out")
	defer out.Close()

	var n float64
	for {
		if fmt.Fscanf(in, "%f", &n); n == 0 {
			break
		}
		a := float64(binarySearch(n))
		dx := (n - a*a*a) / (3 * a * a)
		fmt.Fprintf(out, "%.4f\n", a+dx)
	}
}
