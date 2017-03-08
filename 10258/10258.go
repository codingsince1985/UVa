// UVa 10258 - Contest Scoreboard

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type contestant struct {
	num, penalty, solved int
	problem              map[int]bool
}

func ranking(contest map[int]*contestant) []contestant {
	var c []contestant
	for _, v := range contest {
		c = append(c, *v)
	}
	sort.Slice(c, func(i, j int) bool {
		switch {
		case c[i].solved != c[j].solved:
			return c[i].solved > c[j].solved
		case c[i].penalty != c[j].penalty:
			return c[i].penalty < c[j].penalty
		default:
			return c[i].num < c[j].num
		}
	})
	return c
}

func main() {
	in, _ := os.Open("10258.in")
	defer in.Close()
	out, _ := os.Create("10258.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, c, p, t int
	var l, line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		contest := make(map[int]*contestant)
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			fmt.Sscanf(line, "%d%d%d%s", &c, &p, &t, &l)
			if l == "C" || l == "I" {
				if _, ok := contest[c]; !ok {
					contest[c] = &contestant{c, 0, 0, make(map[int]bool)}
				}
				switch {
				case l == "I" && !contest[c].problem[p]:
					contest[c].penalty += 20
				case l == "C":
					contest[c].penalty += t
					contest[c].solved++
					contest[c].problem[p] = true
				}
			}
		}
		for _, v := range ranking(contest) {
			fmt.Fprintln(out, v.num, v.solved, v.penalty)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
