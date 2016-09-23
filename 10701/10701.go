// UVa 10701 - Pre, in and post

package main

import (
	"fmt"
	"os"
)

func findRoot(root byte, inOrder []byte) int {
	for i := range inOrder {
		if root == inOrder[i] {
			return i
		}
	}
	return -1
}

func postOrder(preOrder, inOrder []byte) []byte {
	if len(preOrder) == 0 {
		return nil
	}
	idx := findRoot(preOrder[0], inOrder)
	return append(append(postOrder(preOrder[1:idx+1], inOrder[:idx]), postOrder(preOrder[idx+1:], inOrder[idx+1:])...), preOrder[0])
}

func main() {
	in, _ := os.Open("10701.in")
	defer in.Close()
	out, _ := os.Create("10701.out")
	defer out.Close()

	var c, n int
	var preOrder, inOrder string
	fmt.Fscanf(in, "%d", &c)
	for c > 0 {
		fmt.Fscanf(in, "%d%s%s", &n, &preOrder, &inOrder)
		fmt.Fprintln(out, string(postOrder([]byte(preOrder), []byte(inOrder))))
		c--
	}
}
