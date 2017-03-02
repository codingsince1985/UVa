// UVa 10131 - Is Bigger Smarter?

package main

import (
	"fmt"
	"os"
	"sort"
)

type (
	node  struct{ w, s int }
	nodes []node
)

func lis(n nodes) ([]int, int) {
	sort.Slice(n, func(i, j int) bool {
		if n[i].w != n[j].w {
			return n[i].w < n[j].w
		}
		return n[i].s > n[j].s
	})
	l := len(n)
	dp := make([]int, l+1)
	for i := range dp {
		dp[i] = 1
	}

	pre, max := make([]int, l), 0
	for i := 1; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if n[j-1].w > n[i-1].w && n[j-1].s < n[i-1].s {
				if dp[j] < dp[i]+1 {
					dp[j] = dp[i] + 1
					pre[j-1] = i - 1
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
		if pre[i] == max-1 {
			max--
			res[max] = n[i]
		}
	}
	return res
}

func output(out *os.File, res, orig []node) {
	fmt.Fprintln(out, len(res))
	for i := range res {
		for j := range orig {
			if res[i] == orig[j] {
				fmt.Fprintln(out, j+1)
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

	var n nodes
	var nd node
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &nd.w, &nd.s); err != nil {
			break
		}
		n = append(n, nd)
	}
	orig := make([]node, len(n))
	copy(orig, n)
	pre, max := lis(n)
	output(out, getResult(pre, max, n), orig)
}
