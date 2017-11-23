// UVa 315 - Network

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	matrix  [][]bool
	visited map[int]bool
)

func dfs(curr int) {
	visited[curr] = true
	for i := range matrix[curr] {
		if matrix[curr][i] && !visited[i] {
			dfs(i)
		}
	}
}

func solve() int {
	cnt := 0
	for i := range matrix {
		var turnOff []int
		for j := range matrix[i] {
			if matrix[i][j] {
				turnOff = append(turnOff, j)
				matrix[i][j], matrix[j][i] = false, false
			}
		}
		visited = make(map[int]bool)
		dfs((i + 1) % len(matrix))
		if len(visited)+1 != len(matrix) {
			cnt++
		}
		for _, j := range turnOff {
			matrix[i][j], matrix[j][i] = true, true
		}
	}
	return cnt
}

func main() {
	in, _ := os.Open("315.in")
	defer in.Close()
	out, _ := os.Create("315.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n, c1, c2 int
	var line string
	for s.Scan() {
		if fmt.Sscanf(s.Text(), "%d", &n); n == 0 {
			break
		}
		matrix = make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		for s.Scan() {
			if line = s.Text(); line == "0" {
				break
			}
			r := strings.NewReader(line)
			fmt.Fscanf(r, "%d", &c1)
			for {
				if _, err := fmt.Fscanf(r, "%d", &c2); err != nil {
					break
				}
				matrix[c1-1][c2-1], matrix[c2-1][c1-1] = true, true
			}
		}
		fmt.Fprintln(out, solve())
	}
}
