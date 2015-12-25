// UVa 108 - Maximum Sum

package main

import (
	"fmt"
	"os"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMax(s [][]int) int {
	l := len(s)
	msf := math.MinInt32
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			meh := 0
			for k := 0; k < l; k++ {
				tmp := s[j][k]
				if i > 0 {
					tmp -= s[i - 1][k]
				}
				meh = max(meh + tmp, tmp)
				msf = max(msf, meh)
			}
		}
	}
	return msf
}

func prefixSum(r [][]int) [][]int {
	l := len(r)
	s := make([][]int, l)
	for i := 0; i < l; i++ {
		s[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			s[i][j] = r[i][j]
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
