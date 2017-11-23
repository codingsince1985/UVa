// UVa 469 - Wetlands of Florida

package main

import (
	"fmt"
	"os"
)

var (
	count   int
	delta   = []int{-1, 0, 1}
	grid    [][]byte
	visited [][]bool
)

func dfs(x, y int) {
	if x < 1 || y < 1 || x > len(grid) || y > len((grid)[0]) {
		return
	}

	if grid[x-1][y-1] == 'W' && !visited[x-1][y-1] {
		visited[x-1][y-1] = true
		count++

		for _, dx := range delta {
			for _, dy := range delta {
				if !(dx == 0 && dy == 0) {
					dfs(x+dx, y+dy)
				}
			}
		}
	}
}

func main() {
	in, _ := os.Open("469.in")
	defer in.Close()
	out, _ := os.Create("469.out")
	defer out.Close()

	var kase, x, y int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		grid = nil
		for {
			if fmt.Fscanf(in, "%s", &line); line[0] != 'L' && line[0] != 'W' {
				break
			}
			grid = append(grid, []byte(line))
		}
		fmt.Sscanf(line, "%d", &x)
		fmt.Fscanf(in, "%d", &y)
		for {
			visited = make([][]bool, len(grid))
			for j := range visited {
				visited[j] = make([]bool, len(grid[0]))
			}
			count = 0
			dfs(x, y)
			fmt.Fprintln(out, count)
			if _, err := fmt.Fscanf(in, "%d%d", &x, &y); err != nil {
				break
			}
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
