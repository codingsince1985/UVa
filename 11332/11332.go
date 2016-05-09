// UVa 11332 - Summing Digits

package main

import (
	"fmt"
	"os"
)

func g(n int) int {
	if n < 10 {
		return n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return g(sum)
}

func main() {
	in, _ := os.Open("11332.in")
	defer in.Close()
	out, _ := os.Create("11332.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, g(n))
	}
}
