// UVa 423 - MPI Maelstrom

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func floydWarshall(matrix [][]int) {
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				if v := matrix[i][k] + matrix[k][j]; v < matrix[i][j] {
					matrix[i][j] = v
				}
			}
		}
	}
}

func main() {
	in, _ := os.Open("423.in")
	defer in.Close()
	out, _ := os.Create("423.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	var token string
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if fmt.Fscanf(in, "%s", &token); token == "x" {
				matrix[i][j] = math.MaxInt32
			} else {
				matrix[i][j], _ = strconv.Atoi(token)
			}
			matrix[j][i] = matrix[i][j]
		}
	}
	floydWarshall(matrix)
	var max int
	for i := 1; i < n; i++ {
		if matrix[0][i] > max {
			max = matrix[0][i]
		}
	}
	fmt.Fprintln(out, max)
}
