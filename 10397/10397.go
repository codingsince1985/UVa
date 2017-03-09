// UVa 10397 - Connect the Campus

package main

import (
	"fmt"
	"math"
	"os"
)

type building struct{ x, y int }

func distance(b1, b2 building) float64 {
	return math.Sqrt(float64((b1.x-b2.x)*(b1.x-b2.x) + (b1.y-b2.y)*(b1.y-b2.y)))
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func prim(n int, campus [][]float64) float64 {
	var total float64
	visited := make(map[int]bool)
	low := make([]float64, n)
	for i := range low {
		low[i] = math.MaxFloat64
	}
	visited[0], low[0] = true, 0
	for i := 0; i < n; i++ {
		var minP int
		minD := math.MaxFloat64
		for j := 0; j < n; j++ {
			if !visited[j] && minD > low[j] {
				minP = j
				minD = low[j]
			}
		}
		visited[minP] = true
		total += low[minP]
		for j := 0; j < n; j++ {
			low[j] = min(low[j], campus[minP][j])
		}
	}
	return total
}

func main() {
	in, _ := os.Open("10397.in")
	defer in.Close()
	out, _ := os.Create("10397.out")
	defer out.Close()

	var n, m, x, y int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		buildings := make([]building, n)
		for i := range buildings {
			fmt.Fscanf(in, "%d%d", &buildings[i].x, &buildings[i].y)
		}
		campus := make([][]float64, n)
		for i := range campus {
			campus[i] = make([]float64, n)
		}
		for i := range campus {
			for j := range campus[i] {
				campus[i][j] = distance(buildings[i], buildings[j])
				campus[j][i] = campus[i][j]
			}
		}
		for fmt.Fscanf(in, "%d", &m); m > 0; m-- {
			fmt.Fscanf(in, "%d%d", &x, &y)
			campus[x-1][y-1], campus[y-1][x-1] = 0, 0
		}
		fmt.Fprintf(out, "%.2f\n", prim(n, campus))
	}
}
