// UVa 10116 - Robot Motion

package main

import (
	"fmt"
	"os"
)

var (
	r, c       int
	directions = map[byte][2]int{'W': {0, -1}, 'E': {0, 1}, 'N': {-1, 0}, 'S': {1, 0}}
)

func move(grid [][]byte, x, y int) (int, int) {
	visited := make(map[[2]int]int)
	cnt := 0
	for x >= 0 && x < r && y >= 0 && y < c {
		curr := [2]int{x, y}
		if i, ok := visited[curr]; ok {
			return i, cnt
		}
		visited[curr] = cnt
		direction := directions[grid[x][y]]
		x += direction[0]
		y += direction[1]
		cnt++
	}
	return 0, cnt
}

func main() {
	in, _ := os.Open("10116.in")
	defer in.Close()
	out, _ := os.Create("10116.out")
	defer out.Close()

	var p int
	var line string
	for {
		if fmt.Fscanf(in, "%d%d%d", &r, &c, &p); r == 0 && c == 0 && p == 0 {
			break
		}
		grid := make([][]byte, r)
		for i := range grid {
			fmt.Fscanf(in, "%s", &line)
			grid[i] = []byte(line)
		}
		loop, cnt := move(grid, 0, p-1)
		if loop == 0 {
			fmt.Fprintf(out, "%d step(s) to exit\n", cnt)
		} else {
			fmt.Fprintf(out, "%d step(s) before a loop of %d step(s)\n", loop, cnt-loop)
		}
	}
}
