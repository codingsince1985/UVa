// UVa 10010 - Where's Waldorf?

package main

import (
	"fmt"
	"os"
)

var (
	m, n       int
	directions = [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
)

func found(x, y int, grid [][]byte, chars []byte) bool {
	if grid[x][y] != chars[0] {
		return false
	}

here:
	for _, dir := range directions {
		px, py := x, y
		for j := 1; j < len(chars); j++ {
			px += dir[0]
			py += dir[1]
			if px < 0 || px >= m || py < 0 || py >= n || grid[px][py] != chars[j] {
				continue here
			}
		}
		return true
	}
	return false
}

func find(out *os.File, grid [][]byte, chars []byte) {
	for x := range grid {
		for y := range grid[x] {
			if found(x, y, grid, chars) {
				fmt.Fprintln(out, x+1, y+1)
				return
			}
		}
	}
}

func toLower(word string) []byte {
	chars := []byte(word)
	for i := range chars {
		if chars[i] >= 'A' && chars[i] <= 'Z' {
			chars[i] = 'a' + chars[i] - 'A'
		}
	}
	return chars
}

func main() {
	in, _ := os.Open("10010.in")
	defer in.Close()
	out, _ := os.Create("10010.out")
	defer out.Close()

	var kase, words int
	var line, word string
	for fmt.Fscanf(in, "%d\n", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d%d", &m, &n)
		grid := make([][]byte, m)
		for i := range grid {
			fmt.Fscanf(in, "%s", &line)
			grid[i] = toLower(line)
		}
		for fmt.Fscanf(in, "%d", &words); words > 0; words-- {
			fmt.Fscanf(in, "%s", &word)
			find(out, grid, toLower(word))
		}
	}
}
