// UVa 10324 - Zeros and Ones

package main

import (
	"fmt"
	"os"
)

func buildDict(line string) []int {
	dict := make([]int, len(line))
	lastDigit := line[0]
	lastIndex := 0
	for i := range line {
		if line[i] == lastDigit {
			dict[i] = lastIndex
		} else {
			lastDigit = line[i]
			lastIndex = i
			dict[i] = lastIndex
		}
	}
	return dict
}

func main() {
	in, _ := os.Open("10324.in")
	defer in.Close()
	out, _ := os.Create("10324.out")
	defer out.Close()

	var line string
	var l, r, n int
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		dict := buildDict(line)
		fmt.Fprintf(out, "Case %d:\n", kase)
		for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
			if fmt.Fscanf(in, "%d%d", &l, &r); dict[l] == dict[r] {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		}
	}
}
