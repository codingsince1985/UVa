// UVa 10515 - Powers Et Al.

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10515.in")
	defer in.Close()
	out, _ := os.Create("10515.out")
	defer out.Close()

	var m, n int
	var ms, ns string
	for {
		if fmt.Fscanf(in, "%s%s", &ms, &ns); ms == "0" {
			break
		}
		fmt.Sscanf(ms[len(ms)-1:], "%d", &m)
		start := len(ns) - 2
		if start < 0 {
			start++
		}
		d := 1
		for fmt.Sscanf(ns[start:], "%d", &n); n > 0; n-- {
			d = (d * m) % 10
		}
		fmt.Fprintln(out, d)
	}
}
