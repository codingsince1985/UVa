// UVa 639 - Don't Get Rooked

package main

import (
	"fmt"
	"os"
)

var (
	n, max int
	grid   [][]byte
)

func ok(x, y int) bool {
	for i := x - 1; i >= 0 && grid[i][y] != 'X'; i-- {
		if grid[i][y] == '0' {
			return false
		}
	}
	for i := x + 1; i < n && grid[i][y] != 'X'; i++ {
		if grid[i][y] == '0' {
			return false
		}
	}
	for i := y - 1; i >= 0 && grid[x][i] != 'X'; i-- {
		if grid[x][i] == '0' {
			return false
		}
	}
	for i := y + 1; i < n && grid[x][i] != 'X'; i++ {
		if grid[x][i] == '0' {
			return false
		}
	}
	return true
}

func dfs(x, y, num int) {
	if num > max {
		max = num
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '.' && ok(i, j) {
				grid[i][j] = '0'
				dfs(i, j, num+1)
				grid[i][j] = '.'
			}
		}
	}
}

func main() {
	in, _ := os.Open("639.in")
	defer in.Close()
	out, _ := os.Create("639.out")
	defer out.Close()

	var c byte
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		grid = make([][]byte, n)
		for i := range grid {
			grid[i] = make([]byte, n)
			for j := range grid[i] {
				fmt.Fscanf(in, "%c", &c)
				grid[i][j] = c
			}
			fmt.Fscanf(in, "\n")
		}
		max = 0
		dfs(0, 0, 0)
		fmt.Fprintln(out, max)
	}
}
