// UVa 10440 - Ferry Loading II

package main

import (
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(n, t int, c []int) (int, int) {
	time, trip := -2 * t, 0
	s := len(c)
	mod := s % n
	if mod != 0 {
		time = c[mod - 1]
		trip ++
	}
	for i := mod + n - 1; i < s; i += n {
		time = max(time + 2 * t, c[i])
		trip ++
	}
	return time + t, trip
}

func main() {
	in, _ := os.Open("10440.in")
	defer in.Close()
	out, _ := os.Create("10440.out")
	defer out.Close()

	var c, n, t, m int
	fmt.Fscanf(in, "%d", &c)
	for i := 0; i < c; i++ {
		fmt.Fscanf(in, "%d%d%d", &n, &t, &m)
		c := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(in, "%d", &c[j])
		}
		time, trip := solve(n, t, c)
		fmt.Fprintln(out, time, trip)
	}
}