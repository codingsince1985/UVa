// UVa 10008 - What's Cryptanalysis?

package main

import (
	"fmt"
	"os"
	"sort"
)

type di struct {
	c byte
	n int
}

type dict []di

func (a dict) Len() int { return len(a) }

func (a dict) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a dict) Less(i, j int) bool {
	if a[i].n == a[j].n {
		return a[i].c < a[j].c
	}
	return a[i].n > a[j].n
}

func output(out *os.File, d map[byte]int) {
	arr := make(dict, len(d))
	i := 0
	for k, v := range d {
		arr[i] = di{k, v}
		i++
	}
	sort.Sort(arr)
	for _, v := range arr {
		fmt.Fprintf(out, "%c %d\n", v.c, v.n)
	}
}

func main() {
	in, _ := os.Open("10008.in")
	defer in.Close()
	out, _ := os.Create("10008.out")
	defer out.Close()

	var n, ln int
	var c byte
	d := make(map[byte]int)
	fmt.Fscanf(in, "%d", &n)
	for {
		if fmt.Fscanf(in, "%c", &c); c == '\n' {
			if ln++; ln == n {
				break
			}
		}
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		if c >= 'A' && c <= 'Z' {
			cnt := d[c]
			cnt++
			d[c] = cnt
		}
	}
	output(out, d)
}
