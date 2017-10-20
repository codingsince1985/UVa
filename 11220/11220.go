// UVa 11220 - Decoding the message.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(lines []string) []string {
	var words []string
	for _, line := range lines {
		var word, tmp string
		for idx, r := 0, strings.NewReader(line); ; {
			if _, err := fmt.Fscanf(r, "%s", &tmp); err != nil {
				break
			}
			if len(tmp) > idx {
				word += string(tmp[idx])
				idx++
			}
		}
		words = append(words, word)
	}
	return words
}

func main() {
	in, _ := os.Open("11220.in")
	defer in.Close()
	out, _ := os.Create("11220.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var t int
	var line string
	s.Scan()
	fmt.Sscanf(s.Text(), "%d", &t)
	s.Scan()
	for kase := 1; kase <= t; kase++ {
		var lines []string
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			lines = append(lines, line)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Case #%d:\n", kase)
		fmt.Fprintln(out, strings.Join(solve(lines), "\n"))
	}
}
