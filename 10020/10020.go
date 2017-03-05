// UVa 10020 - Minimal coverage

package main

import (
	"fmt"
	"os"
	"sort"
)

var out *os.File

func solve(l, r int, lines [][2]int) [][2]int {
	maxRight := l
	idx := -1
	for i, v := range lines {
		if v[0] > l {
			break
		}
		if v[1] > maxRight {
			maxRight = v[1]
			idx = i
		}
	}

	var ret [][2]int
	if idx == -1 {
		return ret
	}

	ret = append(ret, lines[idx])
	if maxRight < r {
		tmp := solve(maxRight, r, lines[idx+1:])
		if len(tmp) == 0 {
			return tmp
		}
		ret = append(ret, tmp...)
	}
	return ret
}

func output(r [][2]int) {
	fmt.Fprintln(out, len(r))
	for _, v := range r {
		fmt.Fprintln(out, v[0], v[1])
	}
}

func main() {
	in, _ := os.Open("10020.in")
	defer in.Close()
	out, _ = os.Create("10020.out")
	defer out.Close()

	var kase, m int
	var line [2]int
	for fmt.Fscanf(in, "%d\n", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &m)
		var lines [][2]int
		for {
			if fmt.Fscanf(in, "%d%d", &line[0], &line[1]); line[0] == 0 && line[1] == 0 {
				break
			}
			lines = append(lines, line)
		}
		sort.Slice(lines, func(i, j int) bool { return lines[i][0] < lines[j][0] })
		output(solve(0, m, lines))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
