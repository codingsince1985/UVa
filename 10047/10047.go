// UVa 10047 - The Monocycle

package main

import (
	"fmt"
	"os"
)

type (
	cell struct{ y, x int }
	node struct {
		cell
		direction, colour, time int
	}
)

var (
	m, n       int
	grid       [][]byte
	turns      = [2]int{1, 3}
	directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

func bfs(s, e cell) int {
	visited := make([][][4][5]bool, m)
	for i := range visited {
		visited[i] = make([][4][5]bool, n)
	}
	visited[s.y][s.x][0][0] = true
	for queue := []node{{s, 0, 0, 0}}; len(queue) > 0; {
		curr := queue[0]
		queue = queue[1:]
		y, x, d, c, t := curr.y, curr.x, curr.direction, curr.colour, curr.time
		if c == 0 && curr.cell == e {
			return t
		}
		for _, turn := range turns {
			if nd := (d + turn) % 4; !visited[y][x][nd][c] {
				visited[y][x][nd][c] = true
				queue = append(queue, node{cell{y, x}, nd, c, t + 1})
			}
		}
		if ny, nx, nc := y+directions[d][0], x+directions[d][1], (c+1)%5; ny >= 0 && ny < m && nx >= 0 && nx < n && grid[ny][nx] != '#' && !visited[ny][nx][d][nc] {
			visited[ny][nx][d][nc] = true
			queue = append(queue, node{cell{ny, nx}, d, nc, t + 1})
		}
	}
	return -1
}

func solve() int {
	var s, e cell
	for i, row := range grid {
		for j, v := range row {
			switch v {
			case 'S':
				s = cell{i, j}
			case 'T':
				e = cell{i, j}
			}
		}
	}
	return bfs(s, e)
}

func main() {
	in, _ := os.Open("10047.in")
	defer in.Close()
	out, _ := os.Create("10047.out")
	defer out.Close()

	var line string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 && n == 0 {
			break
		}
		grid = make([][]byte, m)
		for i := range grid {
			fmt.Fscanf(in, "%s", &line)
			grid[i] = []byte(line)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Case #%d\n", kase)
		if dist := solve(); dist == -1 {
			fmt.Fprintln(out, "destination not reachable")
		} else {
			fmt.Fprintf(out, "minimum time = %d sec\n", dist)
		}
	}
}
