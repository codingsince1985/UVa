// UVa 657 - The die is cast

package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	w, h, count int
	picture     [][]byte
	directions  = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited     [][]bool
	dots        []int
)

func dfs2(y, x int) {
	for _, direction := range directions {
		if newY, newX := y+direction[1], x+direction[0]; newY >= 0 && newY < h && newX >= 0 && newX < w && picture[newY][newX] == 'X' && !visited[newY][newX] {
			visited[newY][newX] = true
			dfs2(newY, newX)
		}
	}
}

func dfs1(y, x int) {
	for _, direction := range directions {
		if newY, newX := y+direction[1], x+direction[0]; newY >= 0 && newY < h && newX >= 0 && newX < w && !visited[newY][newX] {
			visited[newY][newX] = true
			switch picture[newY][newX] {
			case '*':
				dfs1(newY, newX)
			case 'X':
				dfs2(newY, newX)
				count++
			}
		}
	}
}

func dfs(y, x int) {
	for _, direction := range directions {
		if newY, newX := y+direction[1], x+direction[0]; newY >= 0 && newY < h && newX >= 0 && newX < w && !visited[newY][newX] {
			visited[newY][newX] = true
			switch picture[newY][newX] {
			case '.':
				dfs(newY, newX)
			case '*':
				count = 0
				dfs1(newY, newX)
				if count > 0 {
					dots = append(dots, count)
				}
			}
		}
	}
}

func startPoint() (int, int) {
	for y, row := range picture {
		for x := range row {
			if picture[y][x] == '.' {
				return y, x
			}
		}
	}
	return -1, -1
}

func solve() []int {
	visited = make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	y, x := startPoint()
	visited[y][x] = true
	dots = nil
	dfs(y, x)
	sort.Ints(dots)
	return dots
}

func main() {
	in, _ := os.Open("657.in")
	defer in.Close()
	out, _ := os.Create("657.out")
	defer out.Close()

	var line string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &w, &h); w == 0 && h == 0 {
			break
		}
		picture = make([][]byte, h)
		for i := range picture {
			fmt.Fscanf(in, "%s", &line)
			picture[i] = []byte(line)
		}
		fmt.Fprintf(out, "Throw %d\n", kase)
		str := fmt.Sprintln(solve())
		fmt.Fprintf(out, "%s\n\n", str[1:len(str)-2])
	}
}
