// UVa 11110 - Equidivisions

package main

import (
	"fmt"
	"os"
)

var (
	n          int
	cells      [][]int
	visited    [][]bool
	directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

type node struct{ y, x int }

func bfs(y, x, v int) int {
	count := 1
	visited[y][x] = true
	for queue := []node{{y, x}}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		for _, direction := range directions {
			if newY, newX := curr.y+direction[0], curr.x+direction[1]; newY >= 0 && newY < n && newX >= 0 && newX < n && cells[newY][newX] == v && !visited[newY][newX] {
				visited[newY][newX] = true
				count++
				queue = append(queue, node{newY, newX})
			}
		}
	}
	return count
}

func solve() bool {
	visited = make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
here:
	for i := 1; i <= n; i++ {
		for y, row := range cells {
			for x, cell := range row {
				if cell == i {
					if bfs(y, x, i) != n {
						return false
					}
					continue here
				}
			}
		}
	}
	return true
}

func main() {
	in, _ := os.Open("11110.in")
	defer in.Close()
	out, _ := os.Create("11110.out")
	defer out.Close()

	var y, x int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		cells = make([][]int, n)
		for i := range cells {
			cells[i] = make([]int, n)
			for j := range cells[i] {
				cells[i][j] = n
			}
		}
		for i := 1; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Fscanf(in, "%d%d", &y, &x)
				cells[y-1][x-1] = i
			}
		}
		if solve() {
			fmt.Fprintln(out, "good")
		} else {
			fmt.Fprintln(out, "wrong")
		}
	}
}
