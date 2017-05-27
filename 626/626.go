// UVa 626 - Ecosystem

package main

import (
	"fmt"
	"os"
)

var out *os.File

func solve(matrix [][]bool) {
	var count int
	for i := range matrix {
		for j := range matrix {
			for k := range matrix {
				if matrix[i][j] && matrix[j][k] && matrix[k][i] && (i > j && j > k || i < j && j < k) {
					count++
					fmt.Fprintf(out, "%d %d %d\n", i+1, j+1, k+1)
				}
			}
		}
	}
	fmt.Fprintf(out, "total:%d\n\n", count)
}

func main() {
	in, _ := os.Open("626.in")
	defer in.Close()
	out, _ = os.Create("626.out")
	defer out.Close()

	var n, tmp int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		matrix := make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
			for j := range matrix[i] {
				fmt.Fscanf(in, "%d", &tmp)
				matrix[i][j] = tmp == 1
			}
		}
		solve(matrix)
	}
}
