// UVa 10078 - The Art Gallery

package main

import (
	"fmt"
	"os"
)

type point struct{ x, y int }

func cross(p, p1, p2 point) int { return (p1.x-p.x)*(p2.y-p.y) - (p1.y-p.y)*(p2.x-p.x) }

func isConvex(points []point) bool {
	points = append(points, points[:2]...)
	var pre int
	for i := 0; i < len(points)-2; i++ {
		curr := cross(points[i], points[i+1], points[i+2])
		if i > 0 && pre*curr < 0 {
			return false
		}
		pre = curr
	}
	return true
}

func main() {
	in, _ := os.Open("10078.in")
	defer in.Close()
	out, _ := os.Create("10078.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		points := make([]point, n)
		for i := range points {
			fmt.Fscanf(in, "%d%d", &points[i].x, &points[i].y)
		}
		if isConvex(points) {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
	}
}
