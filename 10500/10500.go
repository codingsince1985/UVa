// UVa 10500 - Robot maps

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n, m, move      int
	directions      = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	store, robotMap [][]byte
	visited         [][]bool
	done            bool
)

func output(out io.Writer) {
	fmt.Fprintln(out)
	for _, row := range robotMap {
		fmt.Fprintln(out, strings.Repeat("|---", m)+"|")
		for _, cell := range row {
			fmt.Fprintf(out, "|-%c-", cell)
		}
		fmt.Fprintln(out, "|")
	}
	fmt.Fprintln(out, strings.Repeat("|---", m)+"|")
	fmt.Fprintf(out, "\nNUMBER OF MOVEMENTS: %d\n", move)
}

func dfs(x, y int) {
	visited[x][y], robotMap[x][y] = true, '0'
	for _, direction := range directions {
		if newX, newY := x+direction[0], y+direction[1]; newX >= 0 && newX < n && newY >= 0 && newY < m {
			robotMap[newX][newY] = store[newX][newY]
		}
	}
	for _, direction := range directions {
		if newX, newY := x+direction[0], y+direction[1]; newX >= 0 && newX < n && newY >= 0 && newY < m && store[newX][newY] == '0' && !visited[newX][newY] {
			move++
			dfs(newX, newY)
			if done {
				return
			}
		}
	}
	done = true
}

func solve(x, y int) {
	robotMap = make([][]byte, n)
	for i := range robotMap {
		robotMap[i] = []byte(strings.Repeat("?", m))
	}
	visited = make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	move, done = 0, false
	dfs(x, y)
}

func main() {
	in, _ := os.Open("10500.in")
	defer in.Close()
	out, _ := os.Create("10500.out")
	defer out.Close()

	var x, y int
	var ch byte
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		fmt.Fscanf(in, "%d%d", &x, &y)
		store = make([][]byte, n)
		for i := range store {
			store[i] = make([]byte, m)
			for j := range store[i] {
				if j > 0 {
					fmt.Fscanf(in, "%c", &ch)
				}
				fmt.Fscanf(in, "%c", &store[i][j])
			}
			fmt.Fscanln(in)
		}
		solve(x-1, y-1)
		output(out)
	}
}
