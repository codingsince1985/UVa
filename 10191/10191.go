// UVa 10191 - Longest Nap

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var out *os.File

type (
	job  struct{ start, end int }
	jobs []job
)

func (js jobs) Len() int { return len(js) }

func (js jobs) Less(i, j int) bool { return js[i].start < js[j].start }

func (js jobs) Swap(i, j int) { js[i], js[j] = js[j], js[i] }

func buildJob(h1, m1, h2, m2 int) job { return job{h1*60 + m1, h2*60 + m2} }

func find(js jobs) (int, int) {
	var start, period, max int
	t := 10 * 60
	for _, j := range js {
		period = j.start - t
		if period > max {
			max = period
			start = t
		}
		t = j.end
	}
	period = 18*60 - t
	if period > max {
		max = period
		start = t
	}
	return max, start
}

func output(kase, max, start int) {
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
	out, _ = os.Create("10191.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, h1, m1, h2, m2 int
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		var js jobs
		for n > 0 {
			s.Scan()
			line := s.Text()
			r := strings.NewReader(line)
			fmt.Fscanf(r, "%d:%d %d:%d", &h1, &m1, &h2, &m2)
			js = append(js, buildJob(h1, m1, h2, m2))
			n--
		}
		sort.Sort(js)
		max, start := find(js)
		kase++
		output(kase, max, start)
	}
}
