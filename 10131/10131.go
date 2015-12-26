// UVa 10131 - Is Bigger Smarter?

package main

import (
	"fmt"
	"os"
	"sort"
)

type node struct {
	w, s int
}

type nodes []node

func (a nodes) Len() int {
	return len(a)
}

func (a nodes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a nodes) Less(i, j int) bool {
	if a[i].w != a[j].w {
		return a[i].w < a[j].w
	} else {
		return a[i].s > a[j].s
	}
}

func lis(n []node) ([]int, int) {
	l := len(n)
	dp := make([]int, l + 1)
	for i := 0; i <= l; i++ {
		dp[i] = 1
	}

	pre := make([]int, l)
	max := 0

	for i := 1; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if n[j - 1].w > n[i - 1].w && n[j - 1].s < n[i - 1].s {
				if dp[j] < dp[i] + 1 {
					dp[j] = dp[i] + 1
					pre[j - 1] = i - 1
					if dp[j] > max {
						max = dp[j]
					}
				}
			}
		}
	}
	return pre, max
}

func getResult(pre []int, max int, n []node) []node {
	res := make([]node, max)
	for i := len(pre) - 1; i >= 0; i-- {
		if pre[i] == max - 1 {
			max --
			res[max] = n[i]
		}
	}
	return res
}

func clone(n []node) []node {
	m := make([]node, len(n))
	copy(m, n)
	return m
}

func output(out *os.File, res, orig []node) {
	l := len(res)
	fmt.Fprintln(out, l)
	for i := 0; i < l; i++ {
		for j := 0; j < len(orig); j++ {
			if res[i] == orig[j] {
				fmt.Fprintln(out, j + 1)
				break
			}
		}
	}
}

func main() {
	in, _ := os.Open("10131.in")
	defer in.Close()
	out, _ := os.Create("10131.out")
	defer out.Close()

	n := make(nodes, 0)
	var t1, t2 int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &t1, &t2); err != nil {
			break
		}
		n = append(n, node{t1, t2})
	}
	orig := clone(n)
	sort.Sort(n)
	pre, max := lis(n)
	res := getResult(pre, max, n)
	output(out, res, orig)
}
