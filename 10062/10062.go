// UVa 10062 - Tell me the frequencies!

package main

import (
	"fmt"
	"os"
	"sort"
)

type char struct {
	ch   byte
	freq int
}

func solve(line string) []char {
	table := make(map[byte]int)
	for i := range line {
		if line[i] >= 32 && line[i] < 128 {
			table[line[i]]++
		}
	}
	var lst []char
	for k, v := range table {
		lst = append(lst, char{k, v})
	}
	sort.Slice(lst, func(i, j int) bool {
		if lst[i].freq == lst[j].freq {
			return lst[i].ch > lst[j].ch
		}
		return lst[i].freq < lst[j].freq
	})
	return lst
}

func main() {
	in, _ := os.Open("10062.in")
	defer in.Close()
	out, _ := os.Create("10062.out")
	defer out.Close()

	var line string
	for first := true; ; {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		for _, v := range solve(line) {
			fmt.Fprintln(out, v.ch, v.freq)
		}
	}
}
