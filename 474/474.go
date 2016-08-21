// UVa 474 - Heads / Tails Probability

package main

import (
	"fmt"
	"os"
)

func calc(n int) (float64, int) {
	base, power := 5.0, 1
	for i := 2; i <= n; i++ {
		base *= 0.5
		for base < 1 {
			base *= 10
			power++
		}
	}
	return base, power
}

func main() {
	in, _ := os.Open("474.in")
	defer in.Close()
	out, _ := os.Create("474.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		b, p := calc(n)
		fmt.Fprintf(out, "2^-%d = %.03fe-%d\n", n, b, p)
	}
}
