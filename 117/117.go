// UVa 117 - The Postal Worker Rings Once

package main

import (
	"fmt"
	"math"
	"os"
)

type edge struct {
	n1, n2 string
	l      int
}

func toMatrix(edges []edge) ([]int, int, [][]int) {
	ns := make(map[string]int)
	nl := make(map[string]int)
	idx, sum := 0, 0
	for _, v := range edges {
		sum += v.l
		ns[v.n1]++
		ns[v.n2]++
		if _, ok := nl[v.n1]; !ok {
			nl[v.n1] = idx
			idx++
		}
		if _, ok := nl[v.n2]; !ok {
			nl[v.n2] = idx
			idx++
		}
	}

	matrix := make([][]int, len(nl))
	for i := range matrix {
		matrix[i] = make([]int, len(nl))
		for j := range matrix[i] {
			matrix[i][j] = math.MaxInt32
		}
	}
	for _, v := range edges {
		matrix[nl[v.n1]][nl[v.n2]], matrix[nl[v.n2]][nl[v.n1]] = v.l, v.l
	}

	var odds []int
	for k, v := range ns {
		if v%2 != 0 {
			odds = append(odds, nl[k])
		}
	}
	return odds, sum, matrix
}

func floydWarshall(matrix [][]int) [][]int {
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				matrix[i][j] = min(matrix[i][j], matrix[i][k]+matrix[k][j])
			}
		}
	}
	return matrix
}

func main() {
	in, _ := os.Open("117.in")
	defer in.Close()
	out, _ := os.Create("117.out")
	defer out.Close()

	var word string
	var edges []edge
	for {
		if _, err := fmt.Fscanf(in, "%s", &word); err != nil {
			break
		}
		if word != "deadend" {
			edges = append(edges, edge{word[:1], word[len(word)-1:], len(word)})
		} else {
			if odds, sum, matrix := toMatrix(edges); len(odds) == 0 {
				fmt.Fprintln(out, sum)
			} else {
				fmt.Fprintln(out, sum+floydWarshall(matrix)[odds[0]][odds[1]])
			}
			edges = nil
		}
	}
}
