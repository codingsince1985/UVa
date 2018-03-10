// UVa 699 - The Falling Leaves

package main

import (
	"fmt"
	"io"
	"os"
)

const max = 200

type node struct {
	n           int
	left, right *node
}

var (
	low, high int
	sums      []int
	in        io.ReadCloser
)

func nextNumber() int {
	var n int
	fmt.Fscanf(in, "%d", &n)
	return n
}

func sumTree(root *node, idx int) {
	if root != nil {
		if idx < low {
			low = idx
		}
		if idx > high {
			high = idx
		}
		sums[idx] += root.n
		sumTree(root.left, idx-1)
		sumTree(root.right, idx+1)
	}
}

func buildTree(n int) *node {
	if n != -1 {
		return &node{n, buildTree(nextNumber()), buildTree(nextNumber())}
	}
	return nil
}

func solve(n int) []int {
	root := buildTree(n)
	low, high = max-1, 0
	sums = make([]int, max)
	sumTree(root, max/2)
	return sums[low : high+1]
}

func main() {
	in, _ = os.Open("699.in")
	defer in.Close()
	out, _ := os.Create("699.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if n = nextNumber(); n == -1 {
			break
		}
		fmt.Fprintf(out, "Case %d:\n", kase)
		output := fmt.Sprint(solve(n))
		fmt.Fprintf(out, "%s\n\n", output[1:len(output)-1])
	}
}
