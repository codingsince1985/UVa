// UVa 469 - Wetlands of Florida

package main

import (
	"fmt"
	"os"
	"strconv"
)

var count int

func dfs(grid [][]byte, visited [][]bool, x, y int) {
	if x < 1 || y < 1 || x > len(grid) || y > len((grid)[0]) {
		return
	}

	if grid[x-1][y-1] == 'W' && !visited[x-1][y-1] {
		visited[x-1][y-1] = true
		count++

		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if !(dx == 0 && dy == 0) {
					dfs(grid, visited, x+dx, y+dy)
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

	var n, x, y int
	var line string
	var grid [][]byte
	var visited [][]bool
	fmt.Fscanf(in, "%d", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "\n")
		grid = nil
		for {
			if fmt.Fscanf(in, "%s", &line); line[0] != 'L' && line[0] != 'W' {
				break
			}
			grid = append(grid, []byte(line))
		}
		x, _ = strconv.Atoi(line)
		fmt.Fscanf(in, "%d", &y)
		for {
			visited = make([][]bool, len(grid))
			for j := range visited {
				visited[j] = make([]bool, len(grid[0]))
			}
			count = 0
			dfs(grid, visited, x, y)
			fmt.Fprintln(out, count)
			if _, err := fmt.Fscanf(in, "%d%d", &x, &y); err != nil {
				break
			}
		}
		fmt.Fprintln(out)
	}
}
