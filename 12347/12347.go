// UVa 12347 - Binary Search Tree

package main

import (
	"fmt"
	"os"
)

var out *os.File

func postOrder(n []int) {
	if l := len(n); l > 0 {
		idx := 1
		for i := 1; i < l; i++ {
			if n[i] > n[0] {
				idx = i
				break
			}
		}
		postOrder(n[1:idx])
		postOrder(n[idx:])
		fmt.Fprintln(out, n[0])
	}
}

func main() {
	in, _ := os.Open("12347.in")
	defer in.Close()
	out, _ = os.Create("12347.out")
	defer out.Close()

	var tmp int
	var n []int
	for {
		if _, err := fmt.Fscanf(in, "%d", &tmp); err != nil {
			break
		}
		n = append(n, tmp)
	}
	postOrder(n)
}
