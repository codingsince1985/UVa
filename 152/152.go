// UVa 152 - Tree's a Crowd

package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type point struct{ x, y, z int }

func distance(p1, p2 point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z)))
}

func solve(points []point) [10]int {
	var dist [10]int
	for i, vi := range points {
		closest := math.MaxFloat64
		for j, vj := range points {
			if i != j {
				d := distance(vi, vj)
				if d < closest {
					closest = d
				}
			}
		}
		if closest < 10 {
			dist[int(closest)]++
		}
	}
	return dist
}

func output(out io.Writer, dist [10]int) {
	for _, d := range dist {
		fmt.Fprintf(out, "%4d", d)
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("152.in")
	defer in.Close()
	out, _ := os.Create("152.out")
	defer out.Close()

	var p point
	var points []point
	for {
		if fmt.Fscanf(in, "%d%d%d", &p.x, &p.y, &p.z); p.x == 0 && p.y == 0 && p.z == 0 {
			break
		}
		points = append(points, p)
	}
	output(out, solve(points))
}
