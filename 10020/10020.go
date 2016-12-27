// UVa 10020 - Minimal coverage

package main

import (
	"fmt"
	"os"
	"sort"
)

var out *os.File

type lines [][2]int

func (l lines) Len() int { return len(l) }

func (l lines) Less(i, j int) bool { return l[i][0] < l[j][0] }

func (l lines) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func solve(l, r int, ls [][2]int) [][2]int {
	maxRight := l
	idx := -1
	for i, v := range ls {
		if v[0] <= l {
			if v[1] > maxRight {
				maxRight = v[1]
				idx = i
			}
		} else {
			break
		}
	}

	var ret [][2]int
	if idx == -1 {
		return ret
	}

	ret = append(ret, ls[idx])
	if maxRight < r {
		tmp := solve(maxRight, r, ls[idx+1:])
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

	var kase int
	for fmt.Fscanf(in, "%d\n", &kase); kase > 0; kase-- {
		var m int
		fmt.Fscanf(in, "\n%d", &m)
		var ls lines
		var l [2]int
		for {
			if fmt.Fscanf(in, "%d%d", &l[0], &l[1]); l[0] == 0 && l[1] == 0 {
				break
			}
			ls = append(ls, l)
		}
		sort.Sort(ls)
		output(solve(0, m, ls))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
