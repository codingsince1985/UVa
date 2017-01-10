// UVa 784 - Maze Exploration

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	out        *os.File
	directions = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
)

func dfs(x, y int, maze [][]byte) {
	maze[x][y] = '#'
	for _, direction := range directions {
		newX, newY := x+direction[0], y+direction[1]
		if newX >= 0 && newY >= 0 && newX < len(maze[newY]) && newY < len(maze) && maze[newX][newY] == ' ' {
			dfs(newX, newY, maze)
		}
	}
}

func output(maze [][]byte, line string) {
	for _, vi := range maze {
		fmt.Fprintln(out, string(vi))
	}
	fmt.Fprintln(out, line)
}

func main() {
	in, _ := os.Open("784.in")
	defer in.Close()
	out, _ = os.Create("784.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var x, y int
	var maze [][]byte
	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	for kase > 0 && s.Scan() {
		if line := s.Text(); strings.HasPrefix(line, "_") {
			dfs(x, y, maze)
			output(maze, line)
			maze = nil
			kase--
		} else {
			maze = append(maze, []byte(line))
			if x <= 0 {
				x = strings.Index(line, "*")
				if x <= 0 {
					y++
				}
			}
		}
	}
}
