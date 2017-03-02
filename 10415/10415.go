// UVa 10415 - Eb Alto Saxophone Player

package main

import (
	"fmt"
	"os"
)

var noteMap = map[byte][]bool{
	'c': {false, true, true, true, false, false, true, true, true, true},
	'd': {false, true, true, true, false, false, true, true, true, false},
	'e': {false, true, true, true, false, false, true, true, false, false},
	'f': {false, true, true, true, false, false, true, false, false, false},
	'g': {false, true, true, true, false, false, false, false, false, false},
	'a': {false, true, true, false, false, false, false, false, false, false},
	'b': {false, true, false, false, false, false, false, false, false, false},
	'C': {false, false, true, false, false, false, false, false, false, false},
	'D': {true, true, true, true, false, false, true, true, true, false},
	'E': {true, true, true, true, false, false, true, true, false, false},
	'F': {true, true, true, true, false, false, true, false, false, false},
	'G': {true, true, true, true, false, false, false, false, false, false},
	'A': {true, true, true, false, false, false, false, false, false, false},
	'B': {true, true, false, false, false, false, false, false, false, false},
}

func solve(line string) []int {
	count := make([]int, 10)
	var fingers []bool
	for i := range line {
		for j, finger := range noteMap[line[i]] {
			if (i == 0 || !fingers[j]) && finger {
				count[j]++
			}
		}
		fingers = noteMap[line[i]]
	}
	return count
}

func main() {
	in, _ := os.Open("10415.in")
	defer in.Close()
	out, _ := os.Create("10415.out")
	defer out.Close()

	var kase int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &line)
		output := fmt.Sprint(solve(line))
		fmt.Fprintln(out, output[1:len(output)-1])
	}
}
