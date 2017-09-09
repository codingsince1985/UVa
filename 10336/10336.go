// UVa 10336 - Rank the Languages

package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	h, w       int
	langMap    map[byte]int
	directions = [][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
)

func dfs(world [][]byte, visited [][]bool, x, y int) {
	if !visited[x][y] {
		visited[x][y] = true
		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			if nx >= 0 && nx < h && ny >= 0 && ny < w && world[nx][ny] == world[x][y] {
				dfs(world, visited, nx, ny)
			}
		}
	}
}

func solve(world [][]byte) {
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	langMap = make(map[byte]int)
	for i := range world {
		for j := range world[i] {
			if !visited[i][j] {
				langMap[world[i][j]]++
				dfs(world, visited, i, j)
			}
		}
	}
}

type lang struct {
	l byte
	c int
}

func output(out *os.File, langMap map[byte]int) {
	var lc []lang
	for k, v := range langMap {
		lc = append(lc, lang{k, v})
	}
	sort.Slice(lc, func(i, j int) bool {
		if lc[i].c != lc[j].c {
			return lc[i].c > lc[j].c
		}
		return lc[i].l < lc[j].l
	})
	for _, v := range lc {
		fmt.Fprintf(out, "%c: %d\n", v.l, v.c)
	}
}

func main() {
	in, _ := os.Open("10336.in")
	defer in.Close()
	out, _ := os.Create("10336.out")
	defer out.Close()

	var n int
	var line string
	fmt.Fscanf(in, "%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(in, "%d%d", &h, &w)
		world := make([][]byte, h)
		for j := range world {
			fmt.Fscanf(in, "%s", &line)
			world[j] = []byte(line)
		}
		solve(world)
		fmt.Fprintf(out, "World #%d\n", i)
		output(out, langMap)
	}
}
