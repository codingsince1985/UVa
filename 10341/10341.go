// UVa 10341 - Solve It

package main

import (
	"fmt"
	"math"
	"os"
)

const zero = 0.00001

var p, q, r, s, t, u float64

func f(x float64) float64 {
	return p*math.Pow(math.E, -x) + q*math.Sin(x) + r*math.Cos(x) + s*math.Tan(x) + t*x*x + u
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func binarySearch() string {
	l, r := 0.0, 1.0
	fl, fr := f(l), f(r)
	if fl*fr > 0 {
		return "No solution"
	}
	for m := 0.5; ; m = (l + r) / 2 {
		fm := f(m)
		if abs(fm) <= zero {
			return fmt.Sprintf("%.4f", m)
		}
		if fl*fm > 0 {
			l = m
		} else {
			r = m
		}
	}
}

func main() {
	in, _ := os.Open("10341.in")
	defer in.Close()
	out, _ := os.Create("10341.out")
	defer out.Close()

	for {
		if _, err := fmt.Fscanf(in, "%f%f%f%f%f%f", &p, &q, &r, &s, &t, &u); err != nil {
			break
		}
		fmt.Fprintln(out, binarySearch())
	}
}
