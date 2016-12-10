// UVa 141 - The Spot Game

package main

import (
	"fmt"
	"os"
	"strings"
)

var n int

func turn(board [][]byte) [][]byte {
	newBoard := make([][]byte, n)
	for i := range newBoard {
		newBoard[i] = make([]byte, n)
		for j := range newBoard[i] {
			newBoard[i][j] = board[j][n-i-1]
		}
	}
	return newBoard
}

func stringer(board [][]byte) string {
	var str string
	for _, vi := range board {
		str += string(vi)
	}
	return str
}

func check(board [][]byte, cache map[string]bool) bool {
	for i := 0; i < 4; i++ {
		if cache[stringer(board)] {
			return true
		}
		board = turn(board)
	}
	cache[stringer(board)] = true
	return false
}

func move(board [][]byte, x, y int, c string) {
	if c == "+" {
		board[x-1][y-1] = '.'
	} else {
		board[x-1][y-1] = ' '
	}
}

func main() {
	in, _ := os.Open("141.in")
	defer in.Close()
	out, _ := os.Create("141.out")
	defer out.Close()

	var x, y int
	var c string
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		board := make([][]byte, n)
		for i := range board {
			board[i] = []byte(strings.Repeat(" ", n))
		}
		cache := make(map[string]bool)
		var match int
		for i := 0; i < 2*n; i++ {
			fmt.Fscanf(in, "%d%d%s", &x, &y, &c)
			move(board, x, y, c)
			if match == 0 && check(board, cache) {
				match = i
			}
		}
		if match == 0 {
			fmt.Fprintln(out, "Draw")
		} else {
			if match%2 == 0 {
				fmt.Fprintf(out, "Player 2 wins on move %d\n", match+1)
			} else {
				fmt.Fprintf(out, "Player 1 wins on move %d\n", match+1)
			}
		}
	}
}
