// UVa 10800 - Not That Kind of Graph

package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	MAX = 102
	MID = 52
)

func solve(out *os.File, line string) {
	graph := make([][]byte, MAX)
	for i := range graph {
		graph[i] = []byte(strings.Repeat(" ", MID))
	}
	currX, currY := 0, MID-1
	minY, maxY := MID-1, MID-1
	for i := range line {
		if currY < minY {
			minY = currY
		}
		if currY > maxY {
			maxY = currY
		}
		switch line[i] {
		case 'R':
			graph[currY][currX] = '/'
			currY--
		case 'F':
			currY++
			graph[currY][currX] = '\\'
		default:
			graph[currY][currX] = '_'
		}
		currX++
	}
	for i := minY; i <= maxY; i++ {
		fmt.Fprintln(out, "| "+strings.TrimRight(string(graph[i]), " "))
	}
	fmt.Fprintln(out, "+-"+strings.Repeat("-", currX+1))
}

func main() {
	in, _ := os.Open("10800.in")
	defer in.Close()
	out, _ := os.Create("10800.out")
	defer out.Close()

	var n int
	var line string
	fmt.Fscanf(in, "%d", &n)
	for kase := 1; kase <= n; kase++ {
		fmt.Fscanf(in, "%s", &line)
		fmt.Fprintf(out, "Case #%d:\n", kase)
		solve(out, line)
	}
}
