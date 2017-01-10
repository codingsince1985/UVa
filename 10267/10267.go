// UVa 10267 - Graphical Editor

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	directions = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	m, n       int
)

func clear(picture [][]byte) {
	for i := range picture {
		for j := range picture[i] {
			picture[i][j] = 'O'
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int { return x + y - min(x, y) }

func fill(picture [][]byte, x, y int, c byte) {
	old := picture[y][x]
	picture[y][x] = c
	for _, direction := range directions {
		xn, yn := x+direction[1], y+direction[0]
		if xn >= 0 && xn < m && yn >= 0 && yn < n && picture[yn][xn] == old {
			fill(picture, xn, yn, c)
		}
	}
}

func main() {
	in, _ := os.Open("10267.in")
	defer in.Close()
	out, _ := os.Create("10267.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var x1, y1, x2, y2 int
	var line, tmp, c string
	var picture [][]byte
	for s.Scan() {
		if line = s.Text(); line == "X" {
			break
		}
		switch line[0] {
		case 'I':
			fmt.Sscanf(line, "%s%d%d", &tmp, &m, &n)
			picture = make([][]byte, n)
			for i := range picture {
				picture[i] = make([]byte, m)
			}
			fallthrough
		case 'C':
			clear(picture)
		case 'L':
			fmt.Sscanf(line, "%s%d%d%s", &tmp, &x1, &y1, &c)
			picture[y1-1][x1-1] = c[0]
		case 'V':
			fmt.Sscanf(line, "%s%d%d%d%s", &tmp, &x1, &y1, &y2, &c)
			for y := min(y1, y2) - 1; y <= max(y1, y2)-1; y++ {
				picture[y][x1-1] = c[0]
			}
		case 'H':
			fmt.Sscanf(line, "%s%d%d%d%s", &tmp, &x1, &x2, &y1, &c)
			for x := min(x1, x2) - 1; x <= max(x1, x2)-1; x++ {
				picture[y1-1][x] = c[0]
			}
		case 'K':
			fmt.Sscanf(line, "%s%d%d%d%d%s", &tmp, &x1, &y1, &x2, &y2, &c)
			for x := x1 - 1; x <= x2-1; x++ {
				for y := y1 - 1; y <= y2-1; y++ {
					picture[y][x] = c[0]
				}
			}
		case 'F':
			fmt.Sscanf(line, "%s%d%d%s", &tmp, &x1, &y1, &c)
			if picture[y1-1][x1-1] != c[0] { // so don't need to record visited when doing dfs
				fill(picture, x1-1, y1-1, c[0])
			}
		case 'S':
			fmt.Sscanf(line, "%s%s", &tmp, &c)
			fmt.Fprintln(out, c)
			for _, v := range picture {
				fmt.Fprintln(out, string(v))
			}
		}
	}
}
