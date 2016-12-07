// UVa 10298 - Power Strings

package main

import (
	"bufio"
	"fmt"
	"os"
)

func kmp(line string) int {
	p := make([]int, len(line)+1)
	for i := range p {
		p[i] = -1
	}
	var i, j int
	for i, j = 1, -1; i < len(line); i++ {
		for j >= 0 && line[j+1] != line[i] {
			j = p[j]
		}
		if line[j+1] == line[i] {
			j++
		}
		p[i] = j
	}
	return j
}

func main() {
	in, _ := os.Open("10298.in")
	defer in.Close()
	out, _ := os.Create("10298.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "." {
			break
		}
		if p, l := kmp(line), len(line); l%(l-p-1) != 0 {
			fmt.Fprintln(out, "1")
		} else {
			fmt.Fprintln(out, l/(l-p-1))
		}
	}
}
