// UVa 10077 - The Stern-Brocot Number System

package main

import (
	"fmt"
	"os"
)

type node struct{ m, n int }

func (n1 node) lessThan(n2 node) bool { return n1.m*n2.n < n2.m*n1.n }

func (n1 node) equals(n2 node) bool { return n1.m == n2.m && n1.n == n2.n }

func binarySearch(n node) string {
	var path []byte
	l, r := node{0, 1}, node{1, 0}
	for {
		m := node{l.m + r.m, l.n + r.n}
		switch {
		case m.equals(n):
			return string(path)
		case n.lessThan(m):
			path = append(path, 'L')
			r = m
		default:
			path = append(path, 'R')
			l = m
		}
	}
}

func main() {
	in, _ := os.Open("10077.in")
	defer in.Close()
	out, _ := os.Create("10077.out")
	defer out.Close()

	var m, n int
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 1 && n == 1 {
			break
		}
		fmt.Fprintln(out, binarySearch(node{m, n}))
	}
}
