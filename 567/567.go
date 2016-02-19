// UVa 567 - Risk

package main

import (
	"fmt"
	"os"
)

const N = 20

type node struct {
	n, step int
}

var matrix [][]bool

func bfs(n1, n2 int) int {
	var queue []node
	visited := make([]bool, N+1)
	visited[n1] = true
	queue = append(queue, node{n1, 0})

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for i := 1; i <= N; i++ {
			if !visited[i] && matrix[curr.n][i] {
				if i == n2 {
					return curr.step + 1
				}
				queue = append(queue, node{i, curr.step + 1})
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("567.in")
	defer in.Close()
	out, _ := os.Create("567.out")
	defer out.Close()

	var cnt, tmp, set, kase, n1, n2 int
	for {
		set++
		matrix = make([][]bool, N+1)
		for i := range matrix {
			matrix[i] = make([]bool, N+1)
		}
		for i := 1; i < N; i++ {
			if _, err := fmt.Fscanf(in, "%d", &cnt); err != nil {
				return
			}
			for j := 0; j < cnt; j++ {
				fmt.Fscanf(in, "%d", &tmp)
				matrix[i][tmp] = true
				matrix[tmp][i] = true
			}
		}
		fmt.Fprintf(out, "Test Set #%d\n", set)
		fmt.Fscanf(in, "%d", &kase)
		for i := 0; i < kase; i++ {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			fmt.Fprintf(out, "%2d to %2d: %d\n", n1, n2, bfs(n1, n2))
		}
		fmt.Fprintln(out)
	}
}
