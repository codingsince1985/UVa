// UVa 612 - DNA Sorting

package main

import (
	"fmt"
	"os"
	"sort"
)

func unsorted(s string) int {
	count := 0
	for i, l := 0, len(s); i < l-1; i++ {
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
		s := make([]string, m)
		for j := range s {
			fmt.Fscanf(in, "%s", &s[j])
		}
		sort.Slice(s, func(i, j int) bool { return unsorted(s[i]) < unsorted(s[j]) })
		for _, v := range s {
			fmt.Fprintln(out, v)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
