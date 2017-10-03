// UVa 466 - Mirror, Mirror

package main

import (
	"fmt"
	"os"
)

var (
	d      int
	fm, to [][]byte
)

type node struct {
	matrix [][]byte
	path   []int
}

func initialize() {
	fm = make([][]byte, d)
	for i := range fm {
		fm[i] = make([]byte, d)
	}
	to = make([][]byte, d)
	for i := range to {
		to[i] = make([]byte, d)
	}
}

func set(s1, s2 string, i int) {
	for c := range s1 {
		fm[i][c] = s1[c]
		to[i][c] = s2[c]
	}
}

func isSame(m1, m2 [][]byte) bool {
	for i := range m1 {
		for j := range m1[i] {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}
	return true
}

func rotate(m [][]byte) [][]byte {
	ret := make([][]byte, d)
	for i := range m {
		ret[i] = make([]byte, d)
		for j := range m[i] {
			ret[i][j] = m[d-1-j][i]
		}
	}
	return ret
}

func reflect(m [][]byte) [][]byte {
	ret := make([][]byte, d)
	for i := range m {
		ret[i] = make([]byte, d)
		for j := range m[i] {
			ret[i][j] = m[d-1-i][j]
		}
	}
	return ret
}

func bfs() []int {
	for queue := []node{{fm, nil}}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		if isSame(curr.matrix, to) {
			return curr.path
		}
		if len(curr.path) > 3 {
			break
		}
		n := rotate(curr.matrix)
		queue = append(queue, node{n, append(curr.path, 1)})
		n = rotate(n)
		queue = append(queue, node{n, append(curr.path, 2)})
		n = rotate(n)
		queue = append(queue, node{n, append(curr.path, 3)})
		n = reflect(curr.matrix)
		queue = append(queue, node{n, append(curr.path, 4)})
	}
	return []int{-1}
}

func output(out *os.File, path []int, kase int) {
	fmt.Fprintf(out, "Pattern %d was ", kase)
	switch {
	case len(path) == 0:
		fmt.Fprintln(out, "preserved.")
	case path[0] == -1:
		fmt.Fprintln(out, "improperly transformed.")
	default:
		for i, v := range path {
			if i != 0 {
				fmt.Fprint(out, " and ")
			}
			if v < 4 {
				fmt.Fprintf(out, "rotated %d degrees", v*90)
			} else {
				fmt.Fprint(out, "reflected vertically")
			}
		}
		fmt.Fprintln(out, ".")
	}
}

func main() {
	in, _ := os.Open("466.in")
	defer in.Close()
	out, _ := os.Create("466.out")
	defer out.Close()

	var s1, s2 string
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d", &d); err != nil {
			break
		}
		initialize()
		for i := range fm {
			fmt.Fscanf(in, "%s%s", &s1, &s2)
			set(s1, s2, i)
		}
		output(out, bfs(), kase)
	}
}
