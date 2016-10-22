// UVa 326 - Extrapolation Using a Difference Table

package main

import (
	"fmt"
	"os"
)

func solve(num []int, n, k int) int {
	for j := n - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			num[i] = num[i+1] - num[i]
		}
	}
	for j := 0; j < k; j++ {
		for i := 1; i < n; i++ {
			num[i] += num[i-1]
		}
	}
	return num[n-1]
}

func main() {
	in, _ := os.Open("326.in")
	defer in.Close()
	out, _ := os.Create("326.out")
	defer out.Close()

	var n, k int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		num := make([]int, n)
		for i := range num {
			fmt.Fscanf(in, "%d", &num[i])
		}
		fmt.Fscanf(in, "%d", &k)
		fmt.Fprintf(out, "Term %d of the sequence is %d\n", n+k, solve(num, n, k))
	}
}
