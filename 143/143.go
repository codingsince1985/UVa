// UVa 143 - Orchard Trees

package main

import (
	"fmt"
	"os"
)

var zero = point{0, 0}

type (
	point    struct{ x, y float64 }
	triangle struct{ p1, p2, p3 point }
)

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func area(p1, p2, p3 point) float64 {
	return abs(0.5 * ((p3.y-p1.y)*(p2.x-p1.x) - (p2.y-p1.y)*(p3.x-p1.x)))
}

func (t triangle) contains(p point) bool {
	return abs(area(p, t.p1, t.p2)+area(p, t.p1, t.p3)+area(p, t.p2, t.p3)-area(t.p1, t.p2, t.p3)) <= 0.00001
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 { return a + b - min(a, b) }

func solve(t triangle) int {
	minx := int(min(99, min(min(t.p1.x, t.p2.x), t.p3.x)))
	miny := int(min(99, min(min(t.p1.y, t.p2.y), t.p3.y)))
	maxx := int(max(0, max(max(t.p1.x, t.p2.x), t.p3.x)))
	maxy := int(max(0, max(max(t.p1.y, t.p2.y), t.p3.y)))
	var count int
	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			if t.contains(point{float64(x), float64(y)}) {
				count++
			}
		}
	}
	return count
}

func main() {
	in, _ := os.Open("143.in")
	defer in.Close()
	out, _ := os.Create("143.out")
	defer out.Close()

	var t triangle
	for {
		if fmt.Fscanf(in, "%f%f%f%f%f%f", &t.p1.x, &t.p1.y, &t.p2.x, &t.p2.y, &t.p3.x, &t.p3.y); t.p1 == zero && t.p2 == zero && t.p3 == zero {
			break
		}
		fmt.Fprintf(out, "%4d\n", solve(t))
	}
}
