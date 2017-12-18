// UVa 10178 - Count the Faces.

package main

import (
	"fmt"
	"os"
)

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func solve(n int, edges [][2]int) int {
	f := make([]int, n)
	for i := range f {
		f[i] = i
	}
	for _, edge := range edges {
		if fa, fb := unionFind(edge[0], f), unionFind(edge[1], f); fa != fb {
			f[fa] = fb
		}
	}
	var count int
	for i := range f {
		if f[i] == i {
			count++
		}
	}
	return count
}

func main() {
	in, _ := os.Open("10178.in")
	defer in.Close()
	out, _ := os.Create("10178.out")
	defer out.Close()

	var n, e int
	var n1, n2 byte
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &e); err != nil {
			break
		}
		nodeMap := make(map[byte]int)
		var idx int
		edges := make([][2]int, e)
		for i := range edges {
			fmt.Fscanf(in, "%c %c\n", &n1, &n2)
			if _, ok := nodeMap[n1]; !ok {
				nodeMap[n1] = idx
				idx++
			}
			if _, ok := nodeMap[n2]; !ok {
				nodeMap[n2] = idx
				idx++
			}
			edges[i] = [2]int{nodeMap[n1], nodeMap[n2]}
		}
		fmt.Fprintln(out, 1+solve(len(nodeMap), edges)+e-idx)
	}
}
