// UVa 737 - Gleaming the Cubes

package main

import (
	"fmt"
	"math"
	"os"
)

type cube struct{ lx, ly, lz, hx, hy, hz int }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func solve(cubes []cube) int {
	maxx, maxy, maxz := math.MaxInt32, math.MaxInt32, math.MaxInt32
	minx, miny, minz := math.MinInt32, math.MinInt32, math.MinInt32
	for _, c := range cubes {
		maxx, maxy, maxz = min(maxx, c.hx), min(maxy, c.hy), min(maxz, c.hz)
		minx, miny, minz = max(minx, c.lx), max(miny, c.ly), max(minz, c.lz)
	}
	return max(0, (maxx-minx)*(maxy-miny)*(maxz-minz))
}

func main() {
	in, _ := os.Open("737.in")
	defer in.Close()
	out, _ := os.Create("737.out")
	defer out.Close()

	var n, tmp int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		cubes := make([]cube, n)
		for i := range cubes {
			fmt.Fscanf(in, "%d%d%d%d", &cubes[i].lx, &cubes[i].ly, &cubes[i].lz, &tmp)
			cubes[i].hx, cubes[i].hy, cubes[i].hz = cubes[i].lx+tmp, cubes[i].ly+tmp, cubes[i].lz+tmp
		}
		fmt.Fprintln(out, solve(cubes))
	}
}
