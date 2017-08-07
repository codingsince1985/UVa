// UVa 436 - Arbitrage (II)

package main

import (
	"fmt"
	"os"
)

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func floydWarshall(matrix [][]float64) {
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				matrix[i][j] = max(matrix[i][j], matrix[i][k]*matrix[k][j])
			}
		}
	}
}

func solve(matrix [][]float64) bool {
	floydWarshall(matrix)
	for i := range matrix {
		if matrix[i][i] > 1 {
			return true
		}
	}
	return false
}

func main() {
	in, _ := os.Open("436.in")
	defer in.Close()
	out, _ := os.Create("436.out")
	defer out.Close()

	var n, m int
	var r float64
	var c1, c2 string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		currencyMap := make(map[string]int)
		for i := 0; i < n; i++ {
			fmt.Fscanf(in, "%s", &c1)
			currencyMap[c1] = i
		}
		matrix := make([][]float64, n)
		for i := range matrix {
			matrix[i] = make([]float64, n)
		}
		for fmt.Fscanf(in, "%d", &m); m > 0; m-- {
			fmt.Fscanf(in, "%s%f%s", &c1, &r, &c2)
			matrix[currencyMap[c1]][currencyMap[c2]] = r
		}
		if fmt.Fprintf(out, "Case %d: ", kase); solve(matrix) {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
		fmt.Fscanln(in)
	}
}
