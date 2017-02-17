// UVa 260 - Il Gioco dell'X

package main

import (
	"fmt"
	"os"
)

var (
	n     int
	board [][]byte
)

func dfs(i, j int) bool {
	if j >= n || board[i][j] != 'b' {
		return false
	}
	return i == n-1 || dfs(i+1, j) || dfs(i+1, j+1)
}

func blackWins() bool {
	for j := range board {
		if dfs(0, j) {
			return true
		}
	}
	return false
}

func main() {
	in, _ := os.Open("260.in")
	defer in.Close()
	out, _ := os.Create("260.out")
	defer out.Close()

	var line string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		board = make([][]byte, n)
		for i := range board {
			fmt.Fscanf(in, "%s", &line)
			board[i] = []byte(line)
		}
		if fmt.Fprint(out, kase); blackWins() {
			fmt.Fprintln(out, " B")
		} else {
			fmt.Fprintln(out, " W")
		}
	}
}
