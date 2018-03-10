// UVa 784 - Maze Exploration

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var directions = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func dfs(x, y int, maze [][]byte) {
	maze[x][y] = '#'
	for _, direction := range directions {
		newX, newY := x+direction[0], y+direction[1]
		if newX >= 0 && newY >= 0 && newX < len(maze[newY]) && newY < len(maze) && maze[newX][newY] == ' ' {
			dfs(newX, newY, maze)
		}
	}
}

func output(out io.Writer, maze [][]byte, line string) {
	for _, vi := range maze {
		fmt.Fprintln(out, string(vi))
	}
	fmt.Fprintln(out, line)
}

func find(maze [][]byte) (int, int) {
	for y, row := range maze {
		for x, cell := range row {
			if cell == '*' {
				return x, y
			}
		}
	}
	return -1, -1
}

func main() {
	in, _ := os.Open("784.in")
	defer in.Close()
	out, _ := os.Create("784.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase int
	var maze [][]byte
	s.Scan()
	for fmt.Sscanf(s.Text(), "%d", &kase); kase > 0 && s.Scan(); {
		if line := s.Text(); strings.HasPrefix(line, "_") {
			x, y := find(maze)
			dfs(x, y, maze)
			output(out, maze, line)
			maze = nil
			kase--
		} else {
			maze = append(maze, []byte(line))
		}
	}
}
