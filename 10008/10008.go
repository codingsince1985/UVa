// UVa 10008 - What's Cryptanalysis?

package main

import (
	"fmt"
	"os"
	"sort"
)

type dict struct {
	c byte
	n int
}

func output(out *os.File, d map[byte]int) {
	arr := make([]dict, len(d))
	i := 0
	for k, v := range d {
		arr[i] = dict{k, v}
		i++
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].n == arr[j].n {
			return arr[i].c < arr[j].c
		}
		return arr[i].n > arr[j].n
	})
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
			d[c]++
		}
	}
	output(out, d)
}
