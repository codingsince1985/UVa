// UVa 280 - Vertex

package main

import (
	"fmt"
	"os"
)

var (
	matrix  [][]bool
	visited map[int]bool
)

func dfs(node int) {
	for i, vi := range matrix[node] {
		if vi && !visited[i] {
			visited[i] = true
			dfs(i)
		}
	}
}

func output(out *os.File, n int, visited map[int]bool) {
	fmt.Fprint(out, n-len(visited))
	for i := 0; i < n; i++ {
		if !visited[i] {
			fmt.Fprintf(out, " %d", i+1)
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("280.in")
	defer in.Close()
	out, _ := os.Create("280.out")
	defer out.Close()

	var n, s, e, q, qs int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		matrix = make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		for {
			if fmt.Fscanf(in, "%d", &s); s == 0 {
				break
			}
			for {
				if fmt.Fscanf(in, "%d", &e); e == 0 {
					break
				}
				matrix[s-1][e-1] = true
			}
		}
		for fmt.Fscanf(in, "%d", &q); q > 0; q-- {
			fmt.Fscanf(in, "%d", &qs)
			visited = make(map[int]bool)
			dfs(qs - 1)
			output(out, n, visited)
		}
	}
}
