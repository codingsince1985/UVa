// UVa 10812 - Beat the Spread!

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10812.in")
	defer in.Close()
	out, _ := os.Create("10812.out")
	defer out.Close()

	var kase, s, d int
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		kase--
		fmt.Fscanf(in, "%d%d", &s, &d)
		if s+d < 0 || (s+d)%2 != 0 {
			fmt.Fprintln(out, "impossible")
			continue
		}
		a := (s + d) / 2
		if b := s - a; b >= 0 {
			fmt.Fprintln(out, a, b)
		} else {
			fmt.Fprintln(out, "impossible")
		}
	}
}
