// UVa 632 - Compression (II)

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const length = 50

func solve(n int, s string) (int, []string) {
	ss := make([]string, n)
	ss[0] = s
	for i := 1; i < n; i++ {
		ss[i] = ss[i-1][1:] + string(ss[i-1][0])
	}
	s1 := ss[1]
	sort.Strings(ss)
	lines := n / length
	if n%length != 0 {
		lines++
	}
	lastBytes := make([][]byte, lines)
	for i, line := range ss {
		lastBytes[i/length] = append(lastBytes[i/length], line[n-1])
	}
	lastLines := make([]string, lines)
	for i := range lastBytes {
		lastLines[i] = string(lastBytes[i])
	}
	return sort.Search(n, func(i int) bool { return ss[i] >= s1 }), lastLines
}

func main() {
	in, _ := os.Open("632.in")
	defer in.Close()
	out, _ := os.Create("632.out")
	defer out.Close()

	var kase, n int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		var s string
		fmt.Fscanf(in, "\n%d", &n)
		for {
			fmt.Fscanf(in, "%s", &line)
			if s += line; len(s) == n {
				break
			}
		}
		s1, lines := solve(n, s)
		fmt.Fprintln(out, s1)
		fmt.Fprintln(out, strings.Join(lines, "\n"))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
