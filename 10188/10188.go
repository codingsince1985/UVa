// UVa 10188 - Automated Judge Script

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	s       *bufio.Scanner
	keepNum = func(r rune) rune {
		if r >= rune('0') && r <= rune('9') {
			return r
		}
		return -1
	}
)

func compare(solution, output string) string {
	switch {
	case solution == output:
		return "Accepted"
	case strings.Map(keepNum, solution) == strings.Map(keepNum, output):
		return "Presentation Error"
	default:
		return "Wrong Answer"
	}
}

func getMultiLine(n int) string {
	lines := make([]string, n)
	for i := range lines {
		s.Scan()
		lines[i] = s.Text()
	}
	return strings.Join(lines, "\n")
}

func main() {
	in, _ := os.Open("10188.in")
	defer in.Close()
	out, _ := os.Create("10188.out")
	defer out.Close()

	s = bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n int
	for kase := 1; s.Scan(); kase++ {
		if n, _ = strconv.Atoi(s.Text()); n == 0 {
			break
		}
		solution := getMultiLine(n)
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		output := getMultiLine(n)
		fmt.Fprintf(out, "Run #%d: %s\n", kase, compare(solution, output))
	}
}
