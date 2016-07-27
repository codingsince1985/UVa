// UVa 378 - Intersecting Lines

package main

import (
	"fmt"
	"os"
)

var out *os.File

type (
	point struct{ x, y float64 }
	line  struct{ s, e point }
)

func solve(l1, l2 line) {
	a1, b1 := l1.s.y-l1.e.y, -l1.s.x+l1.e.x
	a2, b2 := l2.s.y-l2.e.y, -l2.s.x+l2.e.x
	c1, c2 := a1*l1.s.x+b1*l1.s.y, a2*l2.s.x+b2*l2.s.y
	d := a1*b2 - a2*b1
	dx, dy := c1*b2-c2*b1, a1*c2-a2*c1

	switch {
	case d == 0:
		switch {
		case dx == 0 && dy == 0:
			fmt.Fprintln(out, "LINE")
		default:
			fmt.Fprintln(out, "NONE")
		}
	default:
		fmt.Fprintf(out, "POINT %.2f %.2f\n", dx/d, dy/d)
	}
}

func main() {
	in, _ := os.Open("378.in")
	defer in.Close()
	out, _ = os.Create("378.out")
	defer out.Close()

	var n int
	var x, y float64
	fmt.Fscanf(in, "%d", &n)
	fmt.Fprintln(out, "INTERSECTING LINES OUTPUT")
	for n > 0 {
		fmt.Fscanf(in, "%f%f", &x, &y)
		p1 := point{x, y}
		fmt.Fscanf(in, "%f%f", &x, &y)
		p2 := point{x, y}
		l1 := line{p1, p2}
		fmt.Fscanf(in, "%f%f", &x, &y)
		p1 = point{x, y}
		fmt.Fscanf(in, "%f%f", &x, &y)
		p2 = point{x, y}
		l2 := line{p1, p2}
		solve(l1, l2)
		n--
	}
	fmt.Fprintln(out, "END OF OUTPUT")
}
