// UVa 514 - Rails

package main

import (
	"fmt"
	"os"
)

func solve(coaches []int) bool {
	var stack []int
	n, idx, coach := len(coaches), 0, 1
	for {
		if coach <= n && (len(stack) == 0 || coaches[idx] != stack[len(stack)-1]) {
			stack = append(stack, coach)
			coach++
		} else if len(stack) > 0 && coaches[idx] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			idx++
		} else {
			break
		}
	}
	return len(stack) == 0 && idx == n
}

func main() {
	in, _ := os.Open("514.in")
	defer in.Close()
	out, _ := os.Create("514.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		coaches := make([]int, n)
	here:
		for {
			for i := range coaches {
				fmt.Fscanf(in, "%d", &coaches[i])
				if i == 0 && coaches[0] == 0 {
					fmt.Fprintln(out)
					break here
				}
			}
			if solve(coaches) {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}

	}
}
