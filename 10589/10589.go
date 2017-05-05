// UVa 10589 - Area

package main

import (
	"fmt"
	"os"
)

type point struct{ x, y float64 }

func distanceGreaterThan(p1, p2 point, a float64) bool {
	return (p1.x-p2.x)*(p1.x-p2.x)+(p1.y-p2.y)*(p1.y-p2.y) > a*a
}

func solve(a float64, points []point) float64 {
	corners := []point{{0, 0}, {0, a}, {a, 0}, {a, a}}
	var m int
here:
	for _, p := range points {
		for _, c := range corners {
			if distanceGreaterThan(p, c, a) {
				continue here
			}
		}
		m++
	}
	return float64(m) * a * a / float64(len(points))
}

func main() {
	in, _ := os.Open("10589.in")
	defer in.Close()
	out, _ := os.Create("10589.out")
	defer out.Close()

	var n int
	var a float64
	for {
		if fmt.Fscanf(in, "%d%f", &n, &a); n == 0 && a == 0 {
			break
		}
		points := make([]point, n)
		for i := range points {
			fmt.Fscanf(in, "%f%f", &points[i].x, &points[i].y)
		}
		fmt.Fprintf(out, "%.5f\n", solve(a, points))
	}
}
