// UVa 496 - Simply Subsets

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse(line string) map[int]bool {
	var n int
	nm := make(map[int]bool)
	r := strings.NewReader(line)
	for {
		if _, err := fmt.Fscanf(r, "%d", &n); err != nil {
			break
		}
		nm[n] = true
	}
	return nm
}

func numIn(a, b map[int]bool) int {
	cnt := 0
	for k := range a {
		if b[k] {
			cnt++
		}
	}
	return cnt
}

func main() {
	in, _ := os.Open("496.in")
	defer in.Close()
	out, _ := os.Create("496.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		a := parse(s.Text())
		s.Scan()
		b := parse(s.Text())
		aInB := numIn(a, b)
		bInA := numIn(b, a)

		if aInB != 0 && bInA != 0 {
			switch {
			case aInB == len(b) && bInA == len(a):
				fmt.Fprintln(out, "A equals B")
			case aInB == len(a):
				fmt.Fprintln(out, "A is a proper subset of B")
			case bInA == len(b):
				fmt.Fprintln(out, "B is a proper subset of A")
			default:
				fmt.Fprintln(out, "I'm confused!")
			}
		} else {
			fmt.Fprintln(out, "A and B are disjoint")
		}
	}
}
