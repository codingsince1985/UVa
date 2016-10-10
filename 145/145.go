// UVa 145 - Gondwanaland Telecom

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	segments = []int{8, 18, 22, 24}
	duration = []int{8 * 60, 10 * 60, 4 * 60, 2 * 60}
	rateMap  = map[byte][]float64{
		'A': {0.02, 0.10, 0.06, 0.02},
		'B': {0.05, 0.25, 0.15, 0.05},
		'C': {0.13, 0.53, 0.33, 0.13},
		'D': {0.17, 0.87, 0.47, 0.17},
		'E': {0.30, 1.44, 0.80, 0.30},
	}
	daily = make(map[byte]float64)
)

func crossDay(h1, m1, h2, m2 int) bool { return h1 == h2 && m2 <= m1 || h1 > h2 }

func calc(h1, m1, h2, m2 int) []int {
	mins := make([]int, len(segments))
	for i, segment := range segments {
		if h1 < segment {
			if h2 < segment {
				mins[i] = (h2-h1)*60 - m1 + m2
				break
			} else {
				mins[i] = (segment-h1)*60 - m1
				h1, m1 = segment, 0
			}
		}
	}
	return mins
}

func main() {
	in, _ := os.Open("145.in")
	defer in.Close()
	out, _ := os.Create("145.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for step, rates := range rateMap {
		total := 0.0
		for i, rate := range rates {
			total += rate * float64(duration[i])
		}
		daily[step] = total
	}

	var step byte
	var line, phone string
	var h1, m1, h2, m2 int
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		r := strings.NewReader(line)
		fmt.Fscanf(r, "%c%s%d%d%d%d", &step, &phone, &h1, &m1, &h2, &m2)
		cross := crossDay(h1, m1, h2, m2)
		if cross {
			h1, m1, h2, m2 = h2, m2, h1, m1
		}
		mins := calc(h1, m1, h2, m2)
		rate, total := rateMap[step], 0.0
		for i, min := range mins {
			total += float64(min) * rate[i]
		}
		if cross {
			total = daily[step] - total
			for i := range mins {
				mins[i] = duration[i] - mins[i]
			}
		}
		mins[len(segments)-1] += mins[0]
		fmt.Fprintf(out, "%10s%6d%6d%6d%3c%8.2f\n", phone, mins[1], mins[2], mins[3], step, total)
	}
}
