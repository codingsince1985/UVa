// UVa 900 - Brick Wall Patterns

package main

import (
	"fmt"
	"os"
)

func fib() []uint64 {
	fibs := make([]uint64, 51)
	fibs[1], fibs[2] = 1, 2
	for i := 3; i <= 50; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs
}

func main() {
	in, _ := os.Open("900.in")
	defer in.Close()
	out, _ := os.Create("900.out")
	defer out.Close()

	var l int
	fibs := fib()
	for {
		if fmt.Fscanf(in, "%d", &l); l == 0 {
			break
		}
		fmt.Fprintln(out, fibs[l])
	}
}
