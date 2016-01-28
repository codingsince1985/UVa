// UVa 572 - Oil Deposits

package main

import (
	"fmt"
	"os"
)

type node struct {
	oil bool
	idx int
}

var (
	m, n int
	g    [][]node
)

func dfs(i, j, count int) {
	if i < 0 || i >= m || j < 0 || j >= n || g[i][j].idx > 0 || !g[i][j].oil {
		return
	}

	g[i][j].idx = count
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				dfs(i+dx, j+dy, count)
			}
		}
	}
}

func solve() int {
	count := 0
	for i := range g {
		for j := range g[i] {
			if g[i][j].idx == 0 && g[i][j].oil {
				count++
				dfs(i, j, count)
			}
		}
	}
	return count
}

func main() {
	in, _ := os.Open("572.in")
	defer in.Close()
	out, _ := os.Create("572.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 && n == 0 {
			break
		}
		g = make([][]node, m)
		for i := range g {
			g[i] = make([]node, n)
			fmt.Fscanf(in, "%s", &line)
			for j := range g[i] {
				g[i][j] = node{line[j] == '@', 0}
			}
		}
		fmt.Fprintln(out, solve())
	}
}
