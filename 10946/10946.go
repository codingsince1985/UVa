// UVa 10946 - You want what filled?

package main

import (
	"fmt"
	"os"
	"sort"
)

type hole struct {
	label byte
	size  int
}

var (
	x, y, count int
	land        [][]byte
	visited     [][]bool
	directions  = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	holes       []hole
)

func dfs(i, j int, cell byte) {
	visited[i][j] = true
	count++
	for _, direction := range directions {
		if newJ, newI := j+direction[1], i+direction[0]; newJ >= 0 && newJ < y && newI >= 0 && newI < x && land[newI][newJ] == cell && !visited[newI][newJ] {
			dfs(newI, newJ, cell)
		}
	}
}

func solve() []hole {
	holes = nil
	visited = make([][]bool, x)
	for i := range visited {
		visited[i] = make([]bool, y)
	}
	for i, row := range land {
		for j, cell := range row {
			if !visited[i][j] {
				count = 0
				dfs(i, j, cell)
				if cell != '.' {
					holes = append(holes, hole{cell, count})
				}
			}
		}
	}
	sort.Slice(holes, func(i, j int) bool {
		if holes[i].size == holes[j].size {
			return holes[i].label < holes[j].label
		}
		return holes[i].size > holes[j].size
	})
	return holes
}

func main() {
	in, _ := os.Open("10946.in")
	defer in.Close()
	out, _ := os.Create("10946.out")
	defer out.Close()

	var line string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &x, &y); x == 0 && y == 0 {
			break
		}
		land = make([][]byte, x)
		for i := range land {
			fmt.Fscanf(in, "%s", &line)
			land[i] = []byte(line)
		}
		fmt.Fprintf(out, "Problem %d:\n", kase)
		holes := solve()
		for _, h := range holes {
			fmt.Fprintf(out, "%c %d\n", h.label, h.size)
		}
	}
}
