// UVa 10287 - Gifts in a Hexagonal Box

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	sqrt_3  = math.Sqrt(3)
	sqrt_21 = math.Sqrt(21)
)

func solve(s float64) [4]float64 {
	r1 := s * sqrt_3 / 2
	r2 := s / (1 + 2*sqrt_3/3)
	r4 := s * (6*sqrt_21 - 21) / (10 * sqrt_3)
	return [4]float64{r1, r2, r1 / 2, r4}
}

func main() {
	in, _ := os.Open("10287.in")
	defer in.Close()
	out, _ := os.Create("10287.out")
	defer out.Close()

	var s float64
	for {
		if _, err := fmt.Fscanf(in, "%f", &s); err != nil {
			break
		}
		rs := solve(s)
		fmt.Fprintf(out, "%.10f %.10f %.10f %.10f\n", rs[0], rs[1], rs[2], rs[3])
	}
}
