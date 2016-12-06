// UVa 10502 - Counting Rectangles

package main

import (
	"fmt"
	"os"
)

var n, m int

func prefixSum(board [][]byte) [][]int {
	sum := make([][]int, n)
	for i := range sum {
		sum[i] = make([]int, m)
		for j := range sum[i] {
			var total int
			if i > 0 {
				total += sum[i-1][j]
			}
			if j > 0 {
				total += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				total -= sum[i-1][j-1]
			}
			sum[i][j] = total + int(board[i][j]-'0')
		}
	}
	return sum
}

func solve(sum [][]int) int {
	var count int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := i; k < n; k++ {
				for l := j; l < m; l++ {
					total := sum[k][l]
					if i > 0 {
						total -= sum[i-1][l]
					}
					if j > 0 {
						total -= sum[k][j-1]
					}
					if i > 0 && j > 0 {
						total += sum[i-1][j-1]
					}
					if total == (k-i+1)*(l-j+1) {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	in, _ := os.Open("10502.in")
	defer in.Close()
	out, _ := os.Create("10502.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fscanf(in, "%d", &m)
		board := make([][]byte, n)
		for i := range board {
			fmt.Fscanf(in, "%s", &line)
			board[i] = []byte(line)
		}
		fmt.Fprintln(out, solve(prefixSum(board)))
	}
}
