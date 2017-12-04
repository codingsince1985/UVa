// UVa 10959 - The Party, Part I

package main

import (
	"fmt"
	"math"
	"os"
)

var distances []int

func bfs(matrix [][]bool, giovanni int) {
	distances[giovanni] = 0
	for visited, queue := map[int]bool{giovanni: true}, []int{giovanni}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		for i, cell := range matrix[curr] {
			if cell && !visited[i] {
				visited[i] = true
				if distances[i] > distances[curr]+1 {
					distances[i] = distances[curr] + 1
				}
				queue = append(queue, i)
			}
		}
	}
}

func solve(p int, couples [][2]int) {
	matrix := make([][]bool, p)
	for i := range matrix {
		matrix[i] = make([]bool, p)
	}
	for _, couple := range couples {
		matrix[couple[0]][couple[1]], matrix[couple[1]][couple[0]] = true, true
	}
	distances = make([]int, p)
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	bfs(matrix, 0)
}

func main() {
	in, _ := os.Open("10959.in")
	defer in.Close()
	out, _ := os.Create("10959.out")
	defer out.Close()

	var kase, p, d int
	first := true
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d%d", &p, &d)
		couples := make([][2]int, d)
		for i := range couples {
			fmt.Fscanf(in, "%d%d", &couples[i][0], &couples[i][1])
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		solve(p, couples)
		for i := 1; i < p; i++ {
			fmt.Fprintln(out, distances[i])
		}
	}
}
