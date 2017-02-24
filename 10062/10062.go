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
		sort.Slice(lst, func(i, j int) bool {
			if lst[i].freq == lst[j].freq {
				return lst[i].ch > lst[j].ch
			}
			return lst[i].freq < lst[j].freq
		})
		for _, v := range lst {
			fmt.Fprintln(out, v.ch, v.freq)
		}
	}
}
