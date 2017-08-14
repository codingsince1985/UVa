// UVa 10279 - Mine Sweeper

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	n          int
	directions = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
)

func check(y, x int, grid [][]byte) int {
	var count int
	for _, direction := range directions {
		if newY, newX := y+direction[0], x+direction[1]; newY >= 0 && newY < n && newX >= 0 && newX < n && grid[newY][newX] == '*' {
			count++
		}
	}
	return count
}

func mark(grid, result [][]byte) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '*' {
				result[y][x] = '*'
			}
		}
	}
}

func solve(grid, touched [][]byte) [][]byte {
	result := make([][]byte, n)
	for i := range result {
		result[i] = []byte(strings.Repeat(".", n))
	}
	lose := false
	for y, row := range touched {
		for x, cell := range row {
			if cell == 'x' {
				if grid[y][x] == '*' {
					lose = true
				} else {
					result[y][x] = byte('0' + check(y, x, grid))
				}
			}
		}
	}
	if lose {
		mark(grid, result)
	}
	return result
}

func main() {
	in, _ := os.Open("10279.in")
	defer in.Close()
	out, _ := os.Create("10279.out")
	defer out.Close()

	var kase int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &n)
		grid := make([][]byte, n)
		for i := range grid {
			fmt.Fscanf(in, "%s", &line)
			grid[i] = []byte(line)
		}
		touched := make([][]byte, n)
		for i := range touched {
			fmt.Fscanf(in, "%s", &line)
			touched[i] = []byte(line)
		}
		result := solve(grid, touched)
		for _, row := range result {
			fmt.Fprintln(out, string(row))
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
