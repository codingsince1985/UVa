// UVa 468 - Key to Success

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type character struct {
	c byte
	f int
}

func freq(line string) []character {
	var chs []character
here:
	for i := range line {
		for j, ch := range chs {
			if ch.c == line[i] {
				chs[j].f++
				continue here
			}
		}
		chs = append(chs, character{line[i], 1})
	}
	sort.Slice(chs, func(i, j int) bool { return chs[i].f > chs[j].f })
	return chs
}

func solve(l1, l2 string) string {
	f1, f2 := freq(l1), freq(l2)
	m := make(map[byte]byte)
	for i, v := range f1 {
		m[f2[i].c] = v.c
	}
	result := make([]byte, len(l2))
	for i := range l2 {
		result[i] = m[l2[i]]
	}
	return string(result)
}

func main() {
	in, _ := os.Open("468.in")
	defer in.Close()
	out, _ := os.Create("468.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase int
	s.Scan()
	for fmt.Sscanf(s.Text(), "%d", &kase); kase > 0 && s.Scan(); kase-- {
		s.Scan()
		line := s.Text()
		s.Scan()
		fmt.Fprintln(out, solve(line, s.Text()))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
