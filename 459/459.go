// UVa 459 - Graph Connectivity

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func find(x int, f []int) int {
	if f[x] != x {
		f[x] = find(f[x], f)
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
	line := s.Text()
	kase, _ := strconv.Atoi(line)
	s.Scan()
	s.Text()
	first := true
	for kase > 0 {
		kase--
		s.Scan()
		line = s.Text()
		num := line[0] - 'A' + 1
		f := make([]int, num)
		for i := range f {
			f[i] = i
		}

		for {
			s.Scan()
			if line = s.Text(); len(line) == 0 {
				break
			}
			a, b := line[0]-'A', line[1]-'A'
			fa, fb := find(int(a), f), find(int(b), f)
			if fa != fb {
				num--
				if fb < fa {
					f[fa] = fb
				} else {
					f[fb] = fa
				}
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
