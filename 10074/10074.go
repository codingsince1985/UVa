// UVa 10074 - Take the Land

package main

import (
	"fmt"
	"os"
)

type cell struct{ tree, up int }

func verticalSum(land [][]cell) {
	for i := range land {
		for j := range land[i] {
			if land[i][j].tree == 1 {
				land[i][j].up = 0
			} else {
				if i == 0 {
					land[i][j].up = 1
				} else {
					land[i][j].up = land[i-1][j].up + 1
				}
			}
		}
	}
}

func solve(n int, land [][]cell) int {
	verticalSum(land)
	var maxArea int
	for i := range land {
		for j := range land[i] {
			sum := land[i][j].up
			for k := j + 1; k < n && land[i][k].up >= land[i][j].up; k++ {
				sum += land[i][j].up
			}
			for k := j - 1; k >= 0 && land[i][k].up >= land[i][j].up; k-- {
				sum += land[i][j].up
			}
			if sum > maxArea {
				maxArea = sum
			}
		}
	}
	return maxArea
}

func main() {
	in, _ := os.Open("10074.in")
	defer in.Close()
	out, _ := os.Create("10074.out")
	defer out.Close()

	var m, n int
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 && n == 0 {
			break
		}
		land := make([][]cell, m)
		for i := range land {
			land[i] = make([]cell, n)
			for j := range land[i] {
				fmt.Fscanf(in, "%d", &land[i][j].tree)
			}
		}
		fmt.Fprintln(out, solve(n, land))
	}
}
