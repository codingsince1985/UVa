// UVa 11417 - GCD

package main

import (
	"fmt"
	"os"
)

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	in, _ := os.Open("11417.in")
	defer in.Close()
	out, _ := os.Create("11417.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		var total int
		for i := 1; i < n; i++ {
			for j := i + 1; j <= n; j++ {
				total += gcd(i, j)
			}
		}
		fmt.Fprintln(out, total)
	}
}
