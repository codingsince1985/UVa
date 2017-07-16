// UVa 10651 - Pebble Solitaire

package main

import (
	"fmt"
	"os"
)

var (
	board   []bool
	removed int
)

func movable(i int) []int {
	if i < len(board)-2 && board[i] && board[i+1] && !board[i+2] {
		return []int{i, i + 1, i + 2}
	}
	if i > 2 && board[i] && board[i-1] && !board[i-2] {
		return []int{i - 2, i - 1, i}
	}
	return nil
}

func dfs(level int) {
	if level > removed {
		removed = level
	}
	for i := range board {
		if idx := movable(i); idx != nil {
			for _, j := range idx {
				board[j] = !board[j]
			}
			dfs(level + 1)
			for _, j := range idx {
				board[j] = !board[j]
			}
		}
	}
}

func solve() int {
	var count int
	for _, p := range board {
		if p {
			count++
		}
	}
	removed = 0
	dfs(0)
	return count - removed
}

func main() {
	in, _ := os.Open("10651.in")
	defer in.Close()
	out, _ := os.Create("10651.out")
	defer out.Close()

	var kase int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &line)
		board = make([]bool, len(line))
		for i := range line {
			board[i] = line[i] == 'o'
		}
		fmt.Fprintln(out, solve())
	}
}
