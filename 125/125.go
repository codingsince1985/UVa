// UVa 125 - Numbering Paths

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var m int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func floydWarshall(matrix [][]int) [][]int {
	for k := 0; k <= m; k++ {
		for i := 0; i <= m; i++ {
			for j := 0; j <= m; j++ {
				if matrix[i][k] > 0 && matrix[k][j] > 0 {
					matrix[i][j] += matrix[i][k] * matrix[k][j]
				}
			}
		}
	}
	return matrix
}

func postProcess(matrix [][]int) [][]int {
	for k := 0; k <= m; k++ {
		if matrix[k][k] > 0 {
			for i := 0; i <= m; i++ {
				for j := 0; j <= m; j++ {
					if i == k && j == k {
						continue
					}
					if matrix[i][k] > 0 && matrix[k][j] > 0 {
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
	return convert(postProcess(floydWarshall(matrix)))
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
