// UVa 10450 - World Cup Noise

// f(n) = f(n-2) + f(n-1)
// f(n-2) is the number of sequences starting with 10, followed by n-2 bit
// f(n-1) is the number of sequences starting with 0, followed by n-1 bit

package main

import (
	"fmt"
	"os"
)

var c = func() [51]uint64 {
	var c [51]uint64
	c[0] = 1
	c[1] = 2
	return c
}()

func f(b int) uint64 {
	if c[b] != 0 {
		return c[b]
	}
	c[b] = f(b-1) + f(b-2)
	return c[b]
}

func main() {
	in, _ := os.Open("10450.in")
	defer in.Close()
	out, _ := os.Create("10450.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)
	var b int
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &b)
		fmt.Fprintf(out, "Scenario #%d:\n%v\n\n", i+1, f(b))
	}
}
