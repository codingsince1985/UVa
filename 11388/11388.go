// UVa 11388 - GCD LCM

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11388.in")
	defer in.Close()
	out, _ := os.Create("11388.out")
	defer out.Close()

	var t, gcd, lcm int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		if fmt.Fscanf(in, "%d%d", &gcd, &lcm); lcm%gcd != 0 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, gcd, lcm)
		}
	}
}
