// UVa 705 - Slash Maze

package main

import (
	"fmt"
	"os"
)

var (
	w, h       int
	maze       [][]bool
	directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

type node struct{ y, x int }

func bfs(y, x int) (bool, int) {
	var longest int
	circle := true
	for queue := []node{{y, x}}; len(queue) > 0; queue = queue[1:] {
		longest++
		curr := queue[0]
		for _, direction := range directions {
			if ny, nx := curr.y+direction[0], curr.x+direction[1]; ny >= 0 && ny < 3*h && nx >= 0 && nx < 3*w && !maze[ny][nx] {
				if ny == 0 || ny == 3*h-1 || nx == 0 || nx == 3*w-1 {
					circle = false
				}
				maze[ny][nx] = true
				queue = append(queue, node{ny, nx})
			}
		}
	}
	return circle, longest / 3
}

func solve() (int, int) {
	var count, longest int
	for i, row := range maze {
		for j, cell := range row {
			if !cell {
				maze[i][j] = true
				if ok, l := bfs(i, j); ok {
					count++
					if l > longest {
						longest = l
					}
				}
			}
		}
	}
	return count, longest
}

func main() {
	in, _ := os.Open("705.in")
	defer in.Close()
	out, _ := os.Create("705.out")
	defer out.Close()

	var line string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &w, &h); w == 0 && h == 0 {
			break
		}
		maze = make([][]bool, 3*h)
		for i := range maze {
			maze[i] = make([]bool, 3*w)
		}
		for i := 0; i < h; i++ {
			fmt.Fscanf(in, "%s", &line)
			for j := range line {
				maze[3*i+1][3*j+1] = true
				if line[j] == '/' {
					maze[3*i][3*j+2] = true
					maze[3*i+2][3*j] = true
				} else {
					maze[3*i][3*j] = true
					maze[3*i+2][3*j+2] = true
				}
			}
		}
		fmt.Fprintf(out, "Maze #%d:\n", kase)
		if circle, longest := solve(); circle > 0 {
			fmt.Fprintf(out, "%d Cycles; the longest has length %d.\n", circle, longest)
		} else {
			fmt.Fprintln(out, "There are no cycles.")
		}
		fmt.Fprintln(out)
	}
}
