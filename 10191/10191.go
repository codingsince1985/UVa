// UVa 10191 - Longest Nap

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type job struct{ start, end int }

func buildJob(h1, m1, h2, m2 int) job { return job{h1*60 + m1, h2*60 + m2} }

func find(jobs []job) (int, int) {
	var start, period, max int
	t := 10 * 60
	for _, j := range jobs {
		if period = j.start - t; period > max {
			max = period
			start = t
		}
		t = j.end
	}
	if period = 18*60 - t; period > max {
		max = period
		start = t
	}
	return max, start
}

func output(out *os.File, kase, max, start int) {
	var duration string
	if max < 60 {
		duration = strconv.Itoa(max) + " minutes"
	} else {
		duration = fmt.Sprintf("%d hours and %d minutes", max/60, max%60)
	}
	startStr := fmt.Sprintf("%02d:%02d", start/60, start%60)
	fmt.Fprintf(out, "Day #%d: the longest nap starts at %s and will last for %s.\n", kase, startStr, duration)
}

func main() {
	in, _ := os.Open("10191.in")
	defer in.Close()
	out, _ := os.Create("10191.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var h1, m1, h2, m2 int
	for kase := 1; s.Scan(); kase++ {
		var js []job
		for n, _ := strconv.Atoi(s.Text()); n > 0 && s.Scan(); n-- {
			fmt.Sscanf(s.Text(), "%d:%d%d:%d", &h1, &m1, &h2, &m2)
			js = append(js, buildJob(h1, m1, h2, m2))
		}
		sort.Slice(js, func(i, j int) bool { return js[i].start < js[j].start })
		max, start := find(js)
		output(out, kase, max, start)
	}
}
