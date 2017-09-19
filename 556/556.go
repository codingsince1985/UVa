// UVa 556 - Amazing

package main

import (
	"fmt"
	"os"
)

type cell struct {
	open  bool
	count int
}

var (
	maze       [][]cell
	directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func openOnRight(y, x, direction int) bool {
	direction = (direction + 1) % 4
	return maze[y+directions[direction][0]][x+directions[direction][1]].open
}

func openInFront(y, x, direction int) bool { return openOnRight(y, x, (direction+3)%4) }

func move(y, x, direction int) (bool, int, int) {
	y, x = y+directions[direction][0], x+directions[direction][1]
	maze[y][x].count++
	return y == len(maze)-2 && x == 1, y, x
}

func stat() map[int]int {
	s := make(map[int]int)
	for _, row := range maze {
		for _, cell := range row {
			if cell.open {
				s[cell.count]++
			}
		}
	}
	return s
}

func solve() map[int]int {
	var ok bool
	y, x, direction := len(maze)-2, 1, 0
here:
	for {
		for openOnRight(y, x, direction) {
			direction = (direction + 1) % 4
			if ok, y, x = move(y, x, direction); ok {
				break here
			}
		}
		if openInFront(y, x, direction) {
			if ok, y, x = move(y, x, direction); ok {
				break
			}
		}
		if !openOnRight(y, x, direction) && !openInFront(y, x, direction) {
			direction = (direction + 3) % 4
		}
	}
	return stat()
}

func main() {
	in, _ := os.Open("556.in")
	defer in.Close()
	out, _ := os.Create("556.out")
	defer out.Close()

	var b, w int
	var line string
	for {
		if fmt.Fscanf(in, "%d%d", &b, &w); b == 0 && w == 0 {
			break
		}
		maze = make([][]cell, b+2)
		for i := range maze {
			maze[i] = make([]cell, w+2)
		}
		for i := 1; i <= b; i++ {
			fmt.Fscanf(in, "%s", &line)
			for j := 1; j <= w; j++ {
				maze[i][j] = cell{line[j-1] == '0', 0}
			}
		}
		s := solve()
		fmt.Fprintf(out, "%3d%3d%3d%3d%3d", s[0], s[1], s[2], s[3], s[4])
	}
}
