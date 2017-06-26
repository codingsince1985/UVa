// UVa 10596 - Morning Walk

package main

import (
	"fmt"
	"os"
)

var (
	n       int
	visited []bool
	matrix  [][]bool
)

func dfs(curr int) {
	for i := 0; i < n; i++ {
		if matrix[curr][i] && !visited[i] {
			visited[i] = true
			dfs(i)
		}
	}
}

func solve(roads []int) bool {
	for _, road := range roads {
		if road%2 == 1 {
			return false
		}
	}
	visited = make([]bool, n)
	visited[0] = true
	dfs(0)
	return len(visited) == n
}

func main() {
	in, _ := os.Open("10596.in")
	defer in.Close()
	out, _ := os.Create("10596.out")
	defer out.Close()

	var r, c1, c2 int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &r); err != nil {
			break
		}
		matrix = make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		roads := make([]int, n)
		for ; r > 0; r-- {
			fmt.Fscanf(in, "%d%d", &c1, &c2)
			matrix[c1][c2] = true
			roads[c1]++
			roads[c2]++
		}
		if solve(roads) {
			fmt.Fprintln(out, "Possible")
		} else {
			fmt.Fprintln(out, "Not Possible")
		}
	}
}
