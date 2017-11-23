// UVa 677 - All Walks of length "n" from the first node

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	paths   []string
	n       int
	matrix  [][]bool
	visited map[int]bool
)

func dfs(curr, level int, path []string) {
	if level == n {
		paths = append(paths, "("+strings.Join(path, ",")+")")
		return
	}
	for i, v := range matrix[curr] {
		if v && !visited[i] {
			visited[i] = true
			dfs(i, level+1, append(path, strconv.Itoa(i+1)))
			visited[i] = false
		}
	}
}

func solve() string {
	paths = nil
	visited = map[int]bool{0: true}
	dfs(0, 0, []string{"1"})
	return strings.Join(paths, "\n")
}

func main() {
	in, _ := os.Open("677.in")
	defer in.Close()
	out, _ := os.Create("677.out")
	defer out.Close()

	var m, tmp int
	first := true
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &m, &n); err != nil {
			break
		}
		matrix = make([][]bool, m)
		for i := range matrix {
			matrix[i] = make([]bool, m)
		}
		for i := range matrix {
			for j := range matrix[i] {
				if fmt.Fscanf(in, "%d", &tmp); tmp == 1 {
					matrix[i][j] = true
				}
			}
		}
		fmt.Fscanf(in, "%d", &tmp)
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, solve())
	}
}
