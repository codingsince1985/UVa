// UVa 190 - Circle Through Three Points

package main

import (
	"fmt"
	"math"
	"os"
)

type point struct{ x, y float64 }

func abs(a float64) float64 {
	if a >= 0 {
		return a
	}
	return -a
}

func sign(condition bool) string {
	if condition {
		return "-"
	}
	return "+"
}

func output(out *os.File, p []point) {
	a1, b1 := p[1].x-p[0].x, p[1].y-p[0].y
	c1 := p[1].x*p[1].x - p[0].x*p[0].x + p[1].y*p[1].y - p[0].y*p[0].y
	a2, b2 := p[2].x-p[0].x, p[2].y-p[0].y
	c2 := p[2].x*p[2].x - p[0].x*p[0].x + p[2].y*p[2].y - p[0].y*p[0].y

	d := (b1*c2 - b2*c1) / (a1*b2 - a2*b1)
	e := (a2*c1 - a1*c2) / (a1*b2 - a2*b1)

	x0 := -d / 2
	y0 := -e / 2
	r := math.Sqrt((p[1].x-x0)*(p[1].x-x0) + (p[1].y-y0)*(p[1].y-y0))
	fmt.Fprintf(out, "(x %s %.3f)^2 + (y %s %.3f)^2 = %.3f^2\n", sign(x0 > 0), abs(x0), sign(y0 > 0), abs(y0), r)

	f := -p[1].x*p[1].x - p[1].y*p[1].y - p[1].x*d - p[1].y*e
	fmt.Fprintf(out, "x^2 + y^2 %s %.3fx %s %.3fy %s %.3f = 0\n", sign(d < 0), abs(d), sign(e < 0), abs(e), sign(f < 0), abs(f))
}

func main() {
	in, _ := os.Open("190.in")
	defer in.Close()
	out, _ := os.Create("190.out")
	defer out.Close()

here:
	for {
		p := make([]point, 3)
		for i := range p {
			if _, err := fmt.Fscanf(in, "%f%f", &p[i].x, &p[i].y); err != nil {
				break here
			}
		}
		output(out, p)
		fmt.Fprintln(out)
	}
}
