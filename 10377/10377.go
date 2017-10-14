// UVa 10377 - Maze Traversal

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	row, column  int
	directions   = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	directionMap = map[int]byte{0: 'N', 1: 'E', 2: 'S', 3: 'W'}
)

func solve(r, c int, maze [][]byte, command string) (int, int, byte) {
	steps := []byte(strings.Replace(command, " ", "", -1))
	var idx int
here:
	for _, step := range steps {
		switch step {
		case 'R':
			if idx++; idx == 4 {
				idx = 0
			}
		case 'L':
			if idx--; idx == -1 {
				idx = 3
			}
		case 'F':
			direction := directions[idx]
			if newR, newC := r+direction[1], c+direction[0]; newR >= 0 && newR < row && newC >= 0 && newC < column && maze[newR][newC] == ' ' {
				r, c = newR, newC
			}
		case 'Q':
			break here
		}
	}
	return r + 1, c + 1, directionMap[idx]
}

func main() {
	in, _ := os.Open("10377.in")
	defer in.Close()
	out, _ := os.Create("10377.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, r, c int
	var maze [][]byte
	s.Scan()
	for fmt.Sscanf(s.Text(), "%d", &kase); kase > 0 && s.Scan(); kase-- {
		s.Scan()
		fmt.Sscanf(s.Text(), "%d%d", &row, &column)
		maze = make([][]byte, row)
		for i := range maze {
			s.Scan()
			maze[i] = []byte(s.Text())
		}
		s.Scan()
		fmt.Sscanf(s.Text(), "%d%d", &r, &c)
		var command string
		for s.Scan() {
			if command += s.Text(); strings.HasSuffix(command, "Q") {
				break
			}
		}
		y, x, d := solve(r-1, c-1, maze, command)
		fmt.Fprintf(out, "%d %d %c\n", y, x, d)
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
