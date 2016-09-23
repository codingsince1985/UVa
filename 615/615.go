// UVa 615 - Is It A Tree?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type edge struct{ n1, n2 int }

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isTree(edges []edge) bool {
	if len(edges) == 0 {
		return true
	}
	var maxNode int
	nodeMap := make(map[int]bool)
	for _, edge := range edges {
		nodeMap[edge.n1], nodeMap[edge.n2] = true, true
		maxNode = max(maxNode, max(edge.n1, edge.n2))
	}

	f := make([]int, maxNode+1)
	for k, _ := range nodeMap {
		f[k] = k
	}
	for _, edge := range edges {
		r1, r2 := unionFind(edge.n1, f), unionFind(edge.n2, f)
		if r1 != r2 {
			f[r1] = r2
		}
	}
	count := 0
	for k, _ := range nodeMap {
		if f[k] == k {
			count++
			if count > 1 {
				return false
			}
		}
	}

	pointingTo := make(map[int]int)
	for _, edge := range edges {
		pointingTo[edge.n2]++
		if pointingTo[edge.n2] > 1 {
			return false
		}
	}
	return len(nodeMap) == len(pointingTo)+1
}

func main() {
	in, _ := os.Open("615.in")
	defer in.Close()
	out, _ := os.Create("615.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n1, n2, kase int
here:
	for {
		done := false
		kase++
		var edges []edge
		for s.Scan() {
			line := s.Text()
			if line == "" {
				continue
			}
			if line[0] == '-' {
				break here
			}
			r := strings.NewReader(line)
			for {
				if _, err := fmt.Fscanf(r, "%d%d", &n1, &n2); err != nil || n1 == 0 && n2 == 0 {
					if n1 == 0 && n2 == 0 {
						done = true
					}
					break
				}
				edges = append(edges, edge{n1, n2})
			}
			if done {
				break
			}
		}
		if isTree(edges) {
			fmt.Fprintf(out, "Case %d is a tree.\n", kase)
		} else {
			fmt.Fprintf(out, "Case %d is not a tree.\n", kase)
		}
	}
}
