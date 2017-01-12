// UVa 612 - DNA Sorting

package main

import (
	"fmt"
	"os"
	"sort"
)

type dna []string

func (d dna) Len() int { return len(d) }

func (d dna) Less(i, j int) bool { return unsorted(d[i]) < unsorted(d[j]) }

func (d dna) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

func unsorted(s string) int {
	l := len(s)
	count := 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] > s[j] {
				count++
			}
		}
	}
	return count
}

func main() {
	in, _ := os.Open("612.in")
	defer in.Close()
	out, _ := os.Create("612.out")
	defer out.Close()

	var kase, n, m int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d%d", &n, &m)
		s := make(dna, m)
		for j := range s {
			fmt.Fscanf(in, "%s", &s[j])
		}
		sort.Sort(s)
		for _, v := range s {
			fmt.Fprintln(out, v)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
