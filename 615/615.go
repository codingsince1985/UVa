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
	for k := range nodeMap {
		f[k] = k
	}
	for _, edge := range edges {
		if r1, r2 := unionFind(edge.n1, f), unionFind(edge.n2, f); r1 != r2 {
			f[r1] = r2
		}
	}
	count := 0
	for k := range nodeMap {
		if f[k] == k {
			if count++; count > 1 {
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

	var e edge
here:
	for kase := 1; ; kase++ {
		var edges []edge
	there:
		for s.Scan() {
			line := s.Text()
			if line == "" {
				continue
			}
			if line[0] == '-' {
				break here
			}
			for r := strings.NewReader(line); ; {
				if _, err := fmt.Fscanf(r, "%d%d", &e.n1, &e.n2); err != nil || e.n1 == 0 && e.n2 == 0 {
					if e.n1 == 0 && e.n2 == 0 {
						break there
					}
					break
				}
				edges = append(edges, e)
			}
		}
		if fmt.Fprintf(out, "Case %d ", kase); isTree(edges) {
			fmt.Fprintln(out, "is a tree.")
		} else {
			fmt.Fprintln(out, "is not a tree.")
		}
	}
}
