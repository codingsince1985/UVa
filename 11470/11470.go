// UVa 11470 - Square Sums

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(square [][]int, n, i int) int {
	var sum int
	for j := i; j < n-i; j++ {
		sum += square[i][j]
		if i != n-i-1 {
			sum += square[n-i-1][j]
		}
	}
	for j := i + 1; j < n-i-1; j++ {
		sum += square[j][i] + square[j][n-i-1]
	}
	return sum
}

func solve(n int, square [][]int) []string {
	var sums []string
	for i := 0; i < (n+1)/2; i++ {
		sums = append(sums, strconv.Itoa(add(square, n, i)))
	}
	return sums
}

func main() {
	in, _ := os.Open("11470.in")
	defer in.Close()
	out, _ := os.Create("11470.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		square := make([][]int, n)
		for i := range square {
			square[i] = make([]int, n)
			for j := range square[i] {
				fmt.Fscanf(in, "%d", &square[i][j])
			}
		}
		fmt.Fprintf(out, "Case %d: %s\n", kase, strings.Join(solve(n, square), " "))
	}
}
