// UVa 10026 - Shoemaker's Problem

package main

import (
	"fmt"
	"os"
	"sort"
)

var out *os.File

type (
	job  struct{ no, time, fine int }
	jobs []job
)

func (jbs jobs) Len() int { return len(jbs) }

func (jbs jobs) Less(i, j int) bool {
	if jbs[i].time*jbs[j].fine == jbs[j].time*jbs[i].fine {
		return i < j
	}
	return jbs[i].time*jbs[j].fine < jbs[j].time*jbs[i].fine
}

func (jbs jobs) Swap(i, j int) { jbs[i], jbs[j] = jbs[j], jbs[i] }

func output(jbs jobs) {
	first := true
	for _, j := range jbs {
		if first {
			first = false
		} else {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, j.no)
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("10026.in")
	defer in.Close()
	out, _ = os.Create("10026.out")
	defer out.Close()

	var kase, n int
	first := true
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &n)
		jbs := make(jobs, n)
		for i := range jbs {
			fmt.Fscanf(in, "%d%d", &jbs[i].time, &jbs[i].fine)
			jbs[i].no = i + 1
		}
		sort.Sort(jbs)
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		output(jbs)
	}
}
