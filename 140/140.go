// UVa 140 - Bandwidth

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

var (
	nodes, result []string
	nodeNum, min  int
	nodeMap       map[string]int
	matrix        [][]bool
)

func setNodes(tokens []string) {
	nodeSet := make(map[string]bool)
	for _, token := range tokens {
		n := strings.Split(token, ":")
		nodeSet[n[0]] = true
		for i := range n[1] {
			nodeSet[string(n[1][i])] = true
		}
	}
	nodeNum = len(nodeSet)
	nodes = make([]string, nodeNum)
	idx := 0
	for k := range nodeSet {
		nodes[idx] = k
		idx++
	}
	sort.Strings(nodes)
	nodeMap = make(map[string]int)
	for i, node := range nodes {
		nodeMap[node] = i
	}
}

func setMatrix(tokens []string) {
	matrix = make([][]bool, nodeNum)
	for i := range matrix {
		matrix[i] = make([]bool, nodeNum)
	}
	for _, token := range tokens {
		n := strings.Split(token, ":")
		idx1 := nodeMap[n[0]]
		for i := range n[1] {
			idx2 := nodeMap[string(n[1][i])]
			matrix[idx1][idx2], matrix[idx2][idx1] = true, true
		}
	}
}

func bandwidth(ordering []string) int {
	var bw int
	for i := 0; i < nodeNum-1; i++ {
		idx1 := nodeMap[ordering[i]]
		for j := i + 1; j < nodeNum; j++ {
			if matrix[idx1][nodeMap[ordering[j]]] && j-i > bw {
				bw = j - i
			}
		}
	}
	return bw
}

func dfs(ordering []string, visited []bool) {
	if len(ordering) == nodeNum {
		bw := bandwidth(ordering)
		if bw < min {
			min = bw
			copy(result, ordering)
		}
		return
	}
	for i, node := range nodes {
		if !visited[i] {
			visited[i] = true
			dfs(append(ordering, node), visited)
			visited[i] = false
		}
	}
}

func main() {
	in, _ := os.Open("140.in")
	defer in.Close()
	out, _ := os.Create("140.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		tokens := strings.Split(line, ";")
		setNodes(tokens)
		setMatrix(tokens)

		min = math.MaxInt32
		result = make([]string, nodeNum)
		dfs(nil, make([]bool, nodeNum))
		fmt.Fprintf(out, "%s -> %d\n", strings.Join(result, " "), min)
	}
}
