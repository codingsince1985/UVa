// UVa 146 - ID Codes

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func nextPermutation(p []string) bool {
	n := len(p)
	i := n - 1
	for i > 0 && p[i-1] >= p[i] {
		i--
	}
	if i == 0 {
		return false
	}

	mp := i
	for j := i + 1; j < n; j++ {
		if p[j] > p[i-1] && p[j] < p[mp] {
			mp = j
		}
	}
	p[mp], p[i-1] = p[i-1], p[mp]
	sort.Strings(p[i:n])
	return true
}

func main() {
	in, _ := os.Open("146.in")
	defer in.Close()
	out, _ := os.Create("146.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%s", &line); line == "#" {
			break
		}
		p := make([]string, len(line))
		for i := range line {
			p[i] = string(line[i])
		}
		if nextPermutation(p) {
			fmt.Fprintln(out, strings.Join(p, ""))
		} else {
			fmt.Fprintln(out, "No Successor")
		}
	}
}
