// UVa 227 - Puzzle

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = map[byte][2]int{'A': {-1, 0}, 'B': {1, 0}, 'R': {0, 1}, 'L': {0, -1}}

func findEmpty(puzzle [][]byte) (int, int) {
	for r := range puzzle {
		for c := range puzzle[r] {
			if puzzle[r][c] == ' ' {
				return r, c
			}
		}
	}
	return -1, -1
}

func move(puzzle [][]byte, moves string) bool {
	r, c := findEmpty(puzzle)
	for i := range moves {
		direction := directions[moves[i]]
		nr, nc := r+direction[0], c+direction[1]
		if nr < 0 || nr > 4 || nc < 0 || nc > 4 {
			return false
		}
		puzzle[r][c], puzzle[nr][nc] = puzzle[nr][nc], puzzle[r][c]
		r, c = nr, nc
	}
	return true
}

func output(out *os.File, puzzle [][]byte) {
	for i := range puzzle {
		for j := range puzzle[i] {
			if j > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, "%c", puzzle[i][j])
		}
		fmt.Fprintln(out)
	}
}

func main() {
	in, _ := os.Open("227.in")
	defer in.Close()
	out, _ := os.Create("227.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
here:
	for kase := 1; ; kase++ {
		puzzle := make([][]byte, 5)
		for i := 0; i < 5 && s.Scan(); i++ {
			if line = s.Text(); i == 0 && line == "Z" {
				break here
			}
			puzzle[i] = []byte(line)
		}
		var moves string
		for s.Scan() {
			if moves += s.Text(); strings.HasSuffix(moves, "0") {
				break
			}
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Puzzle #%d:\n", kase)
		if move(puzzle, moves[:len(moves)-1]) {
			output(out, puzzle)
		} else {
			fmt.Fprintln(out, "This puzzle has no final configuration.")
		}
	}
}
