// UVa 10227 - Forests

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func heardSameTrees(t1, t2 map[int]bool) bool {
	if len(t1) != len(t2) {
		return false
	}
	for k := range t1 {
		if !t2[k] {
			return false
		}
	}
	return true
}

func solve(grid []map[int]bool) int {
	p := len(grid)
	f := make([]int, p)
	for i := range f {
		f[i] = i
	}
	for i := 0; i < p-1; i++ {
		for j := i + 1; j < p; j++ {
			if heardSameTrees(grid[i], grid[j]) {
				if fi, fj := unionFind(i, f), unionFind(j, f); fi != fj {
					f[fi] = fj
				}
			}
		}
	}
	var count int
	for i := range f {
		if i == f[i] {
			count++
		}
	}
	return count
}

func main() {
	in, _ := os.Open("10227.in")
	defer in.Close()
	out, _ := os.Create("10227.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var p, t int
	var line string
	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	for s.Scan(); kase > 0 && s.Scan(); kase-- {
		fmt.Sscanf(s.Text(), "%d%d", &p, &t)
		grid := make([]map[int]bool, p)
		for i := range grid {
			grid[i] = make(map[int]bool)
		}
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			fmt.Sscanf(line, "%d%d", &p, &t)
			grid[p-1][t] = true
		}
		fmt.Fprintln(out, solve(grid))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
