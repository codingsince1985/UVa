// UVa 124 - Following Orders

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	nodes      []byte
	nodeMap    map[byte]int
	constraint [][]bool
	orders     []string
)

func getNodeMap() {
	nodeMap = make(map[byte]int)
	for i, node := range nodes {
		nodeMap[node] = i
	}
}

func getConstraint(line string) {
	constraint = make([][]bool, len(nodes))
	for i := range constraint {
		constraint[i] = make([]bool, len(nodes))
	}
	n := []byte(strings.Replace(line, " ", "", -1))
	for i := 0; i < len(n)-1; i += 2 {
		constraint[nodeMap[n[i]]][nodeMap[n[i+1]]] = true
	}
}

func valid(n byte, order []byte) bool {
	for _, node := range order {
		if node == n || constraint[nodeMap[n]][nodeMap[node]] {
			return false
		}
	}
	return true
}

func backtracking(order []byte) {
	if len(order) == len(nodes) {
		orders = append(orders, string(order))
		return
	}
	for _, node := range nodes {
		if valid(node, order) {
			backtracking(append(order, node))
		}
	}
}

func main() {
	in, _ := os.Open("124.in")
	defer in.Close()
	out, _ := os.Create("124.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	first := true
	for s.Scan() {
		nodes = []byte(strings.Replace(s.Text(), " ", "", -1))
		getNodeMap()
		s.Scan()
		getConstraint(s.Text())
		orders = nil
		backtracking(nil)
		sort.Strings(orders)
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, strings.Join(orders, "\n"))
	}
}
