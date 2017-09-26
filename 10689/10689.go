// UVa 10689 - Yet another Number Sequence

package main

import (
	"fmt"
	"os"
)

var (
	pisano  = []int{60, 300, 1500, 15000}
	modular = []int{10, 100, 1000, 10000}
)

func solve(a, b, n, m int) int {
	length, mod := pisano[m], modular[m]
	seq := make([]int, length)
	seq[0], seq[1] = a%mod, b%mod
	for i := 2; i < length; i++ {
		seq[i] = (seq[i-1] + seq[i-2]) % mod
	}
	return seq[n%length]
}

func main() {
	in, _ := os.Open("10689.in")
	defer in.Close()
	out, _ := os.Create("10689.out")
	defer out.Close()

	var kase, a, b, n, m int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d%d%d", &a, &b, &n, &m)
		fmt.Fprintln(out, solve(a, b, n, m-1))
	}
}
