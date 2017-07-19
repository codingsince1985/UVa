// UVa 501 - Black Box

package main

import (
	"fmt"
	"os"
	"sort"
)

func insert(slice []int, pos, value int) []int {
	slice = slice[0 : len(slice)+1]
	copy(slice[pos+1:], slice[pos:])
	slice[pos] = value
	return slice
}

func solve(out *os.File, m, n int, a, u []int) {
	blackBox := make([]int, 0, m)
	var idx int
	for _, ai := range a {
		for u[idx] == len(blackBox) {
			fmt.Fprintln(out, blackBox[idx])
			if idx++; idx == n {
				return
			}
		}
		blackBox = insert(blackBox, sort.Search(len(blackBox), func(j int) bool { return blackBox[j] >= ai }), ai)
	}
}

func main() {
	in, _ := os.Open("501.in")
	defer in.Close()
	out, _ := os.Create("501.out")
	defer out.Close()

	var kase, m, n int
	first := true
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%d%d", &m, &n)
		a := make([]int, m)
		for i := range a {
			fmt.Fscanf(in, "%d", &a[i])
		}
		u := make([]int, n)
		for i := range u {
			fmt.Fscanf(in, "%d", &u[i])
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		solve(out, m, n, a, u)
	}
}
