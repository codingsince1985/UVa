// UVa 10344 - 23 out of 5

package main

import (
	"fmt"
	"os"
)

var (
	n       [5]int
	visited []bool
)

func dfs(level, result int) bool {
	if level == 5 {
		return result == 23
	}
	for i := range n {
		if !visited[i] {
			visited[i] = true
			if level == 0 {
				if dfs(level+1, n[i]) {
					return true
				}
			} else {
				if dfs(level+1, result+n[i]) || dfs(level+1, result-n[i]) || dfs(level+1, result*n[i]) {
					return true
				}
			}
			visited[i] = false
		}
	}
	return false
}

func main() {
	in, _ := os.Open("10344.in")
	defer in.Close()
	out, _ := os.Create("10344.out")
	defer out.Close()

	for {
		if fmt.Fscanf(in, "%d%d%d%d%d", &n[0], &n[1], &n[2], &n[3], &n[4]); n[0] == 0 && n[1] == 0 && n[2] == 0 && n[3] == 0 && n[4] == 0 {
			break
		}
		visited = make([]bool, 5)
		if dfs(0, 0) {
			fmt.Fprintln(out, "Possible")
		} else {
			fmt.Fprintln(out, "Impossible")
		}
	}
}
