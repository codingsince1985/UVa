// UVa 688 - Mobile Phone Coverage

package main

import (
	"fmt"
	"os"
	"sort"
)

type antenna struct{ x, y, r float64 }

func solve(n int, antennas []antenna) float64 {
	x, y := make([]float64, 2*n), make([]float64, 2*n)
	for i, a := range antennas {
		x[2*i], x[2*i+1] = a.x-a.r, a.x+a.r
		y[2*i], y[2*i+1] = a.y-a.r, a.y+a.r
	}
	sort.Float64s(x)
	sort.Float64s(y)
	var area float64
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			for _, a := range antennas {
				if a.x-a.r <= x[i] && a.x+a.r >= x[i+1] && a.y-a.r <= y[j] && a.y+a.r >= y[j+1] {
					area += (x[i+1] - x[i]) * (y[j+1] - y[j])
					break
				}
			}
		}
	}
	return area
}

func main() {
	in, _ := os.Open("688.in")
	defer in.Close()
	out, _ := os.Create("688.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		antennas := make([]antenna, n)
		for i := range antennas {
			fmt.Fscanf(in, "%f%f%f", &antennas[i].x, &antennas[i].y, &antennas[i].r)
		}
		fmt.Fprintf(out, "%d %.02f\n", kase, solve(n, antennas))
	}
}
