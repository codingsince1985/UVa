// UVa 10364 - Square

package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	edge    int
	sticks  []int
	visited []bool
)

func dfs(side, length int) bool {
	if side == 4 || length == edge && dfs(side+1, 0) {
		return true
	}
	for i, stick := range sticks {
		if newLength := length + stick; !visited[i] && newLength <= edge {
			visited[i] = true
			if dfs(side, newLength) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}

func main() {
	in, _ := os.Open("10364.in")
	defer in.Close()
	out, _ := os.Create("10364.out")
	defer out.Close()

	var n, m int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		var sum int
		fmt.Fscanf(in, "%d", &m)
		sticks = make([]int, m)
		for i := range sticks {
			fmt.Fscanf(in, "%d", &sticks[i])
			sum += sticks[i]
		}
		sort.Ints(sticks)
		visited = make([]bool, m)
		if edge = sum / 4; sum%4 == 0 && sticks[m-1] <= edge && dfs(0, 0) {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
	}
}
