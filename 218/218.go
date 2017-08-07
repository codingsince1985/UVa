// UVa 218 - Moth Eradication

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type point struct{ x, y float64 }

func distance(p1, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func cross(p, p1, p2 point) float64 { return (p1.x-p.x)*(p2.y-p.y) - (p1.y-p.y)*(p2.x-p.x) }

func convexHull(n int, points []point) []point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].y == points[j].y {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})
	m := 0
	ch := make([]point, n+1)
	for _, point := range points {
		for m >= 2 && cross(ch[m-2], ch[m-1], point) <= 0 {
			m--
		}
		ch[m] = point
		m++
	}
	for i, t := n-1, m+1; i >= 0; i-- {
		for m >= t && cross(ch[m-2], ch[m-1], points[i]) <= 0 {
			m--
		}
		ch[m] = points[i]
		m++
	}
	return ch[:m]
}

func output(out *os.File, points []point) {
	var length float64
	for i := len(points) - 1; i >= 0; i-- {
		fmt.Fprintf(out, "(%.1f,%.1f)", points[i].x, points[i].y)
		if i > 0 {
			fmt.Fprint(out, "-")
			length += distance(points[i], points[i-1])
		} else {
			fmt.Fprintln(out)
		}
	}
	fmt.Fprintf(out, "Perimeter length = %.2f\n", length)
}

func main() {
	in, _ := os.Open("218.in")
	defer in.Close()
	out, _ := os.Create("218.out")
	defer out.Close()

	var n int
	first := true
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		points := make([]point, n)
		for i := range points {
			fmt.Fscanf(in, "%f%f", &points[i].x, &points[i].y)
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Region #%d:\n", kase)
		output(out, convexHull(n, points))
	}
}
