// UVa 108 - Maximum Sum

package main

import (
	"fmt"
	"os"
	"math"
)

func findMax(s [][]int) int {
	l := len(s)
	max := math.MinInt32
	for i := 0; i < l - 1; i++ {
		for j := 0; j < l - 1; j++ {
			for k := i; k < l; k++ {
				for m := j; m < l; m++ {
					tmp := s[k][m]
					if i != 0 && j != 0 {
						tmp += s[i - 1][j - 1]
					}
					if i != 0 {
						tmp -= s[i - 1][m]
					}
					if j != 0 {
						tmp -= s[k][j - 1]
					}
					if tmp > max {
						max = tmp
					}
				}
			}
		}
	}
	return max
}

func prefixSum(r [][]int) [][]int {
	l := len(r)
	s := make([][]int, l)
	for i := 0; i < l; i++ {
		s[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		ss := 0
		for j := 0; j < l; j++ {
			ss += r[i][j]
			s[i][j] = ss
			if i > 0 {
				s[i][j] += s[i - 1][j]
			}
		}
	}
	return s
}

func main() {
	in, _ := os.Open("108.in")
	defer in.Close()
	out, _ := os.Create("108.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		r := make([][]int, n)
		for i := 0; i < n; i++ {
			r[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Fscanf(in, "%d", &r[i][j])
			}
		}
		fmt.Fprintln(out, findMax(prefixSum(r)))
	}
}
