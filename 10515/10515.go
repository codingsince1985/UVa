// UVa 10515 - Powers Et Al.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	in, _ := os.Open("10515.in")
	defer in.Close()
	out, _ := os.Create("10515.out")
	defer out.Close()

	var ms, ns string
	for {
		if fmt.Fscanf(in, "%s%s", &ms, &ns); ms == "0" {
			break
		}
		m, _ := strconv.Atoi(ms[len(ms)-1:])
		start := len(ns) - 2
		if start < 0 {
			start++
		}
		n, _ := strconv.Atoi(ns[start:])
		d := 1
		for i := 1; i <= n; i++ {
			d = (d * m) % 10
		}
		fmt.Fprintln(out, d)
	}
}
