// UVa 459 - Graph Connectivity

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

func main() {
	in, _ := os.Open("459.in")
	defer in.Close()
	out, _ := os.Create("459.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	s.Scan()
	first := true
	for ; kase > 0 && s.Scan(); kase-- {
		line := s.Text()
		num := int(line[0] - 'A' + 1)
		f := make([]int, num)
		for i := range f {
			f[i] = i
		}

		for s.Scan() {
			if line = s.Text(); len(line) == 0 {
				break
			}
			a, b := line[0]-'A', line[1]-'A'
			if fa, fb := unionFind(int(a), f), unionFind(int(b), f); fa != fb {
				num--
				f[fa] = fb
			}
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, num)
	}
}
