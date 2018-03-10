// UVa 796 - Critical Links

package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

var (
	matrix, cut [][]bool
	visited     []bool
	depth, low  []int
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(current, d, parent int) int {
	m := math.MaxInt32
	depth[current] = d
	for i := range matrix {
		if matrix[i][current] {
			if !visited[i] {
				visited[i] = true
				l := dfs(i, d+1, current)
				if l > d {
					cut[i][current], cut[current][i] = true, true
				}
				m = min(m, l)
			} else {
				if i != parent {
					m = min(m, depth[i])
				}
			}
		}
	}
	low[current] = m
	return low[current]
}

func solve(out io.Writer, n int) {
	low = make([]int, n)
	depth = make([]int, n)
	cut = make([][]bool, n)
	for i := range cut {
		cut[i] = make([]bool, n)
	}
	visited = make([]bool, n)
	for i := range matrix {
		if !visited[i] {
			visited[i] = true
			dfs(i, 1, -1)
		}
	}
	var count int
	var output strings.Builder
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if cut[i][j] {
				count++
				output.WriteString(fmt.Sprintf("%d - %d\n", i, j))
			}
		}
	}
	fmt.Fprintf(out, "%d critical links\n%s\n", count, output.String())
}

func main() {
	in, _ := os.Open("796.in")
	defer in.Close()
	out, _ := os.Create("796.out")
	defer out.Close()

	var n, m, s, t int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		matrix = make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		for i := 0; i < n; i++ {
			for fmt.Fscanf(in, "%d (%d)", &s, &m); m > 0; m-- {
				fmt.Fscanf(in, "%d", &t)
				matrix[s][t], matrix[t][s] = true, true
			}
		}
		solve(out, n)
		fmt.Fscanln(in)
		fmt.Fscanln(in)
	}
}
