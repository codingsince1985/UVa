// UVa 634 - Polygon

package main

import (
	"fmt"
	"os"
)

type point struct{ x, y int }

func solve(n int, p point, vertices []point) int {
	var count int
	for i, j := 0, n-1; i < n; i, j = i+1, i {
		if (vertices[i].y-p.y)*(vertices[j].y-p.y) <= 0 &&
			p.x-vertices[i].x < (vertices[j].x-vertices[i].x)*(p.y-vertices[i].y)/(vertices[j].y-vertices[i].y) {
			count++
		}
	}
	return count
}

func main() {
	in, _ := os.Open("634.in")
	defer in.Close()
	out, _ := os.Create("634.out")
	defer out.Close()

	var n int
	var p point
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		vertices := make([]point, n)
		for i := range vertices {
			fmt.Fscanf(in, "%d%d", &vertices[i].x, &vertices[i].y)
		}
		fmt.Fscanf(in, "%d%d", &p.x, &p.y)
		if solve(n, p, vertices)%2 == 1 {
			fmt.Fprintln(out, "T")
		} else {
			fmt.Fprintln(out, "F")
		}
	}
}
