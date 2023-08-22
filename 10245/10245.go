// UVa 10245 - The Closest Pair Problem

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

func divide(p []point) float64 {
	if len(p) <= 1 {
		return math.MaxFloat64
	}
	mid := len(p) / 2
	left, right := divide(p[:mid]), divide(p[mid:])
	return conquer(p, min(left, right))
}

func conquer(p []point, min float64) float64 {
	mid := len(p) / 2
	midx := (p[mid-1].x + p[mid].x) / 2
	var mids []point
	for _, vi := range p {
		if vi.x >= midx-min && vi.x <= midx+min {
			mids = append(mids, vi)
		}
	}
	for i := range mids {
		for j := i + 1; j < len(mids); j++ {
			dist := distance(mids[i], mids[j])
			if dist < min {
				min = dist
			}
		}
	}
	return min
}

func main() {
	in, _ := os.Open("10245.in")
	defer in.Close()
	out, _ := os.Create("10245.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		p := make([]point, n)
		for i := range p {
			fmt.Fscanf(in, "%f%f", &p[i].x, &p[i].y)
		}
		sort.Slice(p, func(i, j int) bool { return p[i].x < p[j].x })
		if dist := divide(p); dist >= 10000 {
			fmt.Fprintln(out, "INFINITY")
		} else {
			fmt.Fprintf(out, "%.4f\n", dist)
		}
	}
}
