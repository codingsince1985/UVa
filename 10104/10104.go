// UVa 10104 - Euclid Problem

package main

import (
	"fmt"
	"os"
)

func gcd(a, b int) (int, int, int) {
	if b == 0 {
		return 1, 0, a
	}
	x, y, d := gcd(b, a%b)
	return y, x - (a/b)*y, d
}

func main() {
	in, _ := os.Open("10104.in")
	defer in.Close()
	out, _ := os.Create("10104.out")
	defer out.Close()

	var a, b int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &a, &b); err != nil {
			break
		}
		x, y, d := gcd(a, b)
		fmt.Fprintf(out, "%d %d %d\n", x, y, d)
	}
}
