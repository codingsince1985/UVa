// UVa 10443 - Rock

package main

import (
	"fmt"
	"os"
)

var delta = []int{-1, 0, 1}

func solve(grid [][]byte, r, c, n int) {
	tmp := make([][]byte, r)
	for k := 0; k < n; k++ {
		for x := range grid {
			tmp[x] = make([]byte, c)
			for y := range grid[x] {
				tmp[x][y] = grid[x][y]
				for _, dx := range delta {
					for _, dy := range delta {
						if i, j := x+dx, y+dy; !(dx != 0 && dy != 0 || dx == 0 && dy == 0 || i < 0 || i >= r || j < 0 || j >= r) {
							switch {
							case grid[x][y] == 'R' && grid[i][j] == 'P':
								tmp[x][y] = 'P'
							case grid[x][y] == 'P' && grid[i][j] == 'S':
								tmp[x][y] = 'S'
							case grid[x][y] == 'S' && grid[i][j] == 'R':
								tmp[x][y] = 'R'
							}
						}
					}
				}
			}
		}
		copy(grid, tmp)
	}
}

func main() {
	in, _ := os.Open("10443.in")
	defer in.Close()
	out, _ := os.Create("10443.out")
	defer out.Close()

	var kase, r, c, n int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d%d", &r, &c, &n)
		grid := make([][]byte, r)
		for j := range grid {
			grid[j] = make([]byte, c)
			fmt.Fscanf(in, "%s", &line)
			for k := range line {
				grid[j][k] = line[k]
			}
		}
		solve(grid, r, c, n)
		for _, row := range grid {
			fmt.Fprintln(out, string(row))
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
