// UVa 302 - John's trip

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct{ to, street int }

var (
	matrix  [][]node
	visited map[int]bool
	path    []string
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func even() bool {
	for _, m := range matrix {
		if len(m)%2 != 0 {
			return false
		}
	}
	return true
}

func dfs(curr int) {
	for _, n := range matrix[curr] {
		if !visited[n.street] {
			visited[n.street] = true
			dfs(n.to)
			path = append([]string{strconv.Itoa(n.street)}, path...)
		}
	}
}

func solve(n int, town [][3]int) string {
	var from int
	matrix = make([][]node, n)
	for i, t := range town {
		if i == 0 {
			from = min(t[0]-1, t[1]-1)
		}
		matrix[t[0]-1] = append(matrix[t[0]-1], node{t[1] - 1, t[2]})
		matrix[t[1]-1] = append(matrix[t[1]-1], node{t[0] - 1, t[2]})
	}
	if !even() {
		return "Round trip does not exist."
	}
	for i := range matrix {
		sort.Slice(matrix[i], func(j, k int) bool { return matrix[i][j].street < matrix[i][k].street })
	}
	visited = make(map[int]bool)
	path = nil
	dfs(from)
	return strings.Join(path, " ")
}

func main() {
	in, _ := os.Open("302.in")
	defer in.Close()
	out, _ := os.Create("302.out")
	defer out.Close()

	var x, y, z int
	for {
		var n int
		var town [][3]int
		for {
			if fmt.Fscanf(in, "%d%d", &x, &y); x == 0 && y == 0 {
				break
			}
			fmt.Fscanf(in, "%d", &z)
			n = max(n, z)
			town = append(town, [3]int{x, y, z})
		}
		if len(town) == 0 {
			break
		}
		fmt.Fprintf(out, "%s\n\n", solve(n, town))
	}
}
