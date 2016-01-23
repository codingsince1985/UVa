// UVa 101 - The Blocks Problem

package main

import (
	"fmt"
	"os"
)

var (
	n int
	blocks [][]int
)

func initBlocks() {
	blocks = make([][]int, n)
	for i := range blocks {
		blocks[i] = make([]int, 1)
		blocks[i][0] = i
	}
}

func find(n int) int {
	for i := range blocks {
		for j := range blocks[i] {
			if blocks[i][j] == n {
				return i
			}
		}
	}
	return -1
}

func reset(n int) {
	stack := blocks[find(n)]
	for stack[len(stack) - 1] != n {
		blocks[stack[len(stack) - 1]] = append(blocks[stack[len(stack) - 1]], stack[len(stack) - 1])
		stack = stack[:len(stack) - 1]
	}
}

func move(n1, n2 int) {
	pos1 := find(n1)
	pos2 := find(n2)
	blocks[pos1] = blocks[pos1][:len(blocks[pos1]) - 1]
	blocks[pos2] = append(blocks[pos2], n1)
}

func pile(n1, n2 int) {
	pos1 := find(n1)
	pos2 := find(n2)
	stack := blocks[pos1]
	for i := range stack {
		if stack[i] == n1 {
			blocksToMove := stack[i:]
			blocks[pos1] = stack[:i]
			blocks[pos2] = append(blocks[pos2], blocksToMove...)
			break
		}
	}
}

func process(s1, s2 string, n1, n2 int) {
	switch {
	case s1 == "move":
		reset(n1)
	case s2 == "onto":
		reset(n2)
	}

	switch {
	case s1 == "move":
		move(n1, n2)
	case s1 == "pile":
		pile(n1, n2)
	}
}

func output(out *os.File) {
	for i := range blocks {
		fmt.Fprintf(out, "%d:", i)
		for j := range blocks[i] {
			fmt.Fprintf(out, " %d", blocks[i][j])
		}
		fmt.Fprintln(out)
	}
}

func main() {
	in, _ := os.Open("101.in")
	defer in.Close()
	out, _ := os.Create("101.out")
	defer out.Close()

	var n1, n2 int
	var s1, s2 string
	fmt.Fscanf(in, "%d", &n)
	initBlocks()
	for {
		fmt.Fscanf(in, "%s", &s1)
		if s1 == "quit" {
			break
		}
		fmt.Fscanf(in, "%d%s%d", &n1, &s2, &n2)
		process(s1, s2, n1, n2)
	}
	output(out)
}
