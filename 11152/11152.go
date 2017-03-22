// UVa 11152 - Colourful Flowers

package main

import (
	"fmt"
	"math"
	"os"
)

func area(a, b, c float64) float64 {
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func solve(a, b, c float64) (float64, float64, float64) {
	violetsAndRoses := area(a, b, c)
	r := violetsAndRoses * 2 / (a + b + c)
	roses := r * r * math.Acos(-1)
	r1 := a * b * c / (4 * violetsAndRoses)
	sunflowers := r1*r1*math.Acos(-1) - violetsAndRoses
	return sunflowers, violetsAndRoses - roses, roses
}

func main() {
	in, _ := os.Open("11152.in")
	defer in.Close()
	out, _ := os.Create("11152.out")
	defer out.Close()

	var a, b, c float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f%f", &a, &b, &c); err != nil {
			break
		}
		a1, a2, a3 := solve(a, b, c)
		fmt.Fprintf(out, "%.4f %.4f %.4f\n", a1, a2, a3)
	}
}
