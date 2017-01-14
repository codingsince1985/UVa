// UVa 10452 - Marcus

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	out        *os.File
	m, n       int
	stones     = "IEHOVA#"
	directions = map[string][2]int{"forth": {0, -1}, "left": {-1, 0}, "right": {1, 0}}
	grid       [][]byte
)

func findStart() int {
	for i, vi := range grid[m-1] {
		if vi == '@' {
			return i
		}
	}
	return -1
}

func dfs(step, x, y int, paths []string) {
	if step == len(stones) {
		fmt.Fprintln(out, strings.Join(paths, " "))
		return
	}
	for k, v := range directions {
		if nx, ny := x+v[0], y+v[1]; nx >= 0 && nx < n && ny >= 0 && ny < m && grid[ny][nx] == stones[step] {
			dfs(step+1, nx, ny, append(paths, k))
		}
	}
}

func main() {
	in, _ := os.Open("10452.in")
	defer in.Close()
	out, _ = os.Create("10452.out")
	defer out.Close()

	var line string
	var kase int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d", &m, &n)
		grid = make([][]byte, m)
		for i := range grid {
			fmt.Fscanf(in, "%s", &line)
			grid[i] = []byte(line)
		}
		dfs(0, findStart(), m-1, nil)
	}
}
