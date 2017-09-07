// UVa 11015 - 05-2 Rendezvous

package main

import (
	"fmt"
	"math"
	"os"
)

const max = 1001

func floydWarshall(matrix [][]int) {
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				if cost := matrix[i][k] + matrix[k][j]; cost < matrix[i][j] {
					matrix[i][j] = cost
				}
			}
		}
	}
}

func solve(matrix [][]int) int {
	min := math.MaxInt32
	var idx int
	floydWarshall(matrix)
	for i := range matrix {
		var total int
		for j := range matrix {
			if i != j {
				total += matrix[i][j]
			}
		}
		if total < min {
			min, idx = total, i
		}
	}
	return idx
}

func main() {
	in, _ := os.Open("11015.in")
	defer in.Close()
	out, _ := os.Create("11015.out")
	defer out.Close()

	var n, m, m1, m2, cost int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		members := make([]string, n)
		for i := range members {
			fmt.Fscanf(in, "%s", &members[i])
		}
		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, n)
			for j := range matrix[i] {
				matrix[i][j] = max
			}
		}
		for ; m > 0; m-- {
			fmt.Fscanf(in, "%d%d%d", &m1, &m2, &cost)
			matrix[m1-1][m2-1], matrix[m2-1][m1-1] = cost, cost
		}
		fmt.Fprintf(out, "Case #%d : %s\n", kase, members[solve(matrix)])
	}
}
