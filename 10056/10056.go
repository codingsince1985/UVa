// UVa 10056 - What is the Probability ?

package main

import (
	"fmt"
	"math"
	"os"
)

func calc(n, k int, p float64) float64 {
	switch {
	case p == 0:
		return 0
	case k == 1:
		return p / (1 - math.Pow(1-p, float64(n)))
	default:
		return p * math.Pow(1-p, float64(k-1)) / (1 - math.Pow(1-p, float64(n)))
	}
}

func main() {
	in, _ := os.Open("10056.in")
	defer in.Close()
	out, _ := os.Create("10056.out")
	defer out.Close()

	var s, n, k int
	var p float64
	fmt.Fscanf(in, "%d", &s)
	for s > 0 {
		fmt.Fscanf(in, "%d%f%d", &n, &p, &k)
		fmt.Fprintf(out, "%.4f\n", calc(n, k, p))
		s--
	}
}
