// UVa 291 - The House Of Santa Claus

package main

import (
	"fmt"
	"os"
)

var (
	out    *os.File
	matrix [][]bool
)

func buildMatrix() [][]bool {
	adjacency := [][]int{{}, {2, 3, 5}, {3, 5}, {4, 5}, {5}, {}}
	matrix := make([][]bool, len(adjacency))
	for i := range adjacency {
		matrix[i] = make([]bool, len(adjacency))
	}
	for i := range adjacency {
		for _, v := range adjacency[i] {
			matrix[i][v], matrix[v][i] = true, true
		}
	}
	return matrix
}

func backtracking(steps []int) {
	if len(steps) == 9 {
		for _, v := range steps {
			fmt.Fprint(out, v)
		}
		fmt.Fprintln(out)
		return
	}
	curr := steps[len(steps)-1]
	for i := 1; i <= 5; i++ {
		if matrix[curr][i] {
			matrix[curr][i], matrix[i][curr] = false, false
			backtracking(append(steps, i))
			matrix[curr][i], matrix[i][curr] = true, true
		}
	}
}

func main() {
	out, _ = os.Create("291.out")
	defer out.Close()

	matrix = buildMatrix()
	var steps []int
	backtracking(append(steps, 1))
}
