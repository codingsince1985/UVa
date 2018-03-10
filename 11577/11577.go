// UVa 11577 - Letter Frequency

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(out io.Writer, line string) {
	charMap := make(map[byte]int)
	var max int
	for i := range line {
		if c := line[i]; c >= 'a' && c <= 'z' {
			if charMap[c]++; charMap[c] > max {
				max = charMap[c]
			}
		}
	}
	var i byte
	for i = 'a'; i <= 'z'; i++ {
		if charMap[i] == max {
			fmt.Fprintf(out, "%c", i)
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("11577.in")
	defer in.Close()
	out, _ := os.Create("11577.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase int
	s.Scan()
	for fmt.Sscanf(s.Text(), "%d", &kase); kase > 0 && s.Scan(); kase-- {
		solve(out, strings.ToLower(s.Text()))
	}
}
