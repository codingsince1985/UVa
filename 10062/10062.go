// UVa 10062 - Tell me the frequencies!

package main

import (
	"fmt"
	"os"
	"sort"
)

type (
	char struct {
		ch   byte
		freq int
	}
	chars []char
)

func (a chars) Len() int { return len(a) }

func (a chars) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a chars) Less(i, j int) bool {
	if a[i].freq == a[j].freq {
		return a[i].ch > a[j].ch
	}
	return a[i].freq < a[j].freq
}

func main() {
	in, _ := os.Open("10062.in")
	defer in.Close()
	out, _ := os.Create("10062.out")
	defer out.Close()

	var line string
	first := true
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		if !first {
			fmt.Fprintln(out)
		} else {
			first = false
		}
		table := make(map[byte]int)
		for i := range line {
			if line[i] >= 32 && line[i] < 128 {
				table[line[i]]++
			}
		}
		var lst chars
		for k, v := range table {
			lst = append(lst, char{k, v})
		}
		sort.Sort(lst)
		for _, v := range lst {
			fmt.Fprintln(out, v.ch, v.freq)
		}
	}
}
