// UVa 703 - Triple Ties: The Organizer's Nightmare

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(out io.Writer, matrix [][]bool) {
	var ties [][]string
	var count int
	for i := range matrix {
		for j := range matrix {
			for k := range matrix {
				if (i > j && j > k || i < j && j < k) && !matrix[i][j] && !matrix[j][k] && !matrix[k][i] ||
					i < j && j < k && matrix[i][j] && matrix[j][k] && matrix[i][k] && matrix[j][i] && matrix[k][i] && matrix[k][j] {
					count++
					ties = append(ties, []string{strconv.Itoa(i + 1), strconv.Itoa(j + 1), strconv.Itoa(k + 1)})
				}
			}
		}
	}
	fmt.Fprintln(out, len(ties))
	for _, tie := range ties {
		fmt.Fprintln(out, strings.Join(tie, " "))
	}
}

func main() {
	in, _ := os.Open("703.in")
	defer in.Close()
	out, _ := os.Create("703.out")
	defer out.Close()

	var n, m int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		matrix := make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
			for j := range matrix[i] {
				fmt.Fscanf(in, "%d", &m)
				matrix[i][j] = m == 0
			}
		}
		solve(out, matrix)
	}
}
