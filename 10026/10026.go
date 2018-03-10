// UVa 10026 - Shoemaker's Problem

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type job struct{ no, time, fine int }

func output(out io.Writer, jobs []job) {
	first := true
	for _, j := range jobs {
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
	out, _ := os.Create("10026.out")
	defer out.Close()

	var kase, n int
	first := true
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &n)
		jobs := make([]job, n)
		for i := range jobs {
			fmt.Fscanf(in, "%d%d", &jobs[i].time, &jobs[i].fine)
			jobs[i].no = i + 1
		}
		sort.Slice(jobs, func(i, j int) bool {
			if jobs[i].time*jobs[j].fine == jobs[j].time*jobs[i].fine {
				return i < j
			}
			return jobs[i].time*jobs[j].fine < jobs[j].time*jobs[i].fine
		})
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		output(out, jobs)
	}
}
