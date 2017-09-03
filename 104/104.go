// UVa 104 - Arbitrage

package main

import (
	"fmt"
	"os"
)

var out *os.File

type node struct {
	v float64
	p []int
}

func output(s int, path []int) {
	fmt.Fprint(out, s+1)
	for _, v := range path {
		fmt.Fprint(out, " ", v+1)
	}
	fmt.Fprintln(out, "", s+1)
}

func floydWarshall(n int, matrix [][][]node) {
	for s := 1; s < n; s++ {
		for k := range matrix {
			for i := range matrix {
				for j := range matrix {
					if newV := matrix[i][k][s-1].v * matrix[k][j][0].v; newV > matrix[i][j][s].v {
						matrix[i][j][s] = node{newV, append(matrix[i][k][s-1].p, k)}
					}
				}
			}
		}
		for k := range matrix {
			if matrix[k][k][s].v > 1.01 {
				output(k, matrix[k][k][s].p)
				return
			}
		}
	}
	fmt.Fprintln(out, "no arbitrage sequence exists")
}

func main() {
	in, _ := os.Open("104.in")
	defer in.Close()
	out, _ = os.Create("104.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		matrix := make([][][]node, n)
		for i := range matrix {
			matrix[i] = make([][]node, n)
			for j := range matrix[i] {
				matrix[i][j] = make([]node, n)
				if i == j {
					matrix[i][j][0] = node{1, []int{}}
				} else {
					fmt.Fscanf(in, "%f", &matrix[i][j][0].v)
				}
			}
		}
		floydWarshall(n, matrix)
	}
}
