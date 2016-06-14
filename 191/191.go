// UVa 191 - Intersection

package main

import (
	"fmt"
	"os"
)

type point struct{ x, y int }

type line struct{ p1, p2 point }

type rectangle struct{ p1, p2, p3, p4 point }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func area(p1, p2, p3 point) int { return (p3.y-p1.y)*(p2.x-p1.x) - (p2.y-p1.y)*(p3.x-p1.x) }

func cross(p point, l line) bool {
	return (p.x-l.p1.x)*(p.x-l.p2.x) <= 0 && (p.y-l.p1.y)*(p.y-l.p2.y) <= 0
}

func intersect(l1, l2 line) bool {
	a1 := area(l1.p1, l2.p1, l1.p2)
	a2 := area(l1.p1, l2.p2, l1.p2)
	a3 := area(l2.p1, l1.p1, l2.p2)
	a4 := area(l2.p1, l1.p2, l2.p2)

	switch {
	case a1*a2 < 0 && a3*a4 < 0:
		fallthrough
	case a1 == 0 && cross(l2.p1, l1):
		fallthrough
	case a2 == 0 && cross(l2.p2, l1):
		fallthrough
	case a3 == 0 && cross(l1.p1, l2):
		fallthrough
	case a4 == 0 && cross(l1.p2, l2):
		return true
	default:
		return false
	}
}

func between(a, min, max int) bool { return a >= min && a <= max }

func main() {
	in, _ := os.Open("191.in")
	defer in.Close()
	out, _ := os.Create("191.out")
	defer out.Close()

	var n, x1, y1, x2, y2 int
	fmt.Fscanf(in, "%d", &n)
	for n > 0 {
		n--
		fmt.Fscanf(in, "%d%d%d%d", &x1, &y1, &x2, &y2)
		l := line{point{x1, y1}, point{x2, y2}}
		fmt.Fscanf(in, "%d%d%d%d", &x1, &y1, &x2, &y2)
		x1, x2 := min(x1, x2), max(x1, x2)
		y1, y2 := max(y1, y2), min(y1, y2)
		r := rectangle{point{x1, y1}, point{x2, y1}, point{x2, y2}, point{x1, y2}}

		var ok bool
		switch {
		case intersect(l, line{r.p1, r.p2}):
			fallthrough
		case intersect(l, line{r.p2, r.p3}):
			fallthrough
		case intersect(l, line{r.p3, r.p4}):
			fallthrough
		case intersect(l, line{r.p4, r.p1}):
			fallthrough
		case between(l.p1.x, x1, x2) && between(l.p2.x, x1, x2) && between(l.p1.y, y2, y1) && between(l.p2.y, y2, y1):
			ok = true
		}
		if ok {
			fmt.Fprintln(out, "T")
		} else {
			fmt.Fprintln(out, "F")
		}
	}
}
