// UVa 10283 - The Kissing Circles

package main

import (
	"fmt"
	"math"
	"os"
)

func solve(r, n float64) (float64, float64, float64) {
	if n == 1 {
		return 0, 0, 0
	}
	angle := 180 / n
	rs := r * (math.Sin(angle * math.Pi / 180)) / (1 + math.Sin(angle*math.Pi/180))
	area1 := (r - rs) * math.Cos(angle*math.Pi/180) * rs
	area2 := (90 - angle) / 180 * math.Pi * rs * rs
	blue := n * (area1 - area2)
	return rs, blue, math.Pi*r*r - n*math.Pi*rs*rs - blue
}

func main() {
	in, _ := os.Open("10283.in")
	defer in.Close()
	out, _ := os.Create("10283.out")
	defer out.Close()

	var r, n float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f", &r, &n); err != nil {
			break
		}
		a1, a2, a3 := solve(r, n)
		fmt.Fprintf(out, "%.10f %.10f %.10f\n", a1, a2, a3)
	}
}
