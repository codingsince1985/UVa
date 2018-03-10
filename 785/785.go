// UVa 785 - Grid Colouring

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	grid       [][]byte
	directions = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
)

func dfs(x, y int, cell byte) {
	for _, direction := range directions {
		if newY, newX := y+direction[1], x+direction[0]; grid[newY][newX] == ' ' {
			grid[newY][newX] = cell
			dfs(newX, newY, cell)
		}
	}
}

func solve(out io.Writer) {
	for y, row := range grid {
		for x, cell := range row {
			if cell != 'X' && cell != ' ' {
				dfs(x, y, cell)
			}
		}
	}
	for _, row := range grid {
		fmt.Fprintln(out, string(row))
	}
}

func main() {
	in, _ := os.Open("785.in")
	defer in.Close()
	out, _ := os.Create("785.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		if line := s.Text(); line == "_____________________________" {
			solve(out)
			fmt.Fprintln(out, line)
			grid = nil
		} else {
			grid = append(grid, []byte(line))
		}
	}
}
