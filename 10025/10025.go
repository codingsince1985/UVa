// UVa 10025 - The ? 1 ? 2 ? ... ? n = k problem

package main

import (
	"fmt"
	"os"
)

func solve(k int) int {
	if k < 0 {
		k *= -1
	}
	i, sum := 0, 0
	for sum < k || (sum-k)%2 != 0 {
		i++
		sum += i
	}
	return i
}

func main() {
	in, _ := os.Open("10025.in")
	defer in.Close()
	out, _ := os.Create("10025.out")
	defer out.Close()

	var kase, k int
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%d", &k)
		fmt.Fprintln(out, solve(k))
		kase--
		if kase > 0 {
			fmt.Fprintln(out)
		}
	}
}
