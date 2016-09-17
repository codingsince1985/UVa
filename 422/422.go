// UVa 422 - Word-Search Wonder

package main

import (
	"fmt"
	"os"
)

var (
	l           int
	matrix      [][]byte
	directions  = [][2]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}
	start, stop [2]int
)

func follow(x, y int, direction [2]int, word string) bool {
	if len(word) == 0 {
		stop = [2]int{x - direction[0] + 1, y - direction[1] + 1}
		return true
	}
	if x < 0 || y < 0 || x >= l || y >= l || matrix[x][y] != word[0] {
		return false
	}
	return follow(x+direction[0], y+direction[1], direction, word[1:])
}

func found(word string) bool {
	for i := range matrix {
		for j, ch := range matrix[i] {
			if ch == word[0] {
				start = [2]int{i + 1, j + 1}
				for _, direction := range directions {
					if follow(i+direction[0], j+direction[1], direction, word[1:]) {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	in, _ := os.Open("422.in")
	defer in.Close()
	out, _ := os.Create("422.out")
	defer out.Close()

	var word string
	fmt.Fscanf(in, "%d", &l)
	matrix = make([][]byte, l)
	for i := range matrix {
		fmt.Fscanf(in, "%s", &word)
		matrix[i] = []byte(word)
	}
	for {
		if fmt.Fscanf(in, "%s", &word); word == "0" {
			break
		}
		if found(word) {
			fmt.Fprintf(out, "%d,%d %d,%d\n", start[0], start[1], stop[0], stop[1])
		} else {
			fmt.Fprintln(out, "Not found")
		}
	}
}
