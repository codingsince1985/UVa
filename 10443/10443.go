// UVa 10443 - Rock

package main

import (
	"fmt"
	"os"
)

var (
	delta = []int{-1, 0, 1}
	out   *os.File
)

func solve(grid [][]byte, r, c, n int) {
	tmp := make([][]byte, r)
	for k := 0; k < n; k++ {
		for x := range grid {
			tmp[x] = make([]byte, c)
			for y := range grid[x] {
				tmp[x][y] = grid[x][y]
				for _, dx := range delta {
					for _, dy := range delta {
						i, j := x+dx, y+dy
						if !(dx != 0 && dy != 0 || dx == 0 && dy == 0 || i < 0 || i >= r || j < 0 || j >= r) {
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

func output(grid [][]byte) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Fprintf(out, "%c", grid[i][j])
		}
		fmt.Fprintln(out)
	}
}

func main() {
	in, _ := os.Open("10443.in")
	defer in.Close()
	out, _ = os.Create("10443.out")
	defer out.Close()

	var kase, r, c, n int
	var grid [][]byte
	var line string
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanf(in, "%d%d%d", &r, &c, &n)
		grid = make([][]byte, r)
		for j := range grid {
			grid[j] = make([]byte, c)
			fmt.Fscanf(in, "%s", &line)
			for k := range line {
				grid[j][k] = line[k]
			}
		}
		solve(grid, r, c, n)
		output(grid)
		kase--
		if kase > 0 {
			fmt.Fprintln(out)
		}
	}
}
