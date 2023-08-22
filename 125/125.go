// UVa 125 - Numbering Paths

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var m int

func floydWarshall(matrix [][]int) {
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				if matrix[i][k] > 0 && matrix[k][j] > 0 {
					matrix[i][j] += matrix[i][k] * matrix[k][j]
				}
			}
		}
	}
}

func postProcess(matrix [][]int) [][]int {
	for k := range matrix {
		if matrix[k][k] > 0 {
			for i := range matrix {
				for j := range matrix {
					if (k != i || k != j) && matrix[i][k] > 0 && matrix[k][j] > 0 {
						matrix[i][j] = -1
					}
				}
			}
			matrix[k][k] = -1
		}
	}
	return matrix
}

func convert(matrix [][]int) [][]string {
	str := make([][]string, m+1)
	for i := range str {
		str[i] = make([]string, m+1)
		for j, vj := range matrix[i] {
			str[i][j] = strconv.Itoa(vj)
		}
	}
	return str
}

func solve(street [][2]int) [][]string {
	matrix := make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, m+1)
	}
	for _, vi := range street {
		matrix[vi[0]][vi[1]] = 1
	}
	floydWarshall(matrix)
	return convert(postProcess(matrix))
}

func main() {
	in, _ := os.Open("125.in")
	defer in.Close()
	out, _ := os.Create("125.out")
	defer out.Close()

	var n, n1, n2 int
	for kase := 0; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		var street [][2]int
		m = 0
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			street = append(street, [2]int{n1, n2})
			m = max(m, max(n1, n2))
		}
		fmt.Fprintf(out, "matrix for city %d\n", kase)
		for _, vi := range solve(street) {
			fmt.Fprintln(out, strings.Join(vi, " "))
		}
	}
}
