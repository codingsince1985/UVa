// UVa 10994 - Simple Addition

package main

import (
	"fmt"
	"os"
)

func f(n int64) int64 {
	if n <= 0 {
		return 0
	}
	var sum int64
	for n != 0 {
		sum += 45 * (n / 10)
		for i := n/10*10 + 1; i <= n; i++ {
			sum += i % 10
		}
		n /= 10
	}
	return sum
}

func main() {
	in, _ := os.Open("10994.in")
	defer in.Close()
	out, _ := os.Create("10994.out")
	defer out.Close()

	var p, q int64
	for {
		if fmt.Fscanf(in, "%d%d", &p, &q); p < 0 && q < 0 {
			break
		}
		fmt.Fprintln(out, f(q)-f(p-1))
	}
}
