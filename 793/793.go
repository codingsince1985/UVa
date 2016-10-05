// UVa 793 - Network Connections

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func main() {
	in, _ := os.Open("793.in")
	defer in.Close()
	out, _ := os.Create("793.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var a, b int
	var command byte
	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	s.Scan()
	s.Text()
	for kase > 0 && s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		f := make([]int, num+1)
		for i := range f {
			f[i] = i
		}
		ok, ko := 0, 0
		for s.Scan() {
			line := s.Text()
			if line == "" {
				break
			}
			r := strings.NewReader(line)
			fmt.Fscanf(r, "%c%d%d", &command, &a, &b)
			fa, fb := unionFind(a, f), unionFind(b, f)
			if command == 'c' {
				if fa != fb {
					f[fa] = fb
				}
			} else {
				if fa != fb {
					ko++
				} else {
					ok++
				}
			}
		}
		fmt.Fprintf(out, "%d,%d\n", ok, ko)
		kase--
		if kase > 0 {
			fmt.Fprintln(out)
		}
	}
}
