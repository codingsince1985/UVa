// UVa 10131 - Is Bigger Smarter?

package main

import (
	"fmt"
	"os"
	"sort"
)

type node struct{ w, s int }

func lis(nodes []node) ([]int, int) {
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].w != nodes[j].w {
			return nodes[i].w < nodes[j].w
		}
		return nodes[i].s > nodes[j].s
	})
	l := len(nodes)
	dp := make([]int, l+1)
	for i := range dp {
		dp[i] = 1
	}

	pre, max := make([]int, l), 0
	for i := 1; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if nodes[j-1].w > nodes[i-1].w && nodes[j-1].s < nodes[i-1].s {
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

func getResult(pre []int, max int, nodex []node) []node {
	res := make([]node, max)
	for i := len(pre) - 1; i >= 0; i-- {
		if pre[i] == max-1 {
			max--
			res[max] = nodex[i]
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

	var nodes []node
	var n node
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n.w, &n.s); err != nil {
			break
		}
		nodes = append(nodes, n)
	}
	orig := make([]node, len(nodes))
	copy(orig, nodes)
	pre, max := lis(nodes)
	output(out, getResult(pre, max, nodes), orig)
}
