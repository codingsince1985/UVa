// UVa 10152 - ShellSort

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int

func transform(src, dest []string) []string {
	idx1, idx2 := n-1, n-1
	for ; idx1 >= 0; idx1-- {
		if src[idx1] == dest[idx2] {
			idx2--
		}
	}
	var toMove []string
	for ; idx2 >= 0; idx2-- {
		toMove = append(toMove, dest[idx2])
	}
	return toMove
}

func main() {
	in, _ := os.Open("10152.in")
	defer in.Close()
	out, _ := os.Create("10152.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	for k, _ := strconv.Atoi(s.Text()); k > 0; k-- {
		s.Scan()
		n, _ = strconv.Atoi(s.Text())
		src := make([]string, n)
		for i := range src {
			s.Scan()
			src[i] = s.Text()
		}
		dest := make([]string, n)
		for i := range dest {
			s.Scan()
			dest[i] = s.Text()
		}
		fmt.Fprintln(out, strings.Join(transform(src, dest), "\n")+"\n")
	}
}
