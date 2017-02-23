// UVa 10301 - Rings and Glue

package main

import (
	"fmt"
	"math"
	"os"
)

type ring struct{ x, y, r float64 }

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}

func glued(r1, r2 ring) bool {
	distance := math.Sqrt((r1.x-r2.x)*(r1.x-r2.x) + (r1.y-r2.y)*(r1.y-r2.y))
	return distance < r1.r+r2.r && distance > abs(r1.r-r2.r)
}

func solve(n int, rings []ring) int {
	f, num := make([]int, n), make([]int, n)
	for i := range f {
		f[i], num[i] = i, 1
	}
	var max int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if glued(rings[i], rings[j]) {
				if f1, f2 := unionFind(i, f), unionFind(j, f); f1 != f2 {
					f[f1] = f2
					if num[f2] += num[f1]; num[f2] > max {
						max = num[f2]
					}
				}
			}
		}
	}
	return max
}

func main() {
	in, _ := os.Open("10301.in")
	defer in.Close()
	out, _ := os.Create("10301.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == -1 {
			break
		}
		rings := make([]ring, n)
		for i := range rings {
			fmt.Fscanf(in, "%f%f%f", &rings[i].x, &rings[i].y, &rings[i].r)
		}
		if max := solve(n, rings); max != 1 {
			fmt.Fprintf(out, "The largest component contains %d rings.\n", max)
		} else {
			fmt.Fprintln(out, "The largest component contains 1 ring.")
		}
	}
}
