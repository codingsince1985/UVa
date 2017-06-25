// UVa 11244 - Counting Stars

package main

import (
	"fmt"
	"os"
)

var (
	r, c, size int
	sky        [][]byte
	directions = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	visited    [][]bool
)

func dfs(i, j int) {
	for _, direction := range directions {
		if newI, newJ := i+direction[0], j+direction[1]; newI >= 0 && newI < r && newJ >= 0 && newJ < c && !visited[newI][newJ] && sky[newI][newJ] == '*' {
			size++
			visited[newI][newJ] = true
			dfs(newI, newJ)
		}
	}
}

func solve() int {
	var count int
	visited = make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}
	for i, row := range sky {
		for j, cell := range row {
			if !visited[i][j] {
				visited[i][j] = true
				if cell == '*' {
					size = 1
					dfs(i, j)
					if size == 1 {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	in, _ := os.Open("11244.in")
	defer in.Close()
	out, _ := os.Create("11244.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%d%d", &r, &c); r == 0 && c == 0 {
			break
		}
		sky = make([][]byte, r)
		for i := range sky {
			fmt.Fscanf(in, "%s", &line)
			sky[i] = []byte(line)
		}
		fmt.Fprintln(out, solve())
	}
}
