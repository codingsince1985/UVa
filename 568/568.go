// UVa 568 - Just the Facts

package main

import (
	"fmt"
	"os"
)

func trimZero(n int) int {
	for n%10 == 0 {
		n /= 10
	}
	return n
}

func lastDigit(n int) int {
	if n == 0 {
		return 0
	}

	res := 1
	for i := 2; i <= n; i++ {
		res *= trimZero(i)
		res = trimZero(res) % 100000
	}
	return res % 10
}

func main() {
	in, _ := os.Open("568.in")
	defer in.Close()
	out, _ := os.Create("568.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintf(out, "%5d -> %d\n", n, lastDigit(n))
	}
}
