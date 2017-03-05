// UVa 681 - Convex Hull Finding

package main

import (
	"fmt"
	"os"
	"sort"
)

type point struct{ x, y int }

func cross(p, p1, p2 point) int { return (p1.x-p.x)*(p2.y-p.y) - (p1.y-p.y)*(p2.x-p.x) }

func convexHull(n int, points []point) []point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].y == points[j].y {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})
	m := 0
	ch := make([]point, n)
	for i := 0; i < n; i++ {
		for m >= 2 && cross(ch[m-2], ch[m-1], points[i]) <= 0 {
			m--
		}
		ch[m] = points[i]
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

func main() {
	in, _ := os.Open("681.in")
	defer in.Close()
	out, _ := os.Create("681.out")
	defer out.Close()

	var k, n int
	fmt.Fscanf(in, "%d", &k)
	fmt.Fprintln(out, k)
	for ; k > 0; k-- {
		fmt.Fscanf(in, "%d", &n)
		points := make([]point, n)
		for i := range points {
			fmt.Fscanf(in, "%d%d", &points[i].x, &points[i].y)
		}
		ch := convexHull(n, points)
		fmt.Fprintln(out, len(ch))
		for _, p := range ch {
			fmt.Fprintln(out, p.x, p.y)
		}
		if k > 1 {
			fmt.Fscanf(in, "%d", &n)
			fmt.Fprintln(out, n)
		}
	}
}
