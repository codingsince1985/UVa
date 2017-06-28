// UVa 10001 - Garden of Eden

package main

import (
	"fmt"
	"os"
)

const max = 33

var (
	table               = [][]int{{0, 0, 0}, {0, 0, 1}, {0, 1, 0}, {0, 1, 1}, {1, 0, 0}, {1, 0, 1}, {1, 1, 0}, {1, 1, 1}}
	n                   int
	state, automaton, o []int
)

func dfs(level int) bool {
	if level == n {
		return o[0] == o[level] && o[1] == o[level+1]
	}
	for i, row := range table {
		if state[level] == automaton[i] && (level == 0 || o[level] == row[0] && o[level+1] == row[1]) {
			if level == 0 {
				o[level], o[level+1] = row[0], row[1]
			}
			o[level+2] = row[2]
			if dfs(level + 1) {
				return true
			}
		}
	}
	return false
}

func solve(a int, s string) bool {
	state = make([]int, n)
	for i := range s {
		state[i] = int(s[i] - '0')
	}
	automaton = make([]int, 8)
	for i := range automaton {
		automaton[i] = a % 2
		a /= 2
	}
	o = make([]int, max)
	return dfs(0)
}

func main() {
	in, _ := os.Open("10001.in")
	defer in.Close()
	out, _ := os.Create("10001.out")
	defer out.Close()

	var a int
	var s string
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &a, &n); err != nil {
			break
		}
		fmt.Fscanf(in, "%s", &s)
		if solve(a, s) {
			fmt.Fprintln(out, "REACHABLE")
		} else {
			fmt.Fprintln(out, "GARDEN OF EDEN")
		}
	}
}
