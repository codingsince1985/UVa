// UVa 10038 - Jolly Jumpers

package main

import (
	"fmt"
	"os"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isJolly(lst []int) bool {
	length := len(lst)
	m := make([]bool, length)
	m[0] = true
	for i := 1; i < length; i++ {
		t := abs(lst[i] - lst[i-1])
		if m[t] {
			return false
		}
		m[t] = true
	}
	return true
}

func main() {
	in, _ := os.Open("10038.in")
	defer in.Close()
	out, _ := os.Create("10038.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		lst := make([]int, n)
		for i := range lst {
			fmt.Fscanf(in, "%d", &lst[i])
		}
		if isJolly(lst) {
			fmt.Fprintln(out, "Jolly")
		} else {
			fmt.Fprintln(out, "Not jolly")
		}
	}
}
