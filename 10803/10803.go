// UVa 10803 - Thunder Mountain

package main

import (
	"fmt"
	"math"
	"os"
)

const inf = math.MaxFloat32

type town struct{ x, y float64 }

func distance(t1, t2 town) float64 { return math.Sqrt(math.Pow(t1.x-t2.x, 2) + math.Pow(t1.y-t2.y, 2)) }

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 { return a + b - min(a, b) }

func floydWarshall(t int, matrix [][]float64) {
	for k := 0; k < t; k++ {
		for i := 0; i < t; i++ {
			for j := 0; j < t; j++ {
				matrix[i][j] = min(matrix[i][j], matrix[i][k]+matrix[k][j])
			}
		}
	}
}

func findLongest(t int, matrix [][]float64) float64 {
	var longest float64
	for i := 0; i < t-1; i++ {
		for j := i + 1; j < t; j++ {
			longest = max(longest, matrix[i][j])
		}
	}
	return longest
}

func solve(t int, towns []town) float64 {
	matrix := make([][]float64, t)
	for i := range matrix {
		matrix[i] = make([]float64, t)
	}
	for i := range matrix {
		for j := i; j < t; j++ {
			if i == j {
				matrix[i][j] = inf
				continue
			}
			if dist := distance(towns[i], towns[j]); dist <= 10 {
				matrix[i][j], matrix[j][i] = dist, dist
			} else {
				matrix[i][j], matrix[j][i] = inf, inf
			}
		}
	}
	floydWarshall(t, matrix)
	return findLongest(t, matrix)
}

func main() {
	in, _ := os.Open("10803.in")
	defer in.Close()
	out, _ := os.Create("10803.out")
	defer out.Close()

	var n, t int
	fmt.Fscanf(in, "%d", &n)
	for kase := 1; kase <= n; kase++ {
		fmt.Fscanf(in, "%d", &t)
		towns := make([]town, t)
		for i := range towns {
			fmt.Fscanf(in, "%f%f", &towns[i].x, &towns[i].y)
		}
		fmt.Fprintf(out, "Case #%d:\n", kase)
		if dist := solve(t, towns); dist == inf {
			fmt.Fprintln(out, "Send Kurdy\n")
		} else {
			fmt.Fprintf(out, "%.4f\n\n", dist)
		}
	}
}
