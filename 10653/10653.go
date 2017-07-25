// UVa 10653 - Bombs! NO they are Mines!!

package main

import (
	"fmt"
	"os"
)

type (
	cell struct{ y, x int }
	node struct {
		cell
		l int
	}
)

var (
	r, c       int
	grid       [][]bool
	directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

func bfs(s, d cell) int {
	visited := map[[2]int]bool{[2]int{s.y, s.x}: true}
	queue := []node{{s, 0}}
	for len(queue) > 0 {
		curr := queue[0]
		if curr.cell == d {
			return curr.l
		}
		queue = queue[1:]
		for _, direction := range directions {
			if ny, nx := curr.y+direction[0], curr.x+direction[1]; ny >= 0 && ny < r && nx >= 0 && nx < c && !grid[ny][nx] && !visited[[2]int{ny, nx}] {
				visited[[2]int{ny, nx}] = true
				queue = append(queue, node{cell{ny, nx}, curr.l + 1})
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("10653.in")
	defer in.Close()
	out, _ := os.Create("10653.out")
	defer out.Close()

	var rows, row, n, column int
	var s, d cell
	for {
		if fmt.Fscanf(in, "%d%d", &r, &c); r == 0 && c == 0 {
			break
		}
		grid = make([][]bool, r)
		for i := range grid {
			grid[i] = make([]bool, c)
		}
		for fmt.Fscanf(in, "%d", &rows); rows > 0; rows-- {
			fmt.Fscanf(in, "%d", &row)
			for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
				fmt.Fscanf(in, "%d", &column)
				grid[row][column] = true
			}
		}
		fmt.Fscanf(in, "%d%d", &s.y, &s.x)
		fmt.Fscanf(in, "%d%d", &d.y, &d.x)
		fmt.Fprintln(out, bfs(s, d))
	}
}
