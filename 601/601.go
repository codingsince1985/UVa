// UVa 601 - The PATH

package main

import (
	"fmt"
	"os"
)

var (
	n          int
	board      [][]byte
	visited    [][]bool
	directions = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
)

type cell struct{ x, y int }

func dfs(y, x int, color byte) bool {
	if color == 'W' && x == n-1 || color == 'B' && y == n-1 {
		return true
	}
	for _, direction := range directions {
		if newY, newX := y+direction[1], x+direction[0]; newY >= 0 && newY < n && newX >= 0 && newX < n && !visited[newY][newX] && board[newY][newX] == color {
			visited[newY][newX] = true
			if dfs(newY, newX, color) {
				return true
			}
			visited[newY][newX] = false
		}
	}
	return false
}

func moveToWin() int {
	var unfilled []cell
	for i, row := range board {
		for j, c := range row {
			if c == 'U' {
				unfilled = append(unfilled, cell{j, i})
			}
		}
	}
	for _, c := range unfilled {
		board[c.y][c.x] = 'W'
		if whiteWins() {
			return 3
		}
		board[c.y][c.x] = 'U'
	}
	for _, c := range unfilled {
		board[c.y][c.x] = 'B'
		if blackWins() {
			return 4
		}
		board[c.y][c.x] = 'U'
	}
	return -1
}

func whiteWins() bool {
	for i := range board {
		if board[i][0] == 'W' {
			visited[i][0] = true
			if dfs(i, 0, 'W') {
				return true
			}
			visited[i][0] = false
		}
	}
	return false
}

func blackWins() bool {
	for i := range board {
		if board[0][i] == 'B' {
			visited[0][i] = true
			if dfs(0, i, 'B') {
				return true
			}
			visited[0][i] = false
		}
	}
	return false
}

func solve() int {
	visited = make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	switch {
	case whiteWins():
		return 1
	case blackWins():
		return 2
	default:
		return moveToWin()
	}
}

func main() {
	in, _ := os.Open("601.in")
	defer in.Close()
	out, _ := os.Create("601.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fscanln(in)
		board = make([][]byte, n)
		for i := range board {
			fmt.Fscanf(in, "%s", &line)
			board[i] = []byte(line)
		}
		switch solve() {
		case 1:
			fmt.Fprintln(out, "White has a winning path.")
		case 2:
			fmt.Fprintln(out, "Black has a winning path.")
		case 3:
			fmt.Fprintln(out, "White can win in one move.")
		case 4:
			fmt.Fprintln(out, "Black can win in one move.")
		default:
			fmt.Fprintln(out, "There is no winning path.")
		}
		fmt.Fscanln(in)
	}
}
