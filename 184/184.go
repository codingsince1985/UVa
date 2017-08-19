// UVa 184 - Laser Lines

package main

import (
	"fmt"
	"os"
	"sort"
)

var out *os.File

type node struct{ x, y int }

func sameLine(n1, n2, n3 node) bool { return (n2.x-n1.x)*(n3.y-n1.y) == (n2.y-n1.y)*(n3.x-n1.x) }

func output(nodes []node, lines [][]int) {
	if len(lines) == 0 {
		fmt.Fprintln(out, "No lines were found")
		return
	}
	fmt.Fprintln(out, "The following lines were found:")
	for _, line := range lines {
		for _, n := range line {
			fmt.Fprintf(out, "(%4d,%4d)", nodes[n].x, nodes[n].y)
		}
		fmt.Fprintln(out)
	}
}

func solve(nodes []node) {
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].x != nodes[j].x {
			return nodes[i].x < nodes[j].x
		}
		return nodes[i].y < nodes[j].y
	})
	n := len(nodes)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var lines [][]int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if !visited[i][j] {
				line := []int{i, j}
				for k := j + 1; k < n; k++ {
					if sameLine(nodes[i], nodes[j], nodes[k]) {
						line = append(line, k)
					}
				}
				if len(line) >= 3 {
					lines = append(lines, line)
					for l := 0; l < len(line)-1; l++ {
						for m := l + 1; m < len(line); m++ {
							visited[line[l]][line[m]] = true
						}
					}
				}
			}
		}
	}
	output(nodes, lines)
}

func main() {
	in, _ := os.Open("184.in")
	defer in.Close()
	out, _ = os.Create("184.out")
	defer out.Close()

	var x, y int
	var nodes []node
here:
	for {
		for {
			if fmt.Fscanf(in, "%d%d", &x, &y); x == 0 && y == 0 {
				if len(nodes) == 0 {
					break here
				}
				solve(nodes)
				nodes = nil
				continue
			}
			nodes = append(nodes, node{x, y})
		}
	}
}
