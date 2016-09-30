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
	m := node{l.m + r.m, l.n + r.n}
	for !m.equals(n) {
		if n.lessThan(m) {
			path = append(path, 'L')
			r = m
		} else {
			path = append(path, 'R')
			l = m
		}
		m = node{l.m + r.m, l.n + r.n}
	}
	return string(path)
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
