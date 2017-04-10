// UVa 307 - Sticks

package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	n, l    int
	sticks  []int
	visited []bool
)

func dfs(level, curr, length int) bool {
	if level == n {
		return true
	}
	for i := curr; i < n; i++ {
		if !visited[i] {
			switch {
			case length+sticks[i] < l:
				visited[i] = true
				if dfs(level+1, i+1, length+sticks[i]) {
					return true
				}
				visited[i] = false
				if length == 0 {
					return false
				}
			case length+sticks[i] == l:
				visited[i] = true
				if dfs(level+1, 0, 0) {
					return true
				}
				visited[i] = false
				return false
			}
		}

	}
	return false
}

func solve(sum int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(sticks)))
	for i := n; i > 0; i-- {
		if sum%i == 0 {
			l = sum / i
			visited = make([]bool, n)
			if dfs(0, 0, 0) {
				return l
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("307.in")
	defer in.Close()
	out, _ := os.Create("307.out")
	defer out.Close()

	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		var sum int
		sticks = make([]int, n)
		for i := range sticks {
			fmt.Fscanf(in, "%d", &sticks[i])
			sum += sticks[i]
		}
		fmt.Fprintln(out, solve(sum))
	}
}
