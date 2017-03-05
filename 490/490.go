// UVa 490 - Rotating Sentences

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func rotate(lines []string, max int) []string {
	rotated := make([]string, max)
	for i := range rotated {
		for j := len(lines) - 1; j >= 0; j-- {
			if i < len(lines[j]) {
				rotated[i] += string(lines[j][i])
			} else {
				rotated[i] += " "
			}
		}
	}
	return rotated
}

func main() {
	in, _ := os.Open("490.in")
	defer in.Close()
	out, _ := os.Create("490.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var max int
	var lines []string
	for s.Scan() {
		line := s.Text()
		if size := len(line); max < size {
			max = size
		}
		lines = append(lines, line)
	}

	rotated := rotate(lines, max)
	for _, v := range rotated {
		fmt.Fprintln(out, strings.TrimRight(v, " "))
	}
}
