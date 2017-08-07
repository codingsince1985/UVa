// UVa 534 - Frogger

package main

import (
	"fmt"
	"math"
	"os"
)

type stone struct{ x, y float64 }

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 { return a + b - min(a, b) }

func floydWarshall(n int, stones []stone) float64 {
	minMax := make([][]float64, n)
	for i := range minMax {
		minMax[i] = make([]float64, n)
		for j := range minMax[i] {
			minMax[i][j] = math.Sqrt((stones[i].x-stones[j].x)*(stones[i].x-stones[j].x) + (stones[i].y-stones[j].y)*(stones[i].y-stones[j].y))
		}
	}
	for k := range stones {
		for i := range stones {
			for j := range stones {
				minMax[i][j] = min(minMax[i][j], max(minMax[i][k], minMax[k][j]))
			}
		}
	}
	return minMax[0][1]
}

func main() {
	in, _ := os.Open("534.in")
	defer in.Close()
	out, _ := os.Create("534.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		stones := make([]stone, n)
		for i := range stones {
			fmt.Fscanf(in, "%f%f", &stones[i].x, &stones[i].y)
		}
		fmt.Fscanln(in)
		fmt.Fprintf(out, "Scenario #%d\n", kase)
		fmt.Fprintf(out, "Frog Distance = %.3f\n\n", floydWarshall(n, stones))
	}
}
